package main

import (
	"log"
	"net/http"
)

type CustomMux struct {
	mux *http.ServeMux
}

func NewCustomMux() *CustomMux {
	return &CustomMux{mux: http.NewServeMux()}
}

func (c *CustomMux) Handle(pattern string, handler http.Handler) {
	c.mux.Handle(pattern, handler)
}

func (c *CustomMux) HandleFunc(pattern string, handler http.HandlerFunc) {
	c.mux.HandleFunc(pattern, handler)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.mux.ServeHTTP(w, r)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - Halaman tidak ditemukan"))
}

func internalErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Terjadi kesalahan pada server"))
}

func panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v", err)
				internalErrorHandler(w, r)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})

	mux.HandleFunc("GET /panic", func(w http.ResponseWriter, r *http.Request) {
		panic("something went wrong")
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			notFoundHandler(w, r)
			return
		}
		w.Write([]byte("Home"))
	})

	handler := panicMiddleware(mux)

	log.Println("server started at localhost:9000")
	err := http.ListenAndServe(":9000", handler)
	if err != nil {
		log.Fatal(err)
	}
}
