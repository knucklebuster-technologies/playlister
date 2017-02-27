package models

// StatusMessage used to send responses to json rest calls
type StatusMessage struct {
	Status  string      `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Info    interface{} `json:"info" bson:"info"`
}
