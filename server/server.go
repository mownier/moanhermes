package server

import (
	"net/http"
	"log"
	"crypto/sha1"
	"encoding/base64"
)

type User struct {
	Username string
}

func NewUser(username string) *User {
	return &User {
		Username: username,
	}
}

type Room struct {
	Uid string
	Name string
	Users []*User
}

func NewRoom(name string, creator *User) *Room {
	hasher := sha1.New()
	hasher.Write([]byte(name))
	uid := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	users := make([]*User, 0)
	users = append(users, creator)

	return &Room {
		Uid: uid,
		Name: name,
		Users: users,
	}
}

type Moanhermes struct {
	Rooms []*Room
}

func (m *Moanhermes) broadcastMessage(message, roomUid string) {

}

func (m *Moanhermes) StartServing(address string) {
	log.Fatal(http.ListenAndServe(address, nil))
}

func (m *Moanhermes) CreateRoom(roomName, username string) *Room {
	return NewRoom(roomName, NewUser(username));
}

func NewMoanhermes() *Moanhermes {
	return &Moanhermes {
		Rooms: make([]*Room, 0),
	}
}
