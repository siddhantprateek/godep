package main

import (
	app "godep/routes"
	"log"
)

func main() {

	err := app.ApiServer()
	if err != nil {
		log.Fatal("Unable to start API server.")
	}
}
