package controller

import (
	"dimitrisCBR/bookie-api/v2/v1/modeldel"
	u "dimitrisCBR/bookie-api/v2/v1/utiltil"
	"encoding/json"
	"net/http"
)

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &model.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := model.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
