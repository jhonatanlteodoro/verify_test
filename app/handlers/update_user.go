package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jhonatanlteodoro/verify_test/app/models"
	"gorm.io/gorm"
)

func UpdateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vars = mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondInvalidID(w, vars["id"])
			return
		}

		var user models.User
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "invalid payload")
			return
		}
		defer r.Body.Close()

		user.ID = uint(id)
		tx := db.Model(&user).Updates(&user)
		if !tx.Statement.Changed() {
			respondUserNotFound(w)
			return
		}

		resp := map[string]string{
			"status": "ok",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
