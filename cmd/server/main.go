package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"trading-day-go/internal/handlers"
)

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Read and write messages from/to the WebSocket connection
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		handlers.HandleMessage(message)

		// err = conn.WriteMessage(websocket.TextMessage, []byte("pong"))
		// if err != nil {
		// 	log.Println("write:", err)
		// 	break
		// }
	}
}
