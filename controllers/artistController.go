package controllers

import (
	"encoding/json"
	"fmt"
	"log"
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
func NewArtistController(d *mgo.Database) *ArtistController {
	log.Println("Returning an ArtistController")
	c := d.C("artists")
	return &ArtistController{c}
}

// GetArtist returns a specific Artist in the collection
func (c ArtistController) GetArtist(w http.ResponseWriter, r *http.Request) {
	log.Println("GetArtist called")
	m := decodeArtist(r.Body, models.Artist{})

	if m.First == "" || m.Last == "" {
		log.Println("Can not perform read without first and last json values")
		sendResponse("ERROR", "Can not perform read without first and last json values", m, 404, w)
		return
	}

	err := c.collection.Find(nil).One(&m)
	if err != nil {
		log.Println("Failed to find artist in database")
		sendResponse("ERROR", "failed to get artist", err, 404, w)
		return
	}

	mj, _ := json.Marshal(m)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", mj)
	log.Println("Artist was returned")
}

// PostArtist creates a new Artist document in the database
func (c ArtistController) PostArtist(w http.ResponseWriter, r *http.Request) {
	log.Println("PostArtist called")
	m := decodeArtist(r.Body, models.Artist{})

	if m.First == "" || m.Last == "" {
		log.Println("Can not perform insert without first and last json values")
		sendResponse("ERROR", "Can not perform insert without first and last json values", m, 404, w)
		return
	}

	m.ID = bson.NewObjectId()

	err := c.collection.Insert(&m)
	if err != nil {
		log.Println("Failed to insert artist into database")
		sendResponse("ERROR", "failed to create artist", m, 404, w)
		return
	}

	sendResponse("SUCCESS", "artist was created", m, 201, w)
	log.Println("Artist was created")
}

// DeleteArtist removes an Artist document in the database
func (c ArtistController) DeleteArtist(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteArtist called")
	m := decodeArtist(r.Body, models.Artist{})

	// check that we have a name to use for deletion
	if m.First == "" || m.Last == "" {
		log.Println("Can not perform remove without first and last json values")
		sendResponse("ERROR", "Can not perform remove without first and last json values", m, 404, w)
		return
	}

	err := c.collection.Remove(bson.M{"first": m.First, "last": m.Last})

	if err != nil {
		log.Println("Failed to remove from database -", err)
		sendResponse("ERROR", "Failed to remove from database", err, 404, w)
		return
	}

	log.Println("artist was deleted")
	sendResponse("SUCCESS", "artist was deleted", m, 200, w)
}
