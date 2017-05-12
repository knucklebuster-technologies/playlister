package main

import (
	"log"
	"os"

	"github.com/qawarrior/playlister/config"
	"github.com/qawarrior/playlister/database"
	"github.com/qawarrior/playlister/routes"
	"github.com/qawarrior/playlister/webserver"
)

func main() {
	log.Println("STARTING MAIN")

	log.Println("GETTING WORKING DIRECTORY")
	wdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("WORKING DIRECTORY:", wdir)

	log.Println("READING CONFIGURATION")
	config := config.Config{}
	err = config.Read(wdir + `\config.json`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("STARTING DATABASE SERVER")
	dbsrv := database.NewServer(wdir + `\db`)
	err = dbsrv.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer dbsrv.Stop()

	log.Println("CREATING DATABASE SESSION")
	err = dbsrv.Connect(config.Data.URI)
	if err != nil {
		log.Fatal(err)
	}
	defer dbsrv.Session.Close()

	log.Println("SETTING UP ROUTING")
	router, err := routes.Set(config.Data.DbName, dbsrv.Session)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting Webserver at " + config.Server.Address)
	webserver.Start(config.Server.Address, router)
	defer webserver.Stop()
}
