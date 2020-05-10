package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WsHandler adds new upgrader client to client map
func WsHandler(w http.ResponseWriter, r *http.Request, clients map[*websocket.Conn]bool, upgrader *websocket.Upgrader) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// add client to map
	clients[ws] = true
}
