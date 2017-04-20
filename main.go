package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"github.com/qawarrior/playlister/models"
)

func main() {
	log.Println("Starting Main Loop")

	config := loadConfig()

	db := connectDatabase(config.Data)
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

	userRoutes(config.Data, db, router)

	artistRoutes(config.Data, db, router)

	startServer(config.Server, router)
}

func loadConfig() models.AppConfig {
	log.Println("Reading config.json")
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	configuration := models.AppConfig{}
	json.Unmarshal(file, &configuration)
	log.Println("Application config loaded")
	return configuration
}

func startServer(c models.ServerConfig, router *mux.Router) {
	log.Println("Defining HTTP Server")
	srv := &http.Server{
		Handler:      router,
		Addr:         c.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Starting HTTP Server @", c.Address)
	log.Fatal(srv.ListenAndServe())
}

func connectDatabase(c models.DataConfig) *mgo.Session {
	log.Println("Connecting to database @", c.URI)
	s, err := mgo.Dial(c.URI)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")
	return s
}
