package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qawarrior/playlister/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ArtistController represents the controller for operating on the Artist resource
type ArtistController struct {
	collection *mgo.Collection
}

// NewArtistController returns a controller for a User
func NewArtistController(s *mgo.Session) *ArtistController {
	c := s.DB(dbName).C("artists")
	return &ArtistController{c}
}

// GetArtist returns a specific Artist in the collection
func (c ArtistController) GetArtist(w http.ResponseWriter, r *http.Request) {
	m := decodeArtist(r.Body, models.Artist{})

	if m.Name == "" {
		sendResponse("ERROR", "name property required to get artist", m, 404, w)
		return
	}

	err := c.collection.Find(nil).One(&m)
	if err != nil {
		sendResponse("ERROR", "failed to get artist", err, 404, w)
		return
	}

	mj, _ := json.Marshal(m)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", mj)
}

// PostArtist creates a new Artist document in the database
func (c ArtistController) PostArtist(w http.ResponseWriter, r *http.Request) {
	m := decodeArtist(r.Body, models.Artist{})

	m.ID = bson.NewObjectId()

	err := c.collection.Insert(&m)
	if err != nil {
		sendResponse("ERROR", "failed to create artist", m, 404, w)
		return
	}

	sendResponse("SUCCESS", "artist was created", m, 201, w)
}

// DeleteArtist removes an Artist document in the database
func (c ArtistController) DeleteArtist(w http.ResponseWriter, r *http.Request) {
	m := decodeArtist(r.Body, models.Artist{})

	// check that we have a name to use for deletion
	if m.Name == "" {
		sendResponse("ERROR", "name property needed to delete artist", m, 404, w)
		return
	}

	err := c.collection.Remove(bson.M{"name": m.Name})

	if err != nil {
		sendResponse("ERROR", "delete failed", err, 404, w)
		return
	}

	sendResponse("SUCCESS", "artist was deleted", m, 200, w)
}
