package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", handleWebsocket)

	fmt.Println("Websocket server started")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server: ", err)
	}
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection: ", err)
		return
	}

	defer conn.Close()

	fmt.Println("Client Connected")

	for {
		messageType, message, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("Error reading message: ", err)
		}

		fmt.Println("Message Received: ", string(message))

		time.Sleep(3 * time.Second)

		if err := conn.WriteMessage(messageType, message); err != nil {
			fmt.Println("Error writing message: ", err)
			break
		}

	}
}
