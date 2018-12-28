package main

import (
	"log"
	"net/http"

	"./api"
	"./commonutils"
	"./conf"
	"./data"
	"./routes"
	"github.com/gorilla/mux"
)

type WithCORS struct {
	r *mux.Router
}

func main() {
	log.Printf("Current folder %s", commonutils.CurrentFolder())
	configuration := conf.Get()
	db := data.NewSqliteDB("data.db")
	api := api.NewAPI(db)
	router := routes.NewRoutes(api)

	port := ":" + configuration.Port
	log.Printf("Server listening at: 127.0.0.1%s", port)
	http.ListenAndServe(port, &WithCORS{router})
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
