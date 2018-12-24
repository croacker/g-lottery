package api

import (
	"math/rand"
	"net/http"
	"time"

	"../data"
	"github.com/gorilla/mux"
)

// NominationsAll -
func (api *API) NominationsResultAll(w http.ResponseWriter, req *http.Request) {
	all := api.nominationResult.GetAll()
	toJSON(w, all)
}

// GetNomination -
func (api *API) GetNominationResult(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	nomination := api.nominationResult.ByID(id)
	toJSON(w, nomination)
}

// GetParticipant -
func (api *API) GetNominationResultByNominsation(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	nominationID := params["id"]
	nomination := api.nominations.NominationByID(nominationID)
	nominationResult := api.nominationResult.ByNominationID(nomination)
	toJSON(w, nominationResult)
}

// DeleteNominationResult -
func (api *API) PlayANominsation(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	nomination := api.nominations.NominationByID(id)
	participants := api.participants.ByNominationID(nomination)
	var participantsToPlay []data.Participant
	for _, participant := range participants {
		for i := 0; i < participant.Chance; i++ {
			participantsToPlay = append(participantsToPlay, participant)
		}
	}
	count := len(participantsToPlay)
	winnerIdx := random(0, count-1)
	winner := participantsToPlay[winnerIdx]
	toJSON(w, winner)
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
