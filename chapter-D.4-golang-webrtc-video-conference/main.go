package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type M map[string]interface{}

var connectionsMap sync.Map

func main() {
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadFile("index.html")
		if err != nil {
			http.Error(w, "Could not serve html file", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-type", "text/html")
		w.Write(buf)
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		currentGorillaConn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
			return
		}

		roomID := r.URL.Query().Get("room")
		userID := r.URL.Query().Get("user")
		currentConn := WebSocketConnection{Conn: currentGorillaConn, RoomID: roomID, UserID: userID}
		manager.AppendConnection(&currentConn)

		go handleIO(&currentConn)
	})

	log.Println("Server starting")
	err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleIO(currentConn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	for {
		payload := make(M)
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				manager.EjectConnection(currentConn)
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		log.Println("==============", payload)

		siblingConnections := manager.GetSiblingConnections(currentConn)
		for _, each := range siblingConnections {
			payload["from"] = currentConn.UserID
			payload["room"] = each.RoomID
			each.WriteJSON(payload)
		}
	}
}
