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

// MarshalJSON returns the struct marshalled to json
func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(u)
}

// Decode takes a stream and uses its data to assign values to the properties
func (u *User) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(u)
}
