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
	nominationResult := api.getNominationResultByNomination(nominationID)
	toJSON(w, nominationResult)
}

// PlayANominsation -
func (api *API) PlayANominsation(w http.ResponseWriter, req *http.Request) {
	var winner data.Participant

	params := mux.Vars(req)
	nominationID := params["id"]
	nomination := api.nominations.NominationByID(nominationID)
	nominationResult := api.getNominationResultByNomination(nominationID)
	if nominationResult.ID == 0 {
		participants := api.participants.ByNominationID(nomination)
		var participantsToPlay []data.Participant
		for _, participant := range participants {
			for i := 0; i < participant.Chance; i++ {
				participantsToPlay = append(participantsToPlay, participant)
			}
		}
		count := len(participantsToPlay)
		winnerIdx := random(0, count-1)
		winner = participantsToPlay[winnerIdx]
		api.nominationResult.CreateNominationResult(nomination, &winner)
	} else {
		winner = nominationResult.Participant
	}
	toJSON(w, winner)
}

// PlayANominsation -
func (api *API) DeletePlayANominsation(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	nominationID := params["id"]
	nominationResult := api.getNominationResultByNomination(nominationID)
	if nominationResult.ID != 0 {
		api.nominationResult.Delete(nominationResult)
		toJSON(w, nominationResult)
	} else {
		var emptyI interface{}
		toJSON(w, emptyI)
	}
}

func (api *API) getNominationResultByNomination(nominationID string) *data.NominationResult {
	nomination := api.nominations.NominationByID(nominationID)
	nominationResult := api.nominationResult.ByNominationID(nomination)
	return nominationResult
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
