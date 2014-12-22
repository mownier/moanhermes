package server

import(
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
	"encoding/json"
	"bytes"
)

func TestRegisterHandlerInvalidRequestMethod(t *testing.T) {
	registerHandler := registerHandler()
	params := url.Values{}
	request, _ := http.NewRequest("GET", "localhost:8080/chat/register", bytes.NewBufferString(params.Encode()))
	w := httptest.NewRecorder()
	registerHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code != http.StatusMethodNotAllowed {
		t.Error("Status code is not http.StatusMethodNotAllowed.")
	} else if response == nil {
		t.Error("Should have a response.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["message"]; !ok {
			t.Error("There is no message key 'message'.")
		} else if r["message"] != "Method not allowed." {
			t.Error("Error message should be 'Method not allowed.'")
		}
	}
}

func TestRegisterHandlerEmptyUsernameParameter(t *testing.T) {
	registerHandler := registerHandler()
	params := url.Values{}
	params.Add("username", "")
	request, _ := http.NewRequest("POST", "localhost:8080/chat/register", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-urlencoded; param=value")
	w := httptest.NewRecorder()
	registerHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code != http.StatusBadRequest {
		t.Error("Status code is not http.StatusBadRequest.")
	} else if response == nil {
		t.Error("Should have a response.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["username"]; !ok {
			t.Error("There is no message key 'username'.")
		} else if r["username"] != "Username is required." {
			t.Error("Error message should be 'Username is required.'")
		}
	}
}

func TestRegisterHandlerUsernameParameterIsNotSet(t *testing.T) {
	registerHandler := registerHandler()
	params := url.Values{}
	request, _ := http.NewRequest("POST", "localhost:8080/chat/register", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-urlencoded; param=value")
	w := httptest.NewRecorder()
	registerHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code != http.StatusBadRequest {
		t.Error("Status code is not http.StatusBadRequest.")
	} else if response == nil {
		t.Error("Should have a response.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["username"]; !ok {
			t.Error("There is no message key 'username'.")
		} else if r["username"] != "Username is required." {
			t.Error("Error message should be 'Username is required.'")
		}
	}
}

func TestRegisterHandlerUsernameAlreadyExists(t *testing.T) {
	var newUser *User = NewUser("mownier")
	users = append(users, newUser)

	registerHandler := registerHandler()
	params := url.Values{}
	params.Add("username", "mownier")
	request, _ := http.NewRequest("POST", "localhost:8080/chat/register", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	registerHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code != http.StatusBadRequest {
		t.Error("Status code is not http.StatusBadRequest.")
	} else if response == nil {
		t.Error("Should have a response.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["message"]; !ok {
			t.Error("There is no message key 'message'.")
		} else if r["message"] != "Username already exists." {
			t.Error("Error message should be 'Username already exists.'")
		}
	}
}

func TestRegisterHandlerSuccessfullyRegistered(t *testing.T) {
	var numberOfUsersBeforeSuccessfulRegistration int = len(users)

	registerHandler := registerHandler()
	params := url.Values{}
	params.Add("username", "iammownier")
	request, _ := http.NewRequest("POST", "localhost:8080/chat/register", bytes.NewBufferString(params.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	registerHandler.ServeHTTP(w, request)
	var response interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if w.Code != http.StatusOK {
		t.Error("Status code is not http.StatusOK.")
	} else if response == nil {
		t.Error("Should have a response.")
	} else {
		r := response.(map[string]interface{})
		if _, ok := r["message"]; !ok {
			t.Error("There is no message key 'message'.")
		} else if r["message"] != "Successfully registered." {
			t.Error("Error message should be 'Successfully registered.'")
		} else {
			var numberOfUsersAfterSuccessfulRegistration int = len(users)
			if numberOfUsersAfterSuccessfulRegistration != numberOfUsersBeforeSuccessfulRegistration + 1 {
				t.Error("New user is not added into the array 'users'.")
			}
		}
	}
}






























