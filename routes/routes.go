package routes

import (
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// Set creates router and sets up the applications routes
func Set(dbname string, db *mgo.Session) (*mux.Router, error) {
	router = *mux.NewRouter().StrictSlash(true)

	return router, nil
}

func user(dbname string, db *mgo.Session, router *mux.Router) {

}
