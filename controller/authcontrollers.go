package controller

import (
	"net/http"
	u "dimitrisCBR/bookie-api/util"
	"dimitrisCBR/bookie-api/model"
	"encoding/json"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &model.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Create account
	u.Respond(w, resp)
}

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
