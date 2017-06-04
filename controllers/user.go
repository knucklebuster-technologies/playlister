package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qawarrior/playlister/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User represents the controller for operating on the User resource
type User struct {
	collection *mgo.Collection
}

// NewUser returns a controller for the User endpoint
func NewUser(d *mgo.Database) *User {
	c := d.C("users")
	return &User{c}
}

// Create adds a new user
func (c User) Create(w http.ResponseWriter, r *http.Request) {
	m := models.User{}
	m.Decode(r.Body)
	m.ID = bson.NewObjectId()

	c.collection.Insert(&m)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	m.Encode(w)
}

// Read returns an existing user by email and password
func (c User) Read(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	email := vals.Get("email")
	password := vals.Get("password")

	m := models.User{}

	if err := c.collection.Find(bson.M{"email": email, "password": password}).One(&m); err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	m.Encode(w)
}

// Update modifies an existing user
func (c User) Update(w http.ResponseWriter, r *http.Request) {
	sendResponse("Success", "User Updated", nil, 200, w)
}

// Delete removes an existing user
func (c User) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	password := vars["password"]

	ci, err := c.collection.RemoveAll(bson.M{"email": email, "password": password})
	if err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	cij, _ := json.Marshal(ci)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", cij)
}
