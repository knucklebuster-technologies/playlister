package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Track discribes a single song or entry on a Release
type Track struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Length      time.Time     `json:"length" bson:"length"`
	Description string        `json:"description" bson:"description"`
}
