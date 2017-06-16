package main

import (
	"os"

	"github.com/qawarrior/playlister/config"
	"github.com/qawarrior/playlister/database"
	"github.com/qawarrior/playlister/loggy"
	"github.com/qawarrior/playlister/routes"
	"github.com/qawarrior/playlister/webserver"
)

func main() {
	logger := *loggy.New(os.Stdout, os.Stdout, os.Stderr)
	logger.Info.Println("STARTING MAIN")

	logger.Info.Println("GETTING WORKING DIRECTORY")
	wdir, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info.Println("WORKING DIRECTORY:", wdir)

	logger.Info.Println("READING CONFIGURATION")
	config := config.Config{}
	err = config.Read(wdir + `\config.json`)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info.Println("STARTING DATABASE SERVER")
	dbsrv := database.NewServer(wdir + `\db`)
	err = dbsrv.Start()
	if err != nil {
		logger.Fatal(err)
	}
	defer dbsrv.Stop()

	logger.Info.Println("CREATING DATABASE SESSION")
	err = dbsrv.Connect(config.Data.URI)
	if err != nil {
		logger.Fatal(err)
	}
	defer dbsrv.Session.Close()

	logger.Info.Println("SETTING UP ROUTING")
	router, err := routes.Set(config.Data.DbName, dbsrv.Session)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info.Println("Starting Webserver at " + config.Server.Address)
	webserver.Start(config.Server.Address, router)
	defer webserver.Stop()
}
