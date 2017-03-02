package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Project contains properties of a band or group of artists
type Project struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	Founded   time.Time     `json:"founded" bson:"founded"`
	Disbanded time.Time     `json:"disbanded" bson:"disbanded"`
	Members   []Artist      `json:"members" bson:"members"`
}
