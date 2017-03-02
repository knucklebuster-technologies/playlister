package main

import (
	"log"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"github.com/qawarrior/playlister/controllers"
)

var (
	router = mux.NewRouter().StrictSlash(true)
	db     *mgo.Session
	uri    = "mongodb://localhost"
)

func init() {
	// Connect to our local mongo
	s, err := mgo.Dial(uri)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	db = s
}

func main() {
	//handle /user path
	userRoutes()
	artistRoutes()

	//starts and runs the http server
	startServer()
}

func userRoutes() {
	// Get controll instance
	uc := controllers.NewUserController(db.Copy())

	// Setup routes using the controller functions
	router.HandleFunc("/user/v1/all", uc.GetUsers).Methods("GET")
	router.HandleFunc("/user/v1/{email}/{password}", uc.GetUser).Methods("GET")
	router.HandleFunc("/user/v1/{email}/{password}", uc.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/v1/new", uc.PostUser).Methods("POST")
}

func artistRoutes() {
	ac := controllers.NewArtistController(db.Copy())
	router.HandleFunc("/v1/artist", ac.GetArtist).Methods("GET")
	router.HandleFunc("/v1/artist", ac.DeleteArtist).Methods("DELETE")
	router.HandleFunc("/v1/artist", ac.PostArtist).Methods("POST")
}

func startServer() {
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
