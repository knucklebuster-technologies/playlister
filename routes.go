package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/qawarrior/playlister/controllers"
	"github.com/qawarrior/playlister/models"
	mgo "gopkg.in/mgo.v2"
)

func userRoutes(c models.DataConfig, db *mgo.Session, router *mux.Router) {
	log.Println("Setting up user routes and handlers")
	uc := controllers.NewUserController(db.Copy().DB(c.DBName))
	router.HandleFunc("/v1/user", uc.GetUser).Methods("GET")
	router.HandleFunc("/v1/user", uc.DeleteUser).Methods("DELETE")
	router.HandleFunc("/v1/user", uc.PostUser).Methods("POST")
}

func artistRoutes(c models.DataConfig, db *mgo.Session, router *mux.Router) {
	log.Println("Setting up artists routes and handlers")
	ac := controllers.NewArtistController(db.Copy().DB(c.DBName))
	router.HandleFunc("/v1/artist", ac.GetArtist).Methods("GET")
	router.HandleFunc("/v1/artist", ac.DeleteArtist).Methods("DELETE")
	router.HandleFunc("/v1/artist", ac.PostArtist).Methods("POST")
}
