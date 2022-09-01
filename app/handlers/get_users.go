package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jhonatanlteodoro/verify_test/app/models"
	"gorm.io/gorm"
)

func GetAll(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []models.APIResponseUser
		db.Model(&models.User{}).Find(&users)

		resp := map[string]interface{}{
			"status": "ok",
			"data":   users,
		}
		respondWithJSON(w, http.StatusOK, resp)
	}
}

func GetUserById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vars = mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondInvalidID(w, vars["id"])
			return
		}

		user := models.APIResponseUser{}
		db.Model(&models.User{}).Where("id = ?", id).First(&user)

		if user.ID == 0 {
			respondUserNotFound(w)
			return
		}

		resp := map[string]interface{}{
			"status": "ok",
			"data":   user,
		}
		respondWithJSON(w, http.StatusOK, resp)
	}
}
