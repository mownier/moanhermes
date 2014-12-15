package server

type User struct {
	username string
}

func newUser(username string) *User {
	return &User {
		username: username,
	}
}

type Room struct {
	uid string
	name string
	users []*User
}

func newRoom(name string, creator *User) *Room {
	users := make([]*User, 0)
	users = append(users, creator)
	return &Room {
		name: name,
		users: users,
	}
}

type Moanhermes struct {
	rooms []*Room
}

func (m *Moanhermes) broadcastMessage(message, roomUid string) {

}

func (m *Moanhermes) StartServing(address string) {

}

func (m *Moanhermes) CreateRoom(roomName, username string) *Room {
	return newRoom(roomName, newUser(username));
}

func NewMoanhermes() *Moanhermes {
	return &Moanhermes {
		rooms: make([]*Room, 0),
	}
}
