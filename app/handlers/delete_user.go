package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jhonatanlteodoro/verify_test/app/models"
	"gorm.io/gorm"
)

func DeleteUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vars = mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondInvalidID(w, vars["id"])
			return
		}

		db.Delete(&models.User{}, id)

		resp := map[string]string{
			"status": "ok",
		}
		respondWithJSON(w, http.StatusOK, resp)
	}
}
