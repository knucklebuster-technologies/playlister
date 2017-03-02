package models

import "time"

// StatusMessage used to send responses to json rest calls
type StatusMessage struct {
	DateTime time.Time   `json:"dateTime" bson:"dateTime"`
	Status   string      `json:"status" bson:"status"`
	Message  string      `json:"message" bson:"message"`
	Info     interface{} `json:"info" bson:"info"`
}
