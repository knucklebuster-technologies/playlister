package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Release a unit or body of work released by an Artist or a Project
type Release struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Released    time.Time     `json:"released" bson:"released"`
	Tracks      []Track       `json:"track" bson:"track"`
	Label       string        `json:"label" bson:"label"`
	Country     string        `json:"country" bson:"country"`
	Description string        `json:"description" bson:"description"`
}
