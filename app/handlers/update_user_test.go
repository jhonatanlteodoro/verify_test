package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jhonatanlteodoro/verify_test/app/db"
	"github.com/jhonatanlteodoro/verify_test/app/models"
)

func TestBasicBehaviorOfUpdateUserHandler(t *testing.T) {
	wait := 1
	retry := 0
	db, err := db.GetConnection(wait, retry)
	if err != nil {
		t.Error("fail connecting database")
	}
	models.RunMigrations(db)

	data := "{\"name\": \"teste\", \"email\": \"somseemil@gmail.com\"}"
	bData := bytes.NewBuffer([]byte(data))

	request, request_err := http.NewRequest(http.MethodPut, "/users/1", bData)
	if request_err != nil {
		t.Error(request_err)
	}

	responseRecorder := httptest.NewRecorder()
	router := mux.NewRouter()

	router.HandleFunc("/users/{id}", UpdateUser(db))
	router.ServeHTTP(responseRecorder, request)

	fmt.Println(responseRecorder.Body)
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected 200 as status code but got: %d", responseRecorder.Code)
	}

	contentType := responseRecorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected application/json content-type but got: %s", contentType)
	}
}
