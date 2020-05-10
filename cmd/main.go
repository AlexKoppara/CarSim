package main

import (
	"log"
	"net/http"

	"github.com/AlexKoppara/CarSim/internal/echo"
	"github.com/AlexKoppara/CarSim/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func main() {
	clients := make(map[*websocket.Conn]bool)
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	r := mux.NewRouter()
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Print("test")
		handlers.WsHandler(w, r, clients, upgrader)
	})

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte("healthy"))
	})

	go echo.Echo(clients)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
