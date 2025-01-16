package main

import (
	"log"
	"net/http"

	"github.com/Ashutosh1406/NotificationService/backend/handlers"
)

func main() {
	http.HandleFunc("/ws", handlers.HandleWebSocket)
	log.Println("Server is Running on Port 8044")
	log.Fatal(http.ListenAndServe(":8044", nil))
}
