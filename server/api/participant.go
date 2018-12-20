package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// ParticipantJSON
type ParticipantJSON struct {
	Surname      string `json:"surname"`
	Name         string `json:"name"`
	NominationId string `json:"nominationId"`
}

// ParticipantsAll -
func (api *API) ParticipantsAll(w http.ResponseWriter, req *http.Request) {
	all := api.participants.GetAll()
	toJSON(w, all)
}

// GetParticipant -
func (api *API) GetParticipant(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	nomination := api.participants.ParticipantByID(id)
	toJSON(w, nomination)
}

// CreateParticipant -
func (api *API) CreateParticipant(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	jsondata := ParticipantJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Name == "" || jsondata.Surname == "" || jsondata.NominationId == "" {
		http.Error(w, "Missing property", http.StatusBadRequest)
		return
	}

	if api.participants.HasParticipant(jsondata.Name) {
		http.Error(w, "nomination already exists", http.StatusBadRequest)
		return
	}

	nomination := api.nominations.NominationByID(jsondata.NominationId)
	newParticipant := api.participants.CreateParticipant(jsondata.Surname, jsondata.Name, nomination)
	toJSON(w, newParticipant)
}

// DeleteParticipant -
func (api *API) DeleteParticipant(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	nomination := api.participants.DeleteParticipant(id)
	toJSON(w, nomination)
}
