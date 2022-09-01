package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestBasicBehaviorOfGetUserByIdHandler(t *testing.T) {

	request, err := http.NewRequest(http.MethodGet, "/users/1", nil)
	if err != nil {
		t.Error(err)
	}

	responseRecorder := httptest.NewRecorder()
	router := mux.NewRouter()

	router.HandleFunc("/users/{id}", GetUserById)
	router.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected 200 as status code but got: %d", responseRecorder.Code)
	}

	contentType := responseRecorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected application/json content-type but got: %s", contentType)
	}
}

func TestBasicBehaviorOfGetAllHandler(t *testing.T) {

	request, err := http.NewRequest(http.MethodGet, "/users/", nil)
	if err != nil {
		t.Error(err)
	}

	responseRecorder := httptest.NewRecorder()
	router := mux.NewRouter()

	router.HandleFunc("/users/", GetAll)
	router.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected 200 as status code but got: %d", responseRecorder.Code)
	}

	contentType := responseRecorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected application/json content-type but got: %s", contentType)
	}
}
