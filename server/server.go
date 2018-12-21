package main

import (
	"net/http"

	"./api"
	"./data"
	"./routes"
	"github.com/gorilla/mux"
)

// WithCORS
type WithCORS struct {
	r *mux.Router
}

func main() {
	db := data.NewSqliteDB("data.db")
	api := api.NewAPI(db)
	routes := routes.NewRoutes(api)
	http.ListenAndServe(":3000", &WithCORS{routes})
}

func (s *WithCORS) ServeHTTP(res http.ResponseWriter, req *http.Request) {
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
