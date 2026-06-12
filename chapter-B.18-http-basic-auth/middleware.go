package main

import "net/http"

const USERNAME = "batman"
const PASSWORD = "secret"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		http.Error(w, "wrong username/password", http.StatusUnauthorized)
		return false
	}

	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
		return false
	}

	return true
}
