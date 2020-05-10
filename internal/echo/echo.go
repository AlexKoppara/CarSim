package echo

import (
	"fmt"
	"log"
	"time"

	"github.com/AlexKoppara/CarSim/internal/models"
	"github.com/gorilla/websocket"
)

// Echo endpoint
func Echo(clients map[*websocket.Conn]bool) {
	car := models.Car{
		Make:  "Toyota",
		Model: "Camry",
		Lat:   0.0,
		Long:  0.0,
		Vx:    1.0,
		Vy:    1.0,
	}

	for {
		time.Sleep(time.Second)
		car.Drive(1)

		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%f", car.Lat)))
			if err != nil {
				log.Printf("Websocket error: %s", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
