package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Bernar11296/rest/models"
	"github.com/Bernar11296/rest/utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := account.Create()
	utils.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	utils.Respond(w, resp)
}
