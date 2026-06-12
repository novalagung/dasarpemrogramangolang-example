package main

import (
	"log"
	"net/http"
	"time"

	"github.com/novalagung/gubrak/v2"
)

type M map[string]any

var cookieName = "CookieData"

func ActionIndex(w http.ResponseWriter, r *http.Request) {

	c := &http.Cookie{}

	if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
		c = storedCookie
	}

	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = gubrak.RandomString(32)
		c.Expires = time.Now().Add(5 * time.Minute)
		http.SetCookie(w, c)
	}

	w.Write([]byte(c.Value))
}

func ActionDelete(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{}
	c.Name = cookieName
	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/", ActionIndex)
	http.HandleFunc("/delete", ActionDelete)

	log.Println("server started at localhost:9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
