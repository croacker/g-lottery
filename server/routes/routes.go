package routes

import (
	"log"
	"net/http"

	"../api"
	"../conf"
	"github.com/gorilla/mux"
	// "github.com/gorilla/handlers"
	// "github.com/urfave/negroni"
)

// NewRoutes builds the routes for the api
func NewRoutes(api *api.API) *mux.Router {
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

	// nominations
	// nomination := a.PathPrefix("/nomination").Subrouter()
	// apiRouter.HandleFunc("/nomination", api.NominationOptions).Methods("OPTIONS")
	apiRouter.HandleFunc("/nomination", api.NominationsAll).Methods("GET")
	apiRouter.HandleFunc("/nomination/{id}", api.GetNomination).Methods("GET")
	apiRouter.HandleFunc("/nomination", api.CreateNomination).Methods("POST")
	apiRouter.HandleFunc("/nomination/{id}", api.DeleteNomination).Methods("DELETE")
	apiRouter.HandleFunc("/nominationbycode/{id}", api.GetNominationByCode).Methods("GET")
	// nomination.HandleFunc("/login", api.UserLogin).Methods("POST")
	// nomination.Handle("/info", negroni.New(
	// 	negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
	// 	negroni.Wrap(http.HandlerFunc(api.UserInfo)),
	// ))

	// participant
	// participant := a.PathPrefix("/participant").Subrouter()
	apiRouter.HandleFunc("/participant", api.ParticipantsAll).Methods("GET")
	apiRouter.HandleFunc("/participant/{id}", api.GetParticipant).Methods("GET")
	apiRouter.HandleFunc("/participantbynomination/{id}", api.GetParticipantsByNominsation).Methods("GET")
	apiRouter.HandleFunc("/participant", api.CreateParticipant).Methods("POST")
	apiRouter.HandleFunc("/participant/{id}", api.DeleteParticipant).Methods("DELETE")
	apiRouter.HandleFunc("/participantbynomination/{id}", api.DeleteParticipantsByNominsation).Methods("DELETE")
	// participant.Handle("/protected/random", negroni.New(
	// 	negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
	// 	negroni.Wrap(http.HandlerFunc(api.SecretQuote)),
	// ))
	apiRouter.HandleFunc("/nominationresultbynomination", api.NominationsResultAll).Methods("GET")
	apiRouter.HandleFunc("/nominationresultbynomination/{id}", api.GetNominationResultByNominsation).Methods("GET")
	apiRouter.HandleFunc("/nominationresult/{id}", api.PlayANominsation).Methods("GET")
	apiRouter.HandleFunc("/nominationresult/{id}", api.DeletePlayANominsation).Methods("DELETE")
	// http.ListenAndServe(":3000", handlers.CORS()(apiRouter))

	return mux
}

//YourHandler
func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
