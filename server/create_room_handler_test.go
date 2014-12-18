package server

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
	"bytes"
	"encoding/json"
)

func TestCreateRoomHandlerUsernameErrorMessageWithEmptyParamValue(t *testing.T) {
	// Creating the handler function for the create room handler
	createRoomHandler := createRoomHandler()
	// Creating parameter values for the request
	params := url.Values{}
	// Adding a parameter
	params.Add("username", "")
	// Creating an http request
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/create", bytes.NewBufferString(params.Encode()))
	// Creating a response recorder
	w := httptest.NewRecorder()
	// Serving the http request and the response recorder
	createRoomHandler.ServeHTTP(w, request)
	// Declaring a container for the response
	var response interface{}
	// Converting json to response container
	json.Unmarshal(w.Body.Bytes(), &response)
	// Typecasting the response container to map
	r := response.(map[string]interface{})
	// Checking if the username key exist in the response
	if _, ok := r["username"]; !ok {
		t.Error("There's no username error message.")
	}
}

func TestCreateRoomHandlerUsernameErrorMessageNotSetAsParameter(t *testing.T) {
	// Creating the handler function for the create room handler
	createRoomHandler := createRoomHandler()
	// Creating parameter values for the request
	params := url.Values{}
	// Creating an http request
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/create", bytes.NewBufferString(params.Encode()))
	// Creating a response recorder
	w := httptest.NewRecorder()
	// Serving the http request and the response recorder
	createRoomHandler.ServeHTTP(w, request)
	// Declaring a container for the response
	var response interface{}
	// Converting json to response container
	json.Unmarshal(w.Body.Bytes(), &response)
	// Typecasting the response container to map
	r := response.(map[string]interface{})
	// Checking if the username key exist in the response
	if _, ok := r["username"]; !ok {
		t.Error("There's no username error message.")
	}
}

func TestCreateRoomHandlerRoomNameErrorMessageWithEmptyParamValue(t *testing.T) {
	// Creating the handler function for the create room handler
	createRoomHandler := createRoomHandler()
	// Creating parameter values for the request
	params := url.Values{}
	// Adding a parameter
	params.Add("username", "")
	// Creating an http request
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/create", bytes.NewBufferString(params.Encode()))
	// Creating a response recorder
	w := httptest.NewRecorder()
	// Serving the http request and the response recorder
	createRoomHandler.ServeHTTP(w, request)
	// Declaring a container for the response
	var response interface{}
	// Converting json to response container
	json.Unmarshal(w.Body.Bytes(), &response)
	// Typecasting the response container to map
	r := response.(map[string]interface{})
	// Checking if the room_name key exist in the response
	if _, ok := r["room_name"]; !ok {
		t.Error("There's no room_name error message.")
	}
}

func TestCreateRoomHandlerRoomNameErrorMessageNotSetAsParameter(t *testing.T) {
	// Creating the handler function for the create room handler
	createRoomHandler := createRoomHandler()
	// Creating parameter values for the request
	params := url.Values{}
	// Creating an http request
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/create", bytes.NewBufferString(params.Encode()))
	// Creating a response recorder
	w := httptest.NewRecorder()
	// Serving the http request and the response recorder
	createRoomHandler.ServeHTTP(w, request)
	// Declaring a container for the response
	var response interface{}
	// Converting json to response container
	json.Unmarshal(w.Body.Bytes(), &response)
	// Typecasting the response container to map
	r := response.(map[string]interface{})
	// Checking if the room_name key exist in the response
	if _, ok := r["room_name"]; !ok {
		t.Error("There's no room name error message.")
	}
}

func TestCreateRoomHandlerInvalidMethodResponseMessage(t *testing.T) {
	// Creating the handler function for the create rrom handler
	createRoomHandler := createRoomHandler()
	// Creating parameter values for the request
	params := url.Values{}
	// Creating an http request
	// GET, PUT, DELETE 
	request, _ := http.NewRequest("GET", "localhost:8080/chat/room/create", bytes.NewBufferString(params.Encode()))
	// Creating a response recorder
	w := httptest.NewRecorder()
	// Serving the http request and the response recorder
	createRoomHandler.ServeHTTP(w, request)
	// Declaring a container for the response
	var response interface{}
	// Converting json to response container
	json.Unmarshal(w.Body.Bytes(), &response)
	// Typecasting the response container to map
	r := response.(map[string]interface{})
	// Checking if message key exist in the response
	if _, ok := r["message"]; !ok {
		t.Error("There's no invalid method reponse message.")
	}
}
