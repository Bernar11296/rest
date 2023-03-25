package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bernar11296/rest/models"
	"github.com/Bernar11296/rest/utils"
	"github.com/gorilla/mux"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	contact.UserId = user
	resp := contact.Create()
	utils.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "There was an error in your request"))
		return
	}
	data := models.GetContacts(uint(id))
	resp := utils.Message(true, "succes")
	resp["data"] = data
	utils.Respond(w, resp)
}
