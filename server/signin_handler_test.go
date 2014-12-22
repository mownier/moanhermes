package server

import (
	"testing"
	"net/url"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"bytes"
)

func TestSigninHandlerInvalidRequestMethod(t *testing.T) {
	signinHandler := signinHandler()
	params := url.Values{}
	request, _ := http.NewRequest("GET", "localhost:8080/chat/signin", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	signinHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code != http.StatusMethodNotAllowed {
		t.Error("Status code is not http.StatusMethodNotAllowed")
	} else if response == nil {
		t.Error("Should have a response.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["message"]; !ok {
			t.Error("There is no message key.")
		} else if r["message"] != "Method not allowed." {
			t.Error("Error message should be 'Method not allowed.'")
		}
	}
}

func TestSigninHandlerUsernameParameterNotSet(t *testing.T) {
	signinHandler := signinHandler()
	params := url.Values{}
	request, _ := http.NewRequest("POST", "localhost:8080/chat/signin", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	signinHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code != http.StatusBadRequest {
		t.Error("Status code is not http.StatusBadRequest")
	} else if response == nil {
		t.Error("Should have a response.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["username"]; !ok {
			t.Error("There is no message key.")
		} else if r["username"] != "Username is required." {
			t.Error("Error message should be 'Username is required.'")
		}
	}
}

func TestSigninHandlerEmptyUsernameParameter(t *testing.T) {
	signinHandler := signinHandler()
	params := url.Values{}
	params.Add("username", "")
	request, _ := http.NewRequest("POST", "localhost:8080/chat/signin", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	signinHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code != http.StatusBadRequest {
		t.Error("Status code is not http.StatusBadRequest")
	} else if response == nil {
		t.Error("Should have a response.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["username"]; !ok {
			t.Error("There is no message key.")
		} else if r["username"] != "Username is required." {
			t.Error("Error message should be 'Username is required.'")
		}
	}
}

func TestSigninHandlerUserDoesNotExist(t *testing.T) {
	signinHandler := signinHandler()
	params := url.Values{}
	params.Add("username", "yoyo")
	request, _ := http.NewRequest("POST", "localhost:8080/chat/signin", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	signinHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code != http.StatusNotFound {
		t.Error("Status code is not http.StatusNotFound")
	} else if response == nil {
		t.Error("Should have a response.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["message"]; !ok {
			t.Error("There is no message key.")
		} else if r["message"] != "User does not exist." {
			t.Error("Error message should be 'User does not exist.'")
		}
	}
}

func TestSigninHandlerSuccessfullySignedIn(t *testing.T) {
	var user *User = NewUser("yugi")
	users = append(users, user)
	signinHandler := signinHandler()
	params := url.Values{}
	params.Add("username", user.Username)
	request, _ := http.NewRequest("POST", "localhost:8080/chat/signin", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	signinHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	if w.Code != http.StatusOK {
		t.Error("Status code is not http.StatusOK")
	} else if response == nil {
		t.Error("Should have a response.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["message"]; !ok {
			t.Error("There is no message key.")
		} else if r["message"] != "Successfully signed in." {
			t.Error("Error message should be 'Successfully signed in.'")
		} else {
			var isUserOnline bool
			for i := 0; i < len(users); i++ {
				var u *User = users[i]
				if u.Username == user.Username {
					if u.Online {
						isUserOnline = true
					}
					break
				}
			}
			if !isUserOnline {
				t.Error("User not online.")
			}
		}
	}
}

















