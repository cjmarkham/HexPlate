package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func HandleValidationError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	errBody := map[string]string{
		"title":   "Validation Error",
		"message": message,
		"status":  strconv.Itoa(statusCode),
	}
	_ = json.NewEncoder(w).Encode(errBody)
}

func HandleResponseError(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	errBody := map[string]string{
		"title":   "Internal Error",
		"message": "Something went wrong",
		"status":  strconv.Itoa(http.StatusInternalServerError),
	}
	_ = json.NewEncoder(w).Encode(errBody)
}
