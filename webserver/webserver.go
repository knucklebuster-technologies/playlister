package webserver

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Start invoke the webserver on provides address, port and router
func Start(address string, router *mux.Router) {
	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
