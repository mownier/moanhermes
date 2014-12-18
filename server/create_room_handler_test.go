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
	// Setting the request header content type
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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
	} else if r["username"] != "Username is required." {
		t.Error("Username error message should be 'Username is required.'")
	}
}

func TestCreateRoomHandlerUsernameErrorMessageNotSetAsParameter(t *testing.T) {
	// Creating the handler function for the create room handler
	createRoomHandler := createRoomHandler()
	// Creating parameter values for the request
	params := url.Values{}
	// Creating an http request
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/create", bytes.NewBufferString(params.Encode()))
	// Setting the request header content type
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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
	} else if r["username"] != "Username is required." {
		t.Error("Username error message should be 'Username is required.'")
	}
}

func TestCreateRoomHandlerUsernameErrorStatusCode(t *testing.T) {
	// Creating the handler function for the create room handler
	createRoomHandler := createRoomHandler()
	// Creating parameter values for the request
	params := url.Values{}
	// Creating an http request
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/create", bytes.NewBufferString(params.Encode()))
	// Setting the request header content type
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	// Creating a response recorder
	w := httptest.NewRecorder()
	// Serving the http request and the response recorder
	createRoomHandler.ServeHTTP(w, request)
	// Checking the response status code
	if w.Code != http.StatusBadRequest {
		t.Error("Username error reponse status code is not http.StatusBadRequest")
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
	// Setting the request header content type
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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
	} else if r["room_name"] != "Room name is required." {
		t.Error("Room name error message should be 'Room name is required.'")
	}
}

func TestCreateRoomHandlerRoomNameErrorMessageNotSetAsParameter(t *testing.T) {
	// Creating the handler function for the create room handler
	createRoomHandler := createRoomHandler()
	// Creating parameter values for the request
	params := url.Values{}
	// Creating an http request
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/create", bytes.NewBufferString(params.Encode()))
	// Setting the request header content type
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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
	} else if r["room_name"] != "Room name is required." {
		t.Error("Room name error message should be 'Room name is required.'")
	}
}

func TestCreateRoomHandlerRoomNameErrorStatusCode(t *testing.T) {
	// Creating the handler function for the create room handler
	createRoomHandler := createRoomHandler()
	// Creating parameter values for the request
	params := url.Values{}
	// Creating an http request
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/create", bytes.NewBufferString(params.Encode()))
	// Setting the request header content type
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	// Creating a response recorder
	w := httptest.NewRecorder()
	// Serving the http request and the response recorder
	createRoomHandler.ServeHTTP(w, request)
	// Checking the status code
	if w.Code != http.StatusBadRequest {
		t.Error("Room name error response status code is not http.StatusBadRequest.")
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
	// Setting the request header content type
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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
	} else if r["message"] != "Method not allowed." {
		t.Error("Invalid method error message should be 'Method not allowed.'")
	}
}

func TestCreateRoomHandlerInvalidMethodStatusCode(t *testing.T) {
	// Creating the handler function for the create rrom handler
	createRoomHandler := createRoomHandler()
	// Creating parameter values for the request
	params := url.Values{}
	// Creating an http request
	// GET, PUT, DELETE 
	request, _ := http.NewRequest("GET", "localhost:8080/chat/room/create", bytes.NewBufferString(params.Encode()))
	// Setting the request header content type
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	// Creating a response recorder
	w := httptest.NewRecorder()
	// Serving the http request and the response recorder
	createRoomHandler.ServeHTTP(w, request)
	// Checking the method not allowed status
	if (w.Code != http.StatusMethodNotAllowed) {
		t.Error("Status code is not http.StatusMethodNotAllowed.")
	}
}

func TestCreateRoomHandlerResponseJsonContentType(t *testing.T) {
	// Creating the handler function for the create room handler
	createRoomHandler := createRoomHandler()
	// Creating parameter values for the request
	params := url.Values{}
	// Creating an http request
	// POST, GET, PUT, DELETE
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/create", bytes.NewBufferString(params.Encode()))
	// Setting the request header content type
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	// Creating a response recorder
	w := httptest.NewRecorder()
	// Serving the http request and the response recorder
	createRoomHandler.ServeHTTP(w, request)
	// Checking the header's content type
	if w.Header().Get("Content-Type") != "application/json" {
		t.Error("Content-Type of the response's header is not 'application/json'.")
	}
}

func TestCreateRoomHandlerResponseSuccessfulMessage(t *testing.T) {
	// Creating the handler function for the create room handler
	createRoomHandler := createRoomHandler()
	// Creating parameter values for the request
	params := url.Values{}
	// Adding a username parameter
	params.Add("username", "mownier")
	// Adding a room name parameter
	params.Add("room_name", "room123")
	// Creating an http request
	request, _ := http.NewRequest("POST", "localhost:8080/chat/room/create", bytes.NewBufferString(params.Encode()))
	// Setting the request header content type
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
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
	if _, ok := r["message"]; !ok {
		t.Error("There's no sucessful message.")
	} else if r["message"] != "Successfully created a room." {
		t.Error("Successful message should be 'Successfully created a room.'")
	}
}
