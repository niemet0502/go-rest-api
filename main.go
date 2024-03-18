package main

import (
	"go-server/database"
	"go-server/routes"
	"log"

	"net/http"
)

const (
	PORT = ":1909"
)

func main() {

	// init routes
	r := routes.InitRoutes()

	// init the database
	database.InitDB()

	err := http.ListenAndServe(PORT, r)

	if err != nil {
		log.Fatal("The server didn't start")
	}
}
