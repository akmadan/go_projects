package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// upgrader to upgrade HTTP connection to Websocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	fmt.Println("Websocket server started on ws://localhost:8080/ws")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection to websocket: ", err)
		return
	}

	defer conn.Close()

	fmt.Println("Client connected")

	// read messages from the client

	for {
		messageType, message, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("Error reading message: ", err)
			break
		}

		fmt.Println("Message received: ", string(message))

		// delay of 3  seconds
		time.Sleep(3 * time.Second)

		// write message back to the client
		if err := conn.WriteMessage(messageType, message); err != nil {
			fmt.Println("Error writing message: ", err)
			break
		}
	}
}
