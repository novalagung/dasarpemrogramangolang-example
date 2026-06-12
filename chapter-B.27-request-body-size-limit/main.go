package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

const maxBodyBytes = 1 << 20 // 1 MB

func handleUpload(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	var payload map[string]any
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			http.Error(w, "request body too large", http.StatusRequestEntityTooLarge)
			return
		}
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"received": payload})
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	if err := r.ParseMultipartForm(maxBodyBytes); err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			http.Error(w, "file too large (max 1MB)", http.StatusRequestEntityTooLarge)
			return
		}
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "file is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, _ := io.ReadAll(file)
	log.Printf("received file: %s, size: %d bytes", header.Filename, len(data))
	w.Write([]byte("received: " + header.Filename))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /upload/json", handleUpload)
	mux.HandleFunc("POST /upload/file", handleFileUpload)

	log.Println("server started at localhost:9000")
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
