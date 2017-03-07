package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"github.com/qawarrior/playlister/models"
)

var (
	configuration = models.AppConfig{}
	router        = mux.NewRouter().StrictSlash(true)
	db            *mgo.Session
)

func main() {
	log.Println("Calling loadConfig")
	loadConfig()

	log.Println("Calling connectDatabase")
	connectDatabase()

	log.Println("Calling userRoutes")
	userRoutes()

	log.Println("Calling artistRoutes")
	artistRoutes()

	//starts and runs the http server
	log.Println("Calling startServer")
	startServer()
}

func startServer() {
	srv := &http.Server{
		Handler:      router,
		Addr:         configuration.Server.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func loadConfig() {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &configuration)
}

func connectDatabase() {
	var once sync.Once
	once.Do(func() {
		// Connect to our local mongo
		s, err := mgo.Dial(configuration.Data.URI)

		// Check if connection error, is mongo running?
		if err != nil {
			panic(err)
		}

		db = s
	})
}
