package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"trading-day-go/internal/handlers"
)

var allowedOrigins = []string{"http://localhost:3000", "http://example.com"}

func main() {
	log.Println("Server started on port 8080")
	http.HandleFunc("/ws", handleWebSocket)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Check against allowed origins
			origin := r.Header.Get("Origin")
			for _, allowed := range allowedOrigins {
				if allowed == origin {
					return true
				}
			}
			return false
		},
	}
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
