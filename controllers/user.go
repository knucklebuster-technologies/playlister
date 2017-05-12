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

// NewUser returns a controller for a User
func NewUser(d *mgo.Database) *User {
	c := d.C("users")
	return &User{c}
}

// Create creates a new user
func (c User) Create(w http.ResponseWriter, r *http.Request) {
	// Stub an user to be populated from the body
	m := models.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&m)

	// Add an Id
	m.ID = bson.NewObjectId()

	// Write the user to mongo
	c.collection.Insert(&m)

	// Marshal provided interface into JSON structure
	mj, _ := json.Marshal(m)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", mj)
}

// Read returns an existing user by email and password
func (c User) Read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	password := vars["password"]

	// Stub user
	m := models.User{}

	// Fetch user
	if err := c.collection.Find(bson.M{"email": email, "password": password}).One(&m); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	mj, _ := json.Marshal(m)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", mj)
}

// Update modifies an existing user in the database
func (c User) Update(w http.ResponseWriter, r *http.Request) {
	sendResponse("Success", "User Updated", nil, 200, w)
}

// Delete removes a existing user
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
