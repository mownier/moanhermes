package server

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
	"bytes"
	"encoding/json"
)

func TestLeaveRoomHandlerIfNotMethodDELETE(t *testing.T) {
	var creator *User = NewUser("mownier")
	var member *User = NewUser("juan")
	var room *Room = NewRoom("room123", creator)
	room.Users = append(room.Users, member)
	rooms = append(rooms, room)

	leaveRooomHandler := leaveRoomHandler()
	params := url.Values{}
	params.Add("username", member.Username)
	params.Add("room_id", room.Uid)
	// GET, POST, PUT
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/leave", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-urlencoded; param=value")
	w := httptest.NewRecorder()
	leaveRooomHandler.ServeHTTP(w, request)
	
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	r := response.(map[string]interface{})

	if w.Code != http.StatusMethodNotAllowed {
		t.Error("Status code should be 'http.StatusMethodNotAllowed'.")
	} else if _, ok := r["message"]; !ok {
		t.Error("There is no message key.")
	} else if r["message"] != "Method not allowed." {
		t.Error("Error message should be 'Method not allowed.'")
	}
}
