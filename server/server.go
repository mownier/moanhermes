package server

import (
	"net/http"
	"log"
	"encoding/json"
	"crypto/sha1"
	"encoding/base64"
	"strconv"
	"time"
)

type User struct {
	Username string
	Online bool
}

func NewUser(username string) *User {
	return &User {
		Username: username,
	}
}

type Message struct {
	User *User
	Content string
	DatePosted time.Time
}

func NewMessage(user *User, content string) *Message {
	return &Message {
		User: user,
		Content: content,
		DatePosted: time.Now(),
	}	
}

type Room struct {
	Uid string
	Name string
	Users []*User
	Messages []*Message
}

func NewRoom(name string, creator *User) *Room {
	hasher := sha1.New()
	timestamp := strconv.Itoa(int(time.Now().UnixNano()))
	hasher.Write([]byte(name + timestamp))
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
var users = make([]*User, 0)

type Moanhermes struct {

}

func broadcastMessage(message, roomUid string) {

}

func (m *Moanhermes) StartServing(address string) {
	http.HandleFunc("/chat/room/create"     , createRoomHandler())
	http.HandleFunc("/chat/room/join"       , joinRoomHandler())
	http.HandleFunc("/chat/room/leave"      , leaveRoomHandler())
	http.HandleFunc("/chat/room/invite"     , inviteRoomHandler)
	http.HandleFunc("/chat/room/remove"     , removeRoomHandler())
	http.HandleFunc("/chat/message/compose" , composeMessageHandler)
	http.HandleFunc("/chat/message/remove"  , removeMessageHandler)
	http.HandleFunc("/chat/register"        , registerHandler())
	http.HandleFunc("/chat/signin"          , signinHandler())
	http.HandleFunc("/chat/signout"         , signoutHandler)
	log.Fatal(http.ListenAndServe(address, nil))
}

// HTTP Handlers

// METHOD: POST
// PARAMS: room_name, username
func createRoomHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var responseString []byte
		var responseStatusCode int

		if r.Method != "POST" {
			// Setting method not allowed status
			responseStatusCode = http.StatusMethodNotAllowed
			// Setting the response string
			responseString = []byte("{\"message\" : \"Method not allowed.\"}")
		} else {
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
				// Setting the response status code
				responseStatusCode = http.StatusOK
				// Setting the response string
				responseString = []byte("{\"message\" : \"Successfully created a room.\"}")
			} else {
				// Creating a map for errors
				errors := make(map[string]interface{})
				// For non-existent room name value
				if !hasRoomName {
					errors["room_name"] = "Room name is required."
				} 
				// For non-existent username value
				if !hasUsername {
					errors["username"] = "Username is required."
				}
				// Coverting map to json string 
				jsonString, _ := json.Marshal(errors)
				// Setting the status code
				responseStatusCode = http.StatusBadRequest
				// Setting the response string
				responseString = jsonString
			}
		}
		// Setting the response as json format
		w.Header().Set("Content-Type", "application/json")
		// Returning as 405 Method Not Allowed
		w.WriteHeader(responseStatusCode)
		// Writing the  json response
		w.Write(responseString)
	})
}

