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

// UserController represents the controller for operating on the User resource
type UserController struct {
	collection *mgo.Collection
}

// NewUserController returns a controller for a User
func NewUserController(d *mgo.Database) *UserController {
	c := d.C("users")
	return &UserController{c}
}

// GetUsers returns all users in the collection
func (c UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Stub a slice of user
	m := []models.User{}

	// Fetch All Users
	if err := c.collection.Find(nil).All(&m); err != nil {
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

// GetUser returns an existing user by email and password
func (c UserController) GetUser(w http.ResponseWriter, r *http.Request) {
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

// PostUser creates a new user
func (c UserController) PostUser(w http.ResponseWriter, r *http.Request) {
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

// DeleteUser removes a existing user
func (c UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
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
