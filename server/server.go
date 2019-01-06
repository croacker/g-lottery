package main

import (
	"log"
	"net/http"

	"./api"
	"./commonutils"
	"./conf"
	"./data"
	"./routes"
)

func main() {
	log.Printf("Current folder %s", commonutils.CurrentFolder())
	configuration := conf.Get()
	db := data.NewSqliteDB("data.db")
	api := api.NewAPI(db)
	router := routes.NewCorsRoutes(api)

	port := ":" + configuration.Port
	log.Printf("Server listening at: 127.0.0.1%s", port)
	http.ListenAndServe(port, router)
}
