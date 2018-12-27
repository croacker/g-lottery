package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ParticipantJSON
type ParticipantJSON struct {
	Surname      string `json:"Surname"`
	Name         string `json:"Name"`
	Department   string `json:"Department"`
	Chance       int    `json:"Chance"`
	NominationID int    `json:"NominationID"`
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
	participant := api.participants.ParticipantByID(id)
	toJSON(w, participant)
}

// GetParticipant -
func (api *API) GetParticipantsByNominsation(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	nominationID := params["id"]
	nomination := api.nominations.NominationByID(nominationID)
	participants := api.participants.ByNominationID(nomination)
	toJSON(w, participants)
}

// GetParticipant -
func (api *API) DeleteParticipantsByNominsation(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	nominationID := params["id"]
	nomination := api.nominations.NominationByID(nominationID)
	participants := api.participants.ByNominationID(nomination)
	for _, participant := range participants {
		fmt.Println(participant)
		api.participants.Delete(&participant)
	}
	toJSON(w, participants)
}

// CreateParticipant -
func (api *API) CreateParticipant(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	jsondatas := []ParticipantJSON{}
	err := decoder.Decode(&jsondatas)

	for _, jsondata := range jsondatas {
		if err != nil || (jsondata.Name == "" && jsondata.Surname == "") || jsondata.NominationID == 0 {
			http.Error(w, "Missing property", http.StatusBadRequest)
			return
		}

		// existsParticipant := api.participants.FindParticipant(jsondata.Surname)
		// if existsParticipant.ID != 0 && jsondata.NominationID !=  {
		// 	http.Error(w, "nomination already exists", http.StatusBadRequest)
		// 	return
		// }

		nomination := api.nominations.NominationByID(strconv.Itoa(jsondata.NominationID))
		newParticipant := api.participants.CreateParticipant(jsondata.Surname, jsondata.Name, jsondata.Department, jsondata.Chance, nomination)
		fmt.Println("new Participant ", newParticipant)
	}
	toJSON(w, "")
}

// DeleteParticipant -
func (api *API) DeleteParticipant(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	nomination := api.participants.DeleteParticipant(id)
	toJSON(w, nomination)
}
