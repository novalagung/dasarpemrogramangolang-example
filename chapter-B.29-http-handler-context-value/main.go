package main

import (
	"context"
	"log"
	"net/http"
)

type contextKey string

const (
	contextKeyUser      contextKey = "user"
	contextKeyRequestID contextKey = "request_id"
)

type User struct {
	Username string
	Role     string
}

func middlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, _, ok := r.BasicAuth()
		if !ok || username == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		user := User{Username: username, Role: "admin"}
		ctx := context.WithValue(r.Context(), contextKeyUser, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func middlewareRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = "auto-generated-001"
		}
		ctx := context.WithValue(r.Context(), contextKeyRequestID, requestID)
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(contextKeyUser).(User)
	if !ok {
		http.Error(w, "user not found in context", http.StatusInternalServerError)
		return
	}

	requestID, _ := r.Context().Value(contextKeyRequestID).(string)
	log.Printf("[%s] profile accessed by %s (%s)", requestID, user.Username, user.Role)

	w.Write([]byte("Hello, " + user.Username + " [" + user.Role + "]"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /profile", handleProfile)

	handler := middlewareRequestID(middlewareAuth(mux))

	log.Println("server started at localhost:9000")
	err := http.ListenAndServe(":9000", handler)
	if err != nil {
		log.Fatal(err)
	}
}
