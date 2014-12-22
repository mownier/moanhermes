package server

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func TestRemoveRoomHandlerNotDELETEMethod(t *testing.T) {
	removeRoomHandler := removeRoomHandler()
	request, _ := http.NewRequest("GET", "localhost:8080/chat/room/remove", nil)
	w := httptest.NewRecorder()
	removeRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	

	if w.Code != http.StatusMethodNotAllowed {
		t.Error("Response status code is not http.StatusMethodNotAllowed.")
	} else if response == nil {
		t.Error("Should have a response.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["message"]; !ok {
			t.Error("Should have a message key.")
		} else if r["message"] != "Method not allowed." {
			t.Error("Error message should be 'Method not allowed.'")
		}
	}
}

func TestRemoveRoomHandlerEmptyUsernameParameter(t *testing.T) {
	removeRoomHandler := removeRoomHandler()
	request, _ := http.NewRequest("DELETE", "localhost:8080/chat/room/remove?username=&room_id=", nil)
	w := httptest.NewRecorder()
	removeRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code != http.StatusBadRequest {
		t.Error("Response status code is not http.StatusBadRequest.")
	} else if response == nil {
		t.Error("Should have a response.")
	} else  {
		r := response.(map[string]interface{})
		if _, ok := r["username"]; !ok {
			t.Error("Should have a message key 'username'.")
		} else if r["username"] != "Username is required." {
			t.Error("Error message should be 'Username is required.'")
		}
	}
}

func TestRemoveHandlerUsernameParameterIsNotSet(t *testing.T) {
	removeRoomHandler := removeRoomHandler()
	request, _ := http.NewRequest("DELETE", "localhost:8080/chat/room/remove?room_id=", nil)
	w := httptest.NewRecorder()
	removeRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code != http.StatusBadRequest {
		t.Error("Response status code is not http.StatusBadRequest.")
	} else if response == nil {
		t.Error("Should have a response.")
	} else  {
		r := response.(map[string]interface{})
		if _, ok := r["username"]; !ok {
			t.Error("Should have a message key 'username'.")
		} else if r["username"] != "Username is required." {
			t.Error("Error message should be 'Username is required.'")
		}
	}
}

func TestRemoveHandlerEmptyRoomIdParameter(t *testing.T) {
	removeRoomHandler := removeRoomHandler()
	request, _ := http.NewRequest("DELETE", "localhost:8080/chat/room/remove?username=&room_id=", nil)
	w := httptest.NewRecorder()
	removeRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code != http.StatusBadRequest {
		t.Error("Response status code is not http.StatusBadRequest.")
	} else if response == nil {
		t.Error("Should have a response.")
	} else  {
		r := response.(map[string]interface{})
		if _, ok := r["room_id"]; !ok {
			t.Error("Should have a message key 'room_id'.")
		} else if r["room_id"] != "Room id is required." {
			t.Error("Error message should be 'Room id is required.'")
		}
	}
}

func TestRemoveHandlerRoomIdParameterIsNotSet(t *testing.T) {
	removeRoomHandler := removeRoomHandler()
	request, _ := http.NewRequest("DELETE", "localhost:8080/chat/room/remove?username=", nil)
	w := httptest.NewRecorder()
	removeRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code != http.StatusBadRequest {
		t.Error("Response status code is not http.StatusBadRequest.")
	} else if response == nil {
		t.Error("Should have a response.")
	} else  {
		r := response.(map[string]interface{})
		if _, ok := r["room_id"]; !ok {
			t.Error("Should have a message key 'room_id'.")
		} else if r["room_id"] != "Room id is required." {
			t.Error("Error message should be 'Room id is required.'")
		}
	}
}

func TestRemoveHandlerIfRoomDoesNotExist(t *testing.T) {
	removeRoomHandler := removeRoomHandler()
	request, _ := http.NewRequest("DELETE", "localhost:8080/chat/room/remove?username=mownier&room_id=123", nil)
	w := httptest.NewRecorder()
	removeRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	if w.Code != http.StatusNotFound {
		t.Error("Response status code is not http.StatusNotFound")
	} else if response == nil {
		t.Error("Should have a response.") 
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["message"]; !ok {
			t.Error("Should have a message key 'message'.")
		} else if r["message"] != "Room not found." {
			t.Error("Error message should be 'Room not found.'")
		}
	}
}

func TestRemoveHandlerRemoveTriggeredNotByTheCreator(t *testing.T) {
	var creator *User = NewUser("mownier")
	var room *Room = NewRoom("room123", creator)
	var member *User = NewUser("juan")
	room.Users = append(room.Users, member)
	rooms = append(rooms, room)
	removeRoomHandler := removeRoomHandler()
	request, _ := http.NewRequest("DELETE", "localhost:8080/chat/room/remove?username=" + member.Username + "&room_id=" + room.Uid, nil)
	w := httptest.NewRecorder()
	removeRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	if w.Code != http.StatusUnauthorized {
		t.Error("Response status code is not http.StatusUnauthorized")
	} else if response == nil {
		t.Error("Should have a response.") 
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["message"]; !ok {
			t.Error("Should have a message key 'message'.")
		} else if r["message"] != "Unauthorized to remove the room." {
			t.Error("Error message should be 'Unauthorized to remove the room.'")
		}
	}
}

func TestRemoveHandlerRemoveTriggeredByTheCreator(t *testing.T) {
	var creator *User = NewUser("mownier")
	var room *Room = NewRoom("room123", creator)
	var member *User = NewUser("juan")
	room.Users = append(room.Users, member)
	rooms = append(rooms, room)
	var numberOfRoomsBeforeRemove int = len(rooms)

	removeRoomHandler := removeRoomHandler()
	request, _ := http.NewRequest("DELETE", "localhost:8080/chat/room/remove?username=" + creator.Username + "&room_id=" + room.Uid, nil)
	w := httptest.NewRecorder()
	removeRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	if w.Code != http.StatusOK {
		t.Error("Response status code is not http.StatusOk")
	} else if response == nil {
		t.Error("Should have a response.") 
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["message"]; !ok {
			t.Error("Should have a message key 'message'.")
		} else if r["message"] != "Successfully removed the room." {
			t.Error("Error message should be 'Successfully removed the room.'")
		} else {
			var numberOfRoomsAfterRemove int = len(rooms)
			if numberOfRoomsAfterRemove != numberOfRoomsBeforeRemove - 1 {
				t.Error("Room is not successfully removed from the array 'rooms'.")
			}
		}
	}
}


















