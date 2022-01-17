package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	chi "github.com/go-chi/chi/v5"
	"golang.org/x/sync/singleflight"
)

var singleflightGroupDownloadReport singleflight.Group

func HandlerDownloadReport(w http.ResponseWriter, r *http.Request) {
	reportID := chi.URLParam(r, "reportID")

	// construct the report path
	reportName := fmt.Sprintf("report-%s.txt", reportID)
	path := filepath.Join(os.TempDir(), reportName)

	sharedProcessKey := fmt.Sprintf("generate %s", reportName)
	_, err, shared := singleflightGroupDownloadReport.Do(sharedProcessKey, func() (interface{}, error) {
		// if the report is not exists, generate it first.
		// otherwise, immediatelly download it
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			log.Println("generate report", reportName, path)

			// simulate long-running process to generate report
			time.Sleep(5 * time.Second)

			f, err := os.Create(path)
			if err != nil {
				f.Close()
				return nil, err
			}

			f.Write([]byte("this is a report"))
			f.Close()
		}

		return true, nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if shared {
		log.Printf("generation of report %v is shared with others", reportName)
	}

	// open the file, download it
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contentDisposition := fmt.Sprintf("attachment; filename=%s", reportName)
	w.Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(w, f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/report/download/{reportID}", HandlerDownloadReport)
	})

	host := ":8080"
	fmt.Printf("starting web server at %s \n", host)
	http.ListenAndServe(host, r)
}
