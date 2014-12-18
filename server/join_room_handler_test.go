package server

import (
	"testing"
	"net/url"
	"net/http"
	"net/http/httptest"
	"bytes"
	"encoding/json"
)

func TestJoinRoomHandlerShouldRespondWithErrorMessageIfRoomIdNotSet(t *testing.T) {
	joinRoomHandler := joinRoomHandler()
	params := url.Values{}
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/join", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	joinRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	r := response.(map[string]interface{})
	if _, ok := r["room_id"]; !ok {
		t.Error("Should respond with an error message if room id is not set.")
	} else if r["room_id"] != "Room id is required." {
		t.Error("Room id error message should be 'Room id is required.'")
	}
}

func TestJoinRoomHandlerShouldRespondWithErrorMessageIfEmptyRoomIdParameter(t *testing.T) {
	joinRoomHandler := joinRoomHandler()
	params := url.Values{}
	params.Add("room_id", "")
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/join", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	joinRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	r := response.(map[string]interface{})
	if _, ok := r["room_id"]; !ok {
		t.Error("Should respond with an error message if room id is not set.")
	} else if r["room_id"] != "Room id is required." {
		t.Error("Room id error message should be 'Room id is required.'")
	}
}

func TestJoinRoomHandlerShouldRespondWithErrorMessageIfUsernameNotSet(t *testing.T) {
	joinRoomHandler := joinRoomHandler()
	params := url.Values{}
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/join", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	joinRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	r := response.(map[string]interface{})
	if _, ok := r["username"]; !ok {
		t.Error("Should respond with an error message if username is not set.")
	} else if r["username"] != "Username is required." {
		t.Error("Username error message should be 'Username is required.'")
	}
}

func TestJoinRoomHandlerShouldRespondWithErrorMessageIfEmptyUsernameParameter(t *testing.T) {
	joinRoomHandler := joinRoomHandler()
	params := url.Values{}
	params.Add("username", "")
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/join", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	joinRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	r := response.(map[string]interface{})
	if _, ok := r["username"]; !ok {
		t.Error("Should respond with an error message if username is not set.")
	} else if r["username"] != "Username is required." {
		t.Error("Username error message should be 'Username is required.'")
	}
}

func TestJoinRoomHandlerShouldRespondWithErroMessageIfRequestMethodIsNotPOST(t *testing.T) {
	joinRoomHandler := joinRoomHandler()
	params := url.Values{}
	// GET, PUT, DELETE
	request, _ := http.NewRequest("GET", "localhost:8080/chat/room/join", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	joinRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	r := response.(map[string]interface{})
	if _, ok := r["message"]; !ok {
		t.Error("Should respond with an error message.")
	} else if r["message"] != "Method not allowed." {
		t.Error("Error message content should be 'Method not allowed.'")
	}
}

func TestJoinRoomHandlerResponseContentType(t *testing.T) {
	joinRoomHandler := joinRoomHandler()
	params := url.Values{}
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	joinRoomHandler.ServeHTTP(w, request)
	if w.Header().Get("Content-Type") != "application/json" {
		t.Error("Content-Type of the response should be 'application/json'.")
	}
}

func TestJoinRoomHandlerStatusCodeIfRoomNotFound(t *testing.T) {
	joinRoomHandler := joinRoomHandler()
	params := url.Values{}
	params.Add("room_id", "123123")
	params.Add("username", "mownier")
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	joinRoomHandler.ServeHTTP(w, request)
	if w.Code != http.StatusNotFound {
		t.Error("Status code should http.StatusNotFound.")
	}
}

func TestJoinRoomHandlerResponseMessageIfRoomNotFound(t *testing.T) {
	joinRoomHandler := joinRoomHandler()
	params := url.Values{}
	params.Add("room_id", "123123")
	params.Add("username", "mownier")
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	joinRoomHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	r := response.(map[string]interface{})
	if _, ok := r["message"]; !ok {
		t.Error("There's no message key in the response.")
	} else if r["message"] != "Room not found." {
		t.Error("Message should be 'Room not found.'")
	}
}