// METHOD: POST
// PARAMS: room_id, username
func joinRoomHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var responseString []byte
		var responseStatusCode int

		if r.Method != "POST" {
			responseString = []byte("{\"message\" : \"Method not allowed.\"}")
			responseStatusCode = http.StatusMethodNotAllowed
		} else {
			var roomId string = r.FormValue("room_id")
			var username string = r.FormValue("username")
			var hasRoomId bool = len(roomId) > 0
			var hasUsername bool = len(username) > 0

			if hasRoomId && hasUsername {
				var roomDoesExist bool = false
				var room *Room
				for i := 0; i < len(rooms); i++ {
					var r *Room = rooms[i]
					if r.Uid == roomId {
						room = r
						roomDoesExist = true
						break;
					}
				}
				if !roomDoesExist {
					responseStatusCode = http.StatusNotFound
					responseString = []byte("{\"message\" : \"Room not found.\"}")
				} else {
					var creator *User = room.Users[0]
					if creator.Username == username {
						responseStatusCode = http.StatusBadRequest
						responseString = []byte("{\"message\" : \"Already joined.\"}")
					} else {
						var user *User = NewUser(username)
						room.Users = append(room.Users, user)
						responseStatusCode = http.StatusOK
						responseString = []byte("{\"message\" : \"Successfully joined.\"}")
					}
				}
			} else {
				errors := make(map[string]interface{})
				if !hasRoomId {
					errors["room_id"] = "Room id is required."
				}
				if !hasUsername {
					errors["username"] = "Username is required."
				}

				jsonString, _ := json.Marshal(errors)
				responseString = jsonString
				responseStatusCode = http.StatusBadRequest
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatusCode)
		w.Write(responseString)
	})
}

// METHOD: DELETE
// PARAMS: room_id, username
func leaveRoomHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var responseString []byte
		var responseStatusCode int

		if r.Method != "DELETE" {
			responseStatusCode = http.StatusMethodNotAllowed
			responseString = []byte("{\"message\" : \"Method not allowed.\"}")
		} else {
			r.ParseForm()
			var username string = r.URL.Query().Get("username")
			var roomId string = r.URL.Query().Get("room_id")
			var hasUsername bool = len(username) > 0
			var hasRoomId bool = len(roomId) > 0
			if hasUsername && hasRoomId {
				var roomDoesExist bool
				var room *Room
				var roomIndex int
				for i := 0; i < len(rooms); i++ {
					var r *Room = rooms[i]
					if r.Uid == roomId {
						roomDoesExist = true
						room = r
						roomIndex = i
						break
					}
				} 
				if !roomDoesExist {
					responseStatusCode = http.StatusNotFound
					responseString = []byte("{\"message\" : \"Room not found.\"}")
				} else {
					var userDoesExist bool
					var user *User
					var userIndex int
					for i := 0; i < len(room.Users); i++ {
						var u *User = room.Users[i]
						if u.Username == username {
							userDoesExist = true
							user = u
							userIndex = i
							break
						}
					}
					if !userDoesExist {
						responseStatusCode = http.StatusNotFound
						responseString = []byte("{\"message\" : \"User not found in the room.\"}")
					} else {
						var creator *User = room.Users[0]
						var isTheCreator bool = creator.Username == user.Username
						if isTheCreator {
							// Removing a room in rooms
							rooms = append(rooms[:roomIndex], rooms[roomIndex + 1:]...)
						} else {
							// Removing an user in Users
							room.Users = append(room.Users[:userIndex], room.Users[userIndex + 1:]...)
						}
						responseStatusCode = http.StatusOK
						responseString = []byte("{\"message\" : \"Successfully left the room.\"}")
					}
				}
			} else {
				errors := make(map[string]interface{})
				if !hasUsername {
					errors["username"] = "Username is required."
				}
				if !hasRoomId {
					errors["room_id"] = "Room id is required."
				}
				jsonString, _ := json.Marshal(errors)

				responseStatusCode = http.StatusBadRequest
				responseString = jsonString
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatusCode)
		w.Write(responseString)
	})
}

func inviteRoomHandler(w http.ResponseWriter, r *http.Request) {

}

