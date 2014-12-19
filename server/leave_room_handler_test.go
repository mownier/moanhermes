package server

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func TestLeaveRoomHandlerIfNotMethodDELETE(t *testing.T) {
	leaveRoomHandler := leaveRoomHandler()
	// GET, POST, PUT
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/leave?username=&room_id=", nil)
	w := httptest.NewRecorder()
	leaveRoomHandler.ServeHTTP(w, request)
	
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

func TestLeaveRoomHandlerEmptyUsernameParameterValue(t *testing.T) {
	leaveRoomHandler := leaveRoomHandler()
	request, _ := http.NewRequest("DELETE", "localhost:8080/chat/room/leave?room_id=123&username=", nil)
	w := httptest.NewRecorder()
	leaveRoomHandler.ServeHTTP(w, request)

	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	r := response.(map[string]interface{})

	if w.Code != http.StatusBadRequest {
		t.Error("Status code should be http.StatusBadRequest.")
	} else if _, ok := r["username"]; !ok {
		t.Error("There is no message key.")
	} else if r["username"] != "Username is required." {
		t.Error("Error message should be 'Username is required.'")
	}
}

func TestLeaveRoomHandlerEmptyRoomIdParameterValue(t *testing.T) {
	leaveRoomHandler := leaveRoomHandler()
	request, _ := http.NewRequest("DELETE", "localhost:8080/chat/room/leave?room_id=&username=mownier", nil)
	w := httptest.NewRecorder()
	leaveRoomHandler.ServeHTTP(w, request)

	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	r := response.(map[string]interface{})

	if w.Code != http.StatusBadRequest {
		t.Error("Status code should be http.StatusBadRequest.")
	} else if _, ok := r["room_id"]; !ok {
		t.Error("There is no message key.")
	} else if r["room_id"] != "Room id is required." {
		t.Error("Error message should be 'Room id is required.'")
	}
}

func TestLeaveRoomHandlerRoomIdNotSetAsParameter(t *testing.T) {
	leaveRoomHandler := leaveRoomHandler()
	request, _ := http.NewRequest("DELETE", "localhost:8080/chat/room/leave?username=mownier", nil)
	w := httptest.NewRecorder()
	leaveRoomHandler.ServeHTTP(w, request)

	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	r := response.(map[string]interface{})

	if w.Code != http.StatusBadRequest {
		t.Error("Status code should be http.StatusBadRequest.")
	} else if _, ok := r["room_id"]; !ok {
		t.Error("There is no message key.")
	} else if r["room_id"] != "Room id is required." {
		t.Error("Error message should be 'Room id is required.'")
	}
}

func TestLeaveRoomHandlerUsernameNotSetAsParameter(t *testing.T) {
	leaveRoomHandler := leaveRoomHandler()
	request, _ := http.NewRequest("DELETE", "localhost:8080/chat/room/leave?room_id=123", nil)
	request.Header.Set("Content-Type", "text/plain; charset=utf-8")
	w := httptest.NewRecorder()
	leaveRoomHandler.ServeHTTP(w, request)

	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	r := response.(map[string]interface{})

	if w.Code != http.StatusBadRequest {
		t.Error("Status code should be http.StatusBadRequest.")
	} else if _, ok := r["username"]; !ok {
		t.Error("There is no message key.")
	} else if r["username"] != "Username is required." {
		t.Error("Error message should be 'Username is required.'")
	}
}

func TestLeaveRoomHandlerRoomNotFound(t *testing.T) {
	leaveRoomHandler := leaveRoomHandler()
	request, _ := http.NewRequest("DELETE", "localhost:8080/chat/room/leave?username=mownier&room_id=123", nil)
	w := httptest.NewRecorder()
	leaveRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	r := response.(map[string]interface{})
	if w.Code != http.StatusNotFound {
		t.Error("Status code is not http.StatusNotFound.")
	} else if _, ok := r["message"]; !ok {
		t.Error("There is no error message key.")
	} else if r["message"] != "Room not found." {
		t.Error("Error message should be 'Room not found.'")
	}
}

func TestLeaveRoomHandlerUserNotFoundInRoom(t *testing.T) {
	var creator *User = NewUser("mownier")
	var member *User = NewUser("juan")
	var room *Room = NewRoom("room123", creator)
	room.Users = append(room.Users, member)
	rooms = append(rooms, room)

	leaveRoomHandler := leaveRoomHandler()
	request, _ := http.NewRequest("DELETE", "localhost:8080/chat/room/leave?username=jane&room_id=" + room.Uid, nil)
	w := httptest.NewRecorder()
	leaveRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	r := response.(map[string]interface{})
	if w.Code != http.StatusNotFound {
		t.Error("Status code is not http.StatusNotFound.")
	} else if _, ok := r["message"]; !ok {
		t.Error("There is no message key.")
	} else if r["message"] != "User not found in the room." {
		t.Error("Error message should be 'User not found in the room.'")
	}
}



































































