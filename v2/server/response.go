package server

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, code int, message string) {
	Json(w, code, map[string]string{"error": message})
}

func Json(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func JsonWithToken(w http.ResponseWriter, code int, payload interface{}, token string) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", token)
	w.WriteHeader(code)
	w.Write(response)
}
