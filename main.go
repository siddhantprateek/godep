package main

import (
	"log"

	app "github.com/siddhantprateek/godep/routes"
)

func main() {

	err := app.ApiServer()
	if err != nil {
		log.Fatal("Unable to start API server.")
	}
}
