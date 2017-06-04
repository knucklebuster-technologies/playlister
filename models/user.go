package models

import (
	"encoding/json"
	"io"

	"gopkg.in/mgo.v2/bson"
)

// User represents the user of the playlister web application
type User struct {
	ID       bson.ObjectId `json:"id"  bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
	Gender   string        `json:"gender" bson:"gender"`
	Age      int           `json:"age" bson:"age"`
}

// Encode writes the structs value to a stream
func (a *User) Encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(a)
}

// Decode reads a stream and assigns values to the structs properties
func (a *User) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(a)
}
