package main

import (
	"log"
	"phonebook/server"
	"phonebook/server/setup"
)

func main() {
	openErr, pingErr := setup.ConnectToDB()
	if openErr != nil {
		log.Fatalf("error connect to DB %v", openErr)
	}

	if pingErr != nil {
		log.Fatalf("error ping to DB %v", pingErr)
	}

	server.InitRoutes()
}
