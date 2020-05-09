package main

import (
	"log"
	"net/http"

	"github.com/AlexKoppara/CarSim/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/echo", handlers.Echo)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
