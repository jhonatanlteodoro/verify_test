package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jhonatanlteodoro/verify_test/app/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Printf("Error decoding payload. Error: %v\n", err)
			respondWithError(w, http.StatusInternalServerError, "Error")
			return
		}

		db.Create(&user)

		response := map[string]interface{}{
			"status": "OK",
			"id":     user.ID,
		}

		respondWithJSON(w, http.StatusOK, response)
	}
}
