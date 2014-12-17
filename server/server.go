package server

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
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

// METHOD: POST
// PARAMS: room_name, username
func createRoomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// TODO Method not found
		// Setting the response as json format
		w.Header().Set("Content-Type", "application/json")
		// Returning as 405 Method Not Allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		// Creating a json string
		var jsonString []byte = []byte("{\"message\" : \"Method not allowed.\"}")
		// Writing the json response
		w.Write(jsonString)
		// Writing in the command line
		fmt.Println(string(jsonString))
	} else {
		// TODO Parse parameters: room_name, username
		// Getting the room name from the form values
		var roomName string = r.FormValue("room_name")
		// Getting the username from the form values
		var username string = r.FormValue("username")
		// Checking if room name parameter does exist
		var hasRoomName bool = len(roomName) > 0
		// Checking if username parameter does exist 
		var hasUsername bool = len(username) > 0

		if hasRoomName && hasUsername {
			// Appending a new room
			rooms = append(rooms, NewRoom(roomName, NewUser(username)))
			// Setting the response as json format
			w.Header().Set("Content-Type", "application/json")
			// Return as 200 OK
			w.WriteHeader(http.StatusOK)
			// Creating a json string
			var jsonString []byte = []byte("{\"message\" : \"Successfully created a room.\"}")
			// Writing the json response
			w.Write(jsonString)
			// Writing in the command line
			fmt.Println(string(jsonString))
		} else {
			// Creating a map for errors
			errors := make(map[string]string)
			// For non-existent room name value
			if !hasRoomName {
				errors["room_name"] = "Room name is required."
			} 
			// For non-existent username value
			if !hasUsername {
				errors["username"] = "Username is required."
			}
			// Coverting map to json stirng 
			jsonString, _ := json.Marshal(errors)
			// Setting the response as json format
			w.Header().Set("Content-Type", "application/json")
			// Writing the json response
			w.Write(jsonString)
			// Writing in the command line
			fmt.Println(string(jsonString))
		}
	}
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
