package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jhonatanlteodoro/verify_test/app/models"
	"github.com/jhonatanlteodoro/verify_test/app/sqlite_connector"
)

func TestBasicBehaviorOfCreateUserHandler(t *testing.T) {
	wait := 1
	retry := 0
	db, err := sqlite_connector.GetConnection(wait, retry)

	if err != nil {
		t.Error("fail connecting database")
	}
	models.RunMigrations(db)

	data := "{\"name\": \"test\", \"password\": \"some password\"}"
	bData := bytes.NewBuffer([]byte(data))

	request, request_err := http.NewRequest(http.MethodPost, "/users/", bData)
	if request_err != nil {
		t.Error(request_err)
	}
	request.Header.Set("Content-type", "application/json")

	responseRecorder := httptest.NewRecorder()
	router := mux.NewRouter()

	router.HandleFunc("/users/", CreateUser(db))
	router.ServeHTTP(responseRecorder, request)

	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected 200 as status code but got: %d", responseRecorder.Code)
	}

	contentType := responseRecorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected application/json content-type but got: %s", contentType)
	}
}
