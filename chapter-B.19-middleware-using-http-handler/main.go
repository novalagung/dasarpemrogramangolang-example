package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.Write([]byte("\n"))
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func main() {
	mux := http.DefaultServeMux

	mux.HandleFunc("/student", ActionStudent)

	var handler http.Handler = mux
	handler = MiddlewareAuth(handler)
	handler = MiddlewareAllowOnlyGet(handler)

	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = handler

	log.Println("server started at localhost:9000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// ===== Test

// curl -X GET --user batman:secret http://localhost:9000/student
// curl -X GET --user batman:secret http://localhost:9000/student?id=s001
