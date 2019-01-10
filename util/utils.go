package util

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(rw http.ResponseWriter, data map[string] interface{})  {
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}