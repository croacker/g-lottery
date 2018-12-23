package api

import (
	"encoding/json"
	"net/http"

	"../data"
)

// API -
type API struct {
	nominations  *data.NominationManager
	participants *data.ParticipantManager
}

// NewAPI -
func NewAPI(db *data.DB) *API {

	nominationmgr, _ := data.NewNominationManager(db)
	participantmgr, _ := data.NewParticipantManager(db)

	nominationmgr.Predefined()
	return &API{
		nominations:  nominationmgr,
		participants: participantmgr,
	}
}

func toJSON(w http.ResponseWriter, v interface{}) {
	jsondata, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsondata)
}
