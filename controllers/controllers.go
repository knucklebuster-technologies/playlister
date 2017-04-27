package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/qawarrior/playlister/models"
)

// Controller defines a type that controls an http endpoint
type Controller interface {
	Create(http.ResponseWriter, *http.Request)
	Read(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

func decodeArtist(b io.ReadCloser, a models.Artist) models.Artist {
	json.NewDecoder(b).Decode(&a)

	return a
}

func sendResponse(status string, message string, info interface{}, code int, w http.ResponseWriter) {
	s := models.StatusMessage{
		Status:  status,
		Message: message,
		Info:    info,
	}

	mj, _ := json.Marshal(s)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintf(w, "%s", mj)
}
