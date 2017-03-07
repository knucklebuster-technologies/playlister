package main

import "github.com/qawarrior/playlister/controllers"

func userRoutes() {
	uc := controllers.NewUserController(db.Copy().DB(configuration.Data.DBName))
	router.HandleFunc("/v1/user", uc.GetUser).Methods("GET")
	router.HandleFunc("/v1/user", uc.DeleteUser).Methods("DELETE")
	router.HandleFunc("/v1/user", uc.PostUser).Methods("POST")
}

func artistRoutes() {
	ac := controllers.NewArtistController(db.Copy().DB(configuration.Data.DBName))
	router.HandleFunc("/v1/artist", ac.GetArtist).Methods("GET")
	router.HandleFunc("/v1/artist", ac.DeleteArtist).Methods("DELETE")
	router.HandleFunc("/v1/artist", ac.PostArtist).Methods("POST")
}
