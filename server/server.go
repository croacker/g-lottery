package main

import (
	"./api"
	"./data"
	"./routes"
	"github.com/urfave/negroni"
)

func main() {
	db := data.NewSqliteDB("data.db")
	api := api.NewAPI(db)
	routes := routes.NewRoutes(api)
	n := negroni.Classic()
	n.UseHandler(routes)
	n.Run(":3000")
}
