package routes

import (
	"log"
	"net/http"

	"../api"
	"../conf"
	"github.com/gorilla/mux"
)

type WithCORS struct {
	r *mux.Router
}

//NewCorsRoutes builds the CORS routes for the api
func NewCorsRoutes(api *api.API) *WithCORS {
	return &WithCORS{newRoutes(api)}
}

// NewRoutes builds the routes for the api
func newRoutes(api *api.API) *mux.Router {
	configuration := conf.Get()
	mux := mux.NewRouter()

	// client static files
	clientFolder := configuration.ClientFolder
	log.Printf("Clietn folder %s", clientFolder)
	mux.Handle("/", http.FileServer(http.Dir(clientFolder+"/"))).Methods("GET")
	mux.PathPrefix("/js").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(clientFolder+"/js/"))))
	mux.PathPrefix("/img").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir(clientFolder+"/img/"))))
	mux.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(clientFolder+"/css/"))))
	mux.PathPrefix("/font").Handler(http.StripPrefix("/font/", http.FileServer(http.Dir(clientFolder+"/font/"))))

	// api
	mux.HandleFunc("/api", YourHandler)
	apiRouter := mux.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/nomination", api.NominationsAll).Methods("GET")
	apiRouter.HandleFunc("/nomination/{id}", api.GetNomination).Methods("GET")
	apiRouter.HandleFunc("/nomination/code/{code}", api.GetNominationByCode).Methods("GET")
	apiRouter.HandleFunc("/nomination", api.CreateNomination).Methods("POST")
	apiRouter.HandleFunc("/nomination/{id}", api.DeleteNomination).Methods("DELETE")

	apiRouter.HandleFunc("/participant", api.ParticipantsAll).Methods("GET")
	apiRouter.HandleFunc("/participant/{id}", api.GetParticipant).Methods("GET")
	apiRouter.HandleFunc("/participant/nomination/{id}", api.GetParticipantsByNominsation).Methods("GET")
	apiRouter.HandleFunc("/participant", api.CreateParticipant).Methods("POST")
	apiRouter.HandleFunc("/participant/{id}", api.DeleteParticipant).Methods("DELETE")
	apiRouter.HandleFunc("/participant/nomination/{id}", api.DeleteParticipantsByNominsation).Methods("DELETE")

	apiRouter.HandleFunc("/nominationresult/nomination", api.NominationsResultAll).Methods("GET")
	apiRouter.HandleFunc("/nominationresult/nomination/{id}", api.GetNominationResultByNominsation).Methods("GET")
	apiRouter.HandleFunc("/nominationresult/{id}", api.PlayANominsation).Methods("GET")
	apiRouter.HandleFunc("/nominationresult/{id}", api.DeletePlayANominsation).Methods("DELETE")

	return mux
}

//YourHandler
func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func (s *WithCORS) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s\n", req.RemoteAddr, req.Method, req.URL)
	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
		res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		res.Header().Set("Access-Control-Allow-Headers", "*")
	}

	if req.Method == "OPTIONS" {
		return
	}
	s.r.ServeHTTP(res, req)
}
