package server

import (
	"net/http"
	"fmt"
	"log"
)

type User struct {
	Username string
}

func NewUser(username string) *User {
	return &User {
		username: username,
	}
}

type Room struct {
	Uid string
	Name strings
	Users []*User
}

func NewRoom(name string, creator *User) *Room {
	users := make([]*User, 0)
	users = append(users, creator)
	return &Room {
		name: name,
		users: users,
	}
}

var rooms = make([]*Room, 0)

type Moanhermes struct {

}

func broadcastMessage(message, roomUid string) {

}

func (m *Moanhermes) StartServing(address string) {
	http.HandleFunc("/chat/room/create"     , createRoomHandler)
	http.HandleFunc("/chat/room/join"       , joinRoomHandler)
	http.HandleFunc("/chat/room/leave"      , leaveRoomHandler)
	http.HandleFunc("/chat/room/invite"     , inviteRoomHandler)
	http.HandleFunc("/chat/room/remove"     , removeRoomHandler)
	http.HandleFunc("/chat/message/compose" , composeMessageHandler)
	http.HandleFunc("/chat/message/remove"  , removeMessageHandler)
	http.HandleFunc("/chat/register"        , registerHandler)
	http.HandleFunc("/chat/signin"          , signinHandler)
	http.HandleFunc("/chat/signout"         , signoutHandler)
	log.Fatal(http.ListenAndServe(address, nil))
}

// HTTP Handlers

func createRoomHandler(w http.ResponseWriter, r *http.Request) {
	rooms = append(rooms, NewRoom("room-name", NewUser("mownier")))
	fmt.Println(rooms[0].)
}

func joinRoomHandler(w http.ResponseWriter, r *http.Request) {
	
}

func leaveRoomHandler(w http.ResponseWriter, r *http.Request) {

}

func inviteRoomHandler(w http.ResponseWriter, r *http.Request) {

}

func removeRoomHandler(w http.ResponseWriter, r *http.Request) {

}

func composeMessageHandler(w http.ResponseWriter, r *http.Request) {

}

func removeMessageHandler(w http.ResponseWriter, r *http.Request) {
	
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	
}

func signoutHandler(w http.ResponseWriter, r *http.Request) {

}

func NewMoanhermes() *Moanhermes {
	return &Moanhermes {
	}
}
