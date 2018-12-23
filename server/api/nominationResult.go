package api

import (
	"net/http"

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
func (api *API) DeleteNominationResult(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	nomination := api.nominations.DeleteNomination(id)
	toJSON(w, nomination)
}
