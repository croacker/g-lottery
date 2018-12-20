package routes

import (
	"net/http"

	"../api"
	"github.com/gorilla/mux"
	// "github.com/urfave/negroni"
)

// NewRoutes builds the routes for the api
func NewRoutes(api *api.API) *mux.Router {

	mux := mux.NewRouter()

	// client static files
	mux.Handle("/", http.FileServer(http.Dir("../client/dist/"))).Methods("GET")
	mux.PathPrefix("/js").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("../client/dist/js/"))))
	mux.PathPrefix("/img").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("../client/dist/img/"))))
	mux.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("../client/dist/css/"))))

	// api
	apiRouter := mux.PathPrefix("/api").Subrouter()

	// nominations
	// nomination := a.PathPrefix("/nomination").Subrouter()
	apiRouter.HandleFunc("/nomination", api.NominationsAll).Methods("GET")
	apiRouter.HandleFunc("/nomination/{id}", api.GetNomination).Methods("GET")
	apiRouter.HandleFunc("/nomination", api.CreateNomination).Methods("POST")
	apiRouter.HandleFunc("/nomination/{id}", api.DeleteNomination).Methods("DELETE")
	// nomination.HandleFunc("/login", api.UserLogin).Methods("POST")
	// nomination.Handle("/info", negroni.New(
	// 	negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
	// 	negroni.Wrap(http.HandlerFunc(api.UserInfo)),
	// ))

	// participant
	// participant := a.PathPrefix("/participant").Subrouter()
	apiRouter.HandleFunc("/participant", api.ParticipantsAll).Methods("GET")
	apiRouter.HandleFunc("/participant/{id}", api.GetParticipant).Methods("GET")
	apiRouter.HandleFunc("/participant", api.CreateParticipant).Methods("POST")
	apiRouter.HandleFunc("/participant/{id}", api.DeleteParticipant).Methods("DELETE")
	// participant.Handle("/protected/random", negroni.New(
	// 	negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
	// 	negroni.Wrap(http.HandlerFunc(api.SecretQuote)),
	// ))

	return mux
}
