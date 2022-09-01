package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error while marshaling response payload, ERROR: %v", err)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	payload := map[string]string{"error": message}

	respondWithJSON(w, code, payload)
}

func respondInvalidID(w http.ResponseWriter, id string) {
	log.Printf("Invalid ID. id: %s", id)
	respondWithError(w, http.StatusBadRequest, "invalid id")
}

func respondUserNotFound(w http.ResponseWriter) {
	respondWithError(w, http.StatusNotFound, "user not found")
}
