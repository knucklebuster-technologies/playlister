package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"github.com/qawarrior/playlister/controllers"
	"github.com/qawarrior/playlister/models"
)

var router = mux.NewRouter().StrictSlash(true)

func main() {
	log.Println("Starting Main Loop")

	config := loadConfig()

	db := connectDatabase(config.Data)
	defer db.Close()

	userRoutes(config.Data, db)

	artistRoutes(config.Data, db)

	startServer(config.Server)
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

func startServer(c models.ServerConfig) {
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

func userRoutes(c models.DataConfig, db *mgo.Session) {
	log.Println("Setting up user routes and handlers")
	uc := controllers.NewUserController(db.Copy().DB(c.DBName))
	router.HandleFunc("/v1/user", uc.GetUser).Methods("GET")
	router.HandleFunc("/v1/user", uc.DeleteUser).Methods("DELETE")
	router.HandleFunc("/v1/user", uc.PostUser).Methods("POST")
}

func artistRoutes(c models.DataConfig, db *mgo.Session) {
	log.Println("Setting up artidt routes and handlers")
	ac := controllers.NewArtistController(db.Copy().DB(c.DBName))
	router.HandleFunc("/v1/artist", ac.GetArtist).Methods("GET")
	router.HandleFunc("/v1/artist", ac.DeleteArtist).Methods("DELETE")
	router.HandleFunc("/v1/artist", ac.PostArtist).Methods("POST")
}
