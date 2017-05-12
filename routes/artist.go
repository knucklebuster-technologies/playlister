package routes

import (
	"github.com/gorilla/mux"
	"github.com/qawarrior/playlister/controllers"
	mgo "gopkg.in/mgo.v2"
)

func artist(dbname string, db *mgo.Session, router *mux.Router) {
	c := controllers.NewArtistController(db.Copy().DB(dbname))
	router.HandleFunc("/v1/Artist", c.Create).Methods("PODT")
	router.HandleFunc("/v1/artist", c.Read).Methods("GET")
	router.HandleFunc("/v1/artist", c.Update).Methods("PUT")
	router.HandleFunc("/v1/artist", c.Delete).Methods("DELETE")
}
