package webserver

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var srv *http.Server

// Start invoke the webserver on provided address, port and router
func Start(address string, router *mux.Router) {
	srv = &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// Stop the running webserver
func Stop() {
	srv.Shutdown(nil)
}
