package api

import (
	"encoding/json"
	"net/http"
)

// ParticipantsAll -
func (api *API) ParticipantsAll(w http.ResponseWriter, req *http.Request) {
	all := api.participants.GetAll()
	allJSON, err := json.Marshal(all)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(allJSON)
}
