package api

import (
	"encoding/json"
	"net/http"

	"../data"
)

// API -
type API struct {
	nominations      *data.NominationManager
	nominationResult *data.NominationResultManager
	participants     *data.ParticipantManager
}

// NewAPI -
func NewAPI(db *data.DB) *API {

	nominationMgr, _ := data.NewNominationManager(db)
	nominationResultMgr, _ := data.NewNominationResultManager(db)
	participantMgr, _ := data.NewParticipantManager(db)

	nominationMgr.Predefined()
	return &API{
		nominations:      nominationMgr,
		nominationResult: nominationResultMgr,
		participants:     participantMgr,
	}
}

func toJSON(reposnseWriter http.ResponseWriter, v interface{}) {
	jsondata, err := json.Marshal(v)
	if err != nil {
		http.Error(reposnseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	reposnseWriter.Header().Set("Content-Type", "application/json")
	reposnseWriter.Write(jsondata)
}
