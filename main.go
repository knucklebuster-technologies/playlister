package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/qawarrior/playlister/models"
)

func main() {
	log.Println("STARTING MAIN")
	wdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("WORKING DIRECTORY:", wdir)

	config := readConfig(wdir + `\config.json`)

	dbsrv := newDBServer(wdir + `\db`)
	err = dbsrv.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer dbsrv.Stop()

	db := connectDatabase(config.Data)
	defer db.Close()

	log.Println("SETTING UP HTTP ROUTER")
	router := mux.NewRouter().StrictSlash(true)

	userRoutes(config.Data, db, router)

	artistRoutes(config.Data, db, router)

	startHTTPServer(config.Server, router)
}

func readConfig(path string) models.AppConfig {
	log.Println("READING CONFIGURATION FROM:", path)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	configuration := models.AppConfig{}
	json.Unmarshal(file, &configuration)
	log.Println("CONFIGURATION READ SUCCESSFULLY")
	return configuration
}

func startHTTPServer(c models.ServerConfig, router *mux.Router) {
	log.Println("SETTING UP HTTP SERVER")
	srv := &http.Server{
		Handler:      router,
		Addr:         c.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("STARTING HTTP SERVER AT:", c.Address)
	log.Fatal(srv.ListenAndServe())
}
