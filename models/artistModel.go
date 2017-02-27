package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Artist represents a music artists info
type Artist struct {
	ID        bson.ObjectId `json:"id"  bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Founded   time.Time     `json:"founded" bson:"founded"`
	Disbanded time.Time     `json:"disbanded" bson:"disbanded"`
	Members   []string      `json:"members" bson:"members"`
	Releases  []string      `json:"releases" bson:"releases"`
}
