package handlers

import (
	"log"
)

func HandleMessage(msg []byte) {
	log.Printf("received: %s\n", msg)
}