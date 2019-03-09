package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// NominationJSON
type NominationJSON struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// NominationOptions -
func (api *API) NominationOptions(responseWriter http.ResponseWriter, req *http.Request) {
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	responseWriter.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, YourOwnHeader")
}

// NominationsAll get all Nominations
func (api *API) NominationsAll(responseWriter http.ResponseWriter, req *http.Request) {
	all := api.nominations.GetAll()
	toJSON(responseWriter, all)
}

// GetNomination - get Nomination by id
func (api *API) GetNomination(responseWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	nomination := api.nominations.NominationByID(id)
	toJSON(responseWriter, nomination)
}

// GetNominationByCode - get Nomination by code
func (api *API) GetNominationByCode(responseWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["code"]
	nomination := api.nominations.NominationByCode(id)
	toJSON(responseWriter, nomination)
}

// CreateNomination - create Nomination
func (api *API) CreateNomination(responseWriter http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	jsondata := NominationJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Name == "" {
		http.Error(responseWriter, "Missing Name", http.StatusBadRequest)
		return
	}

	if api.nominations.HasNomination(jsondata.Name) {
		http.Error(responseWriter, "nomination already exists", http.StatusBadRequest)
		return
	}

	newNomination := api.nominations.CreateNomination(jsondata.Name, jsondata.Code)
	toJSON(responseWriter, newNomination)
}

// DeleteNomination -
func (api *API) DeleteNomination(responseWriter http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	nomination := api.nominations.DeleteNomination(id)
	toJSON(responseWriter, nomination)
}
