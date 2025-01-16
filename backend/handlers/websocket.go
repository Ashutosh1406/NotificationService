package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)
var mutex sync.Mutex

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error Upgrading connection", err)
		return
	}

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(clients, conn)
		mutex.Unlock()
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading messages", err)
			break
		}
		broadcast <- string(message)
	}
}

func BroadcastMessages() {
	for {
		message := <-broadcast
		mutex.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("error broadcasting messages to client", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}