// METHOD : DELETE
// PARAMS : username, room_id
func removeRoomHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var responseString []byte
		var responseStatusCode int

		if r.Method != "DELETE" {
			responseStatusCode = http.StatusMethodNotAllowed
			responseString = []byte("{\"message\" : \"Method not allowed.\"}")
		} else {
			var username string = r.URL.Query().Get("username")
			var roomId string = r.URL.Query().Get("room_id")
			var hasUsername bool = len(username) > 0
			var hasRoomId bool = len(roomId) > 0

			if hasUsername && hasRoomId {
				var roomDoesExist bool
				var room *Room
				var roomIndex int
				for i := 0; i < len(rooms); i++ {
					var r *Room = rooms[i]
					if r.Uid == roomId {
						room = r
						roomDoesExist = true
						roomIndex = i
						break
					}
				}
				if !roomDoesExist {
					responseStatusCode = http.StatusNotFound
					responseString = []byte("{\"message\" : \"Room not found.\"}")
				} else {
					var creator *User = room.Users[0]
					var isTheCreator bool = username == creator.Username
					if !isTheCreator {
						responseStatusCode = http.StatusUnauthorized
						responseString = []byte("{\"message\" : \"Unauthorized to remove the room.\"}")
					} else {
						rooms = append(rooms[:roomIndex], rooms[roomIndex + 1:]...)
						responseStatusCode = http.StatusOK
						responseString = []byte("{\"message\" : \"Successfully removed the room.\"}")
					}
				}
			} else {
				errors := make(map[string]interface{})
				if !hasUsername {
					errors["username"] = "Username is required."
				}
				if !hasRoomId {
					errors["room_id"] = "Room id is required."
				}
				jsonString, _ := json.Marshal(errors)
				responseString = jsonString
				responseStatusCode = http.StatusBadRequest
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatusCode)
		w.Write(responseString)
	})
}

func composeMessageHandler(w http.ResponseWriter, r *http.Request) {

}

func removeMessageHandler(w http.ResponseWriter, r *http.Request) {
	
}

// METHOD : POST
// PARAMS : username
func registerHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var responseString []byte
		var responseStatusCode int

		if r.Method != "POST" {
			responseStatusCode = http.StatusMethodNotAllowed
			responseString = []byte("{\"message\" : \"Method not allowed.\"}")
		} else {
			var username string = r.FormValue("username")
			var hasUsername bool = len(username) > 0

			if !hasUsername {
				errors := make(map[string]interface{})
				errors["username"] = "Username is required."
				responseStatusCode = http.StatusBadRequest
				responseString = []byte("{\"username\" : \"Username is required.\"}")
			} else {
				var userDoesExist bool
				// var userIndex int
				for i := 0; i < len(users); i++ {
					var u *User = users[i]
					if u.Username == username {
						// userIndex = i
						userDoesExist = true
						break
					}
				}
				if userDoesExist {
					responseStatusCode = http.StatusBadRequest
					responseString = []byte("{\"message\" : \"Username already exists.\"}")
				} else {
					var newUser *User = NewUser(username)
					users = append(users, newUser)
					responseStatusCode = http.StatusOK
					responseString = []byte("{\"message\" :\"Successfully registered.\"}")
				}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatusCode)
		w.Write(responseString)
	})
}

func signinHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var responseString []byte
		var responseStatusCode int

		if r.Method != "POST" {
			responseStatusCode = http.StatusMethodNotAllowed
			responseString = []byte("{\"message\" :\"Method not allowed.\"}")
		} else {
			var username string = r.FormValue("username")
			var hasUsername bool = len(username) > 0
			if hasUsername {
				var userDoesExist bool
				var user *User
				for i := 0; i < len(users); i++ {
					var u *User = users[i]
					if u.Username == username {
						userDoesExist = true
						user = u
						break
					}
				}
				if !userDoesExist {
					responseStatusCode = http.StatusNotFound
					responseString = []byte("{\"message\" : \"User does not exist.\"}")
				} else {
					user.Online = true
					responseStatusCode = http.StatusOK
					responseString = []byte("{\"message\" : \"Successfully signed in.\"}")
				}
			} else {
				errors := make(map[string]interface{})
				errors["username"] = "Username is required."
				jsonString, _ := json.Marshal(errors)
				responseString = jsonString
				responseStatusCode = http.StatusBadRequest
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatusCode)
		w.Write(responseString)
	})
}

func signoutHandler(w http.ResponseWriter, r *http.Request) {

}

func NewMoanhermes() *Moanhermes {
	return &Moanhermes {
	}
}
