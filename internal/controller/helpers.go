package controller

import (
	"encoding/json"
	"net/http"
)

func parseRequestBody[T any](r *http.Request) (T, error) {
	var dto T
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return dto, err
	}
	return dto, nil
}

func respondWithError(w http.ResponseWriter, status int, message error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": message.Error(),
		"status":  status,
		"success": false,
	})
}

func responseOk(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": message,
		"status":  status,
		"data":    data,
		"success": true,
	})
}
