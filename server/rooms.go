package server

import (
	"log"
	"sync"
	"time"

	"math/rand"

	"github.com/gorilla/websocket"
)

// Participant describe a single entity in hashmap
type Participant struct {
	Host bool
	Conn *websocket.Conn
}

// RoomMap is main hashmap [roomID string] -> []Participant
type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

// INIT initializes the RoomMap struct
func (r *RoomMap) Init() {
	// r.Mutex.Lock()
	// defer r.Mutex.Unlock()
	r.Map = make(map[string][]Participant)
}

func (r *RoomMap) Get(roomID string) []Participant {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.Map[roomID]

}

// CreateRoom generates a unique room id and return it -> insert it in hashmap
func (r *RoomMap) CreateRoom() string {
	// generate a uniqueid -> insert it in hashmap
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	roomID := string(b)
	r.Map[roomID] = []Participant{}
	return roomID
}

// InsertIntoRoom will create a participant and add it to hashmap
func (r *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{host, conn}
	log.Println("Inserting into Room with RoomID:", roomID)
	r.Map[roomID] = append(r.Map[roomID], p)

}

// DeleteRoom deletes the room with the roomID
func (r *RoomMap) DeleteRoom(roomID string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomID)

}
