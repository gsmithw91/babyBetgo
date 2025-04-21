// json.go
package utils

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, payload JSONResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)

}

func ErrorResponse(w http.ResponseWriter, message string, status int) {
	WriteJSON(w, status, JSONResponse{
		Status:  "error",
		Message: message,
	})

}
