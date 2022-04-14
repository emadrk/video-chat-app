package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// AllRooms in the global hashmap
var AllRooms RoomMap

// createRoomRequestHandler and return room id
func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	roomID := AllRooms.CreateRoom()

	type resp struct {
		RoomID string `json:"room_id"`
	}
	json.NewEncoder(w).Encode(resp{RoomID: roomID})
	log.Println(AllRooms.Map)

}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type broadcastMsg struct {
	Message map[string]interface{}
	RoomID  string
	Client  *websocket.Conn
}

var broadcast = make(chan broadcastMsg)

func broadcaster() {
	for {
		msg := <-broadcast

		for _, client := range AllRooms.Map[msg.RoomID] {
			fmt.Println(client)
			if client.Conn != msg.Client {
				err := client.Conn.WriteJSON(msg.Message)
				if err != nil {
					log.Fatal(err)
					client.Conn.Close()
				}

			}

		}
	}

}

// JoinRoomRequestHandler will join thr client in particular room
func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Query())
	roomID, ok := r.URL.Query()["roomID"]
	log.Println(roomID)

	if !ok {
		log.Println("roomID missing in URL papameters")
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("websocket upgrade error", err)
	}
	AllRooms.InsertIntoRoom(roomID[0], false, ws)
	go broadcaster()
	for {
		var msg broadcastMsg
		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Fatal("Read error:", err)
		}

		msg.Client = ws
		msg.RoomID = roomID[0]
		log.Println(msg.Message)

		broadcast <- msg

	}

}
