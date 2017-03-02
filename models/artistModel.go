package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Artist represents a music artists info
type Artist struct {
	ID         bson.ObjectId `json:"id"  bson:"_id"`
	First      string        `json:"first" bson:"first"`
	Last       string        `json:"last" bson:"last"`
	Born       time.Time     `json:"born" bson:"born"`
	Death      time.Time     `json:"death" bson:"death"`
	Birthplace string        `json:"birthplace" bson:"birthplace"`
}
