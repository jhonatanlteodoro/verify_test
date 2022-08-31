package handlers

import (
	"encoding/json"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"status": "Not Implemented",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
