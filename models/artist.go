package models

import (
	"encoding/json"
	"io"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Artist represents a music artists info
type Artist struct {
	ID         bson.ObjectId `json:"id"  bson:"_id"`
	First      string        `json:"first" bson:"first"`
	Last       string        `json:"last" bson:"last"`
	Born       time.Time     `json:"born" bson:"born"`
	Died       time.Time     `json:"died" bson:"died"`
	Birthplace string        `json:"birthplace" bson:"birthplace"`
	Bio        string        `json:"bio" bson:"bio"`
}

// Encode writes the structs value to a stream
func (a *Artist) Encode(w io.Writer) error {
	return json.NewEncoder(w).Encode(a)
}

// Decode reads a stream and assigns values to the structs properties
func (a *Artist) Decode(r io.Reader) error {
	return json.NewDecoder(r).Decode(a)
}
