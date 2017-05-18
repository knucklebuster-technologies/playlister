package routes

import (
	"github.com/gorilla/mux"
	"github.com/qawarrior/playlister/controllers"
	mgo "gopkg.in/mgo.v2"
)

func user(dbname string, db *mgo.Session, router *mux.Router) {
	c := controllers.NewUser(db.Copy().DB(dbname))
	router.HandleFunc("/v1/user", c.Create).Methods("POST")
	router.HandleFunc("/v1/user", c.Read).Methods("GET")
	router.HandleFunc("/v1/user", c.Update).Methods("PUT")
	router.HandleFunc("/vi/user", c.Delete).Methods("DELETE")
}
