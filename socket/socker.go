package socket

import (
	"encoding/json"
	"fmt"
	"go-lang/blinkchat/models"
	"go-lang/blinkchat/services"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}

	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		fmt.Println(messageType)
		mess := models.Message{}

		err = json.Unmarshal(message, &mess)

		if err != nil {
			fmt.Println(err)
		}

		// fmt.Println(mess)

		if err != nil {
			fmt.Println(err)
		}

		res, err := services.CreateMessage(&mess)

		if err != nil {
			fmt.Println("Error creating message in database:", err)
			response := map[string]string{"error": "Failed to save message"}
			responseJSON, _ := json.Marshal(response)
			_ = conn.WriteMessage(websocket.TextMessage, responseJSON)
			continue
		}

		responseJson, err := json.Marshal(res)

		err = conn.WriteMessage(messageType, responseJson)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}

	}
}
