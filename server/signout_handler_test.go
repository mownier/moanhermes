package server

import (
	"testing"
	"net/url"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"bytes"
)

func TestSignoutHandlerInvalidRequestMethod(t *testing.T) {
	signoutHandler := signoutHandler()
	request, _ := http.NewRequest("DELETE", "http://localhost:8080/chat/signout", nil)
	w := httptest.NewRecorder()
	signoutHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	if w.Code != http.StatusMethodNotAllowed {
		t.Error("Status code is not http.StatusMethodNotAllowed.")
	} else if response == nil {
		t.Error("Response is nil.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["message"]; !ok {
			t.Error("There is no message key.")
		} else if r["message"] != "Method not allowed." {
			t.Error("Response message is not 'Method not allowed.'")
		}
	}
}

func TestSignoutHandlerEmptyUseranme(t *testing.T) {
	signoutHandler := signoutHandler()
	params := url.Values{}
	request, _ := http.NewRequest("POST", "http://localhost:8080/chat/signout", bytes.NewBufferString(params.Encode()))
	w := httptest.NewRecorder()
	signoutHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	if w.Code != http.StatusBadRequest {
		t.Error("Status code is not http.StatusBadRequest.")
	} else if response == nil {
		t.Error("Response is nil.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["username"]; !ok {
			t.Error("There is no message key.")
		} else if r["username"] != "Username is required." {
			t.Error("Error message is not 'Username is required.'")
		}
	}
}

func TestSignoutHandlerUsernameDoesNotExist(t *testing.T) {
	signoutHandler := signoutHandler()
	params := url.Values{}
	params.Add("username", "adfadf")
	request, _ := http.NewRequest("POST", "http://localhost:8080/chat/signout/", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	signoutHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	if w.Code != http.StatusNotFound {
		t.Error("Status code is not http.StatusNotFound.")
	} else if response == nil {
		t.Error("Response is nil.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["message"]; !ok {
			t.Error("There is no message key.")
		} else if r["message"] != "Username not found." {
			t.Error("Error message is not 'Username not found.'")
		}
	}
}

func TestSignoutHandlerSetOnlineToFalse(t *testing.T) {
	var user *User = NewUser("edgar")
	user.Online = true
	users = append(users, user)

	signoutHandler := signoutHandler()
	params := url.Values{}
	params.Add("username", user.Username)
	request, _ := http.NewRequest("POST", "http://localhost:8080/chat/signout", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	signoutHandler.ServeHTTP(w, request)
	
	for i := 0; i < len(users); i++ {
		if users[i].Username == user.Username {
			if users[i].Online == true {
				t.Error("User's online status is not set to false.")
			}
			break
		}
	}
}

func TestSignoutHandlerSuccessResponse(t *testing.T) {
	var user *User = NewUser("edgar")
	user.Online = true
	users = append(users, user)

	signoutHandler := signoutHandler()
	params := url.Values{}
	params.Add("username", user.Username)
	request, _ := http.NewRequest("POST", "http://localhost:8080/chat/signout", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	signoutHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code != http.StatusOK {
		t.Error("Status code is not http.StatusOK")
	} else if response == nil {
		t.Error("Response is nil.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["message"]; !ok {
			t.Error("There is no message key.")
		} else if r["message"] != "Signed out successfully." {
			t.Error("Response message is not 'Signed out successfully.'")
		}
	}
}






























