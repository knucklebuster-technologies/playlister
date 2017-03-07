package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"github.com/qawarrior/playlister/controllers"
	"github.com/qawarrior/playlister/models"
)

var (
	configuration = models.AppConfig{}
	router        = mux.NewRouter().StrictSlash(true)
	db            *mgo.Session
	once          sync.Once
)

func main() {
	log.Println("Starting Main Loop")
	defer db.Close()

	loadConfig()

	connectDatabase()

	userRoutes()

	artistRoutes()

	startServer()
}

func startServer() {
	log.Println("Defining HTTP Server")
	srv := &http.Server{
		Handler:      router,
		Addr:         configuration.Server.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Starting HTTP Server @", configuration.Server.Address)
	log.Fatal(srv.ListenAndServe())
}

func loadConfig() {
	log.Println("Reading config.json")
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(file, &configuration)
	log.Println("Application config loaded")
}

func connectDatabase() {
	once.Do(func() {
		log.Println("Connecting to database @", configuration.Data.URI)
		s, err := mgo.Dial(configuration.Data.URI)

		if err != nil {
			log.Fatal(err)
		}

		db = s
		log.Println("Database connected")
	})
}

func userRoutes() {
	log.Println("Setting up user routes and handlers")
	uc := controllers.NewUserController(db.Copy().DB(configuration.Data.DBName))
	router.HandleFunc("/v1/user", uc.GetUser).Methods("GET")
	router.HandleFunc("/v1/user", uc.DeleteUser).Methods("DELETE")
	router.HandleFunc("/v1/user", uc.PostUser).Methods("POST")
}

func artistRoutes() {
	log.Println("Setting up artidt routes and handlers")
	ac := controllers.NewArtistController(db.Copy().DB(configuration.Data.DBName))
	router.HandleFunc("/v1/artist", ac.GetArtist).Methods("GET")
	router.HandleFunc("/v1/artist", ac.DeleteArtist).Methods("DELETE")
	router.HandleFunc("/v1/artist", ac.PostArtist).Methods("POST")
}
