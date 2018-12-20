package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// NominationJSON - json data expected for login/signup
type NominationJSON struct {
	Name string `json:"name"`
}

// NominationsAll -
func (api *API) NominationsAll(w http.ResponseWriter, req *http.Request) {
	all := api.nominations.GetAll()
	// jsondata, err := json.Marshal(all)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(jsondata)
	toJSON(w, all)
}

// GetNomination -
func (api *API) GetNomination(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	nomination := api.nominations.NominationByID(id)
	// jsondata, err := json.Marshal(nomination)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(jsondata)
	toJSON(w, nomination)
}

// CreateNomination -
func (api *API) CreateNomination(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	jsondata := NominationJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Name == "" {
		http.Error(w, "Missing Name", http.StatusBadRequest)
		return
	}

	if api.nominations.HasNomination(jsondata.Name) {
		http.Error(w, "nomination already exists", http.StatusBadRequest)
		return
	}

	newNomination := api.nominations.CreateNomination(jsondata.Name)
	toJSON(w, newNomination)
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(jsondata)
}

// DeleteNomination -
func (api *API) DeleteNomination(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	nomination := api.nominations.DeleteNomination(id)

	toJSON(w, nomination)
}
