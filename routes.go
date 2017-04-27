package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/qawarrior/playlister/controllers"
	"github.com/qawarrior/playlister/models"
	mgo "gopkg.in/mgo.v2"
)

func userRoutes(c models.DataConfig, db *mgo.Session, router *mux.Router) {
	log.Println("SETTING UP USER ROUTES AND HANDLERS")
	uc := controllers.NewUserController(db.Copy().DB(c.DBName))
	router.HandleFunc("/v1/user", uc.Create).Methods("POST")
	router.HandleFunc("/v1/user", uc.Read).Methods("GET")
	router.HandleFunc("/v1/user", uc.Update).Methods("PUT")
	router.HandleFunc("/v1/user", uc.Delete).Methods("DELETE")
}

func artistRoutes(c models.DataConfig, db *mgo.Session, router *mux.Router) {
	log.Println("SETTING UP ARTIST ROUTES AND HANDLERS")
	ac := controllers.NewArtistController(db.Copy().DB(c.DBName))
	router.HandleFunc("/v1/artist", ac.Create).Methods("POST")
	router.HandleFunc("/v1/artist", ac.Read).Methods("GET")
	router.HandleFunc("/v1/artist", ac.Update).Methods("PUT")
	router.HandleFunc("/v1/artist", ac.Delete).Methods("DELETE")
}
