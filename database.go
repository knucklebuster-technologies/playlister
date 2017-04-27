package main

import (
	"log"
	"os/exec"

	"github.com/qawarrior/playlister/models"
	mgo "gopkg.in/mgo.v2"
)

type dbserver struct {
	path string
	cmd  *exec.Cmd
}

func (srv *dbserver) Start() error {
	log.Println("STARTING MONGODB SERVER")
	log.Println("DATA PATH:", srv.path)
	srv.cmd = exec.Command("mongod", "--dbpath", srv.path)
	err := srv.cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("MONGODB SERVER STARTED")
	return nil
}

func (srv *dbserver) Stop() {
	log.Println("STOPPING MONGODB SERVER")
	err := srv.cmd.Process.Kill()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("MONGODB SERVER STOPPED")
}

func newDBServer(dirpath string) *dbserver {
	return &dbserver{
		dirpath,
		nil,
	}
}

func connectDatabase(c models.DataConfig) *mgo.Session {
	log.Println("CONNECTING TO DATABASE AT", c.URI)
	s, err := mgo.Dial(c.URI)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DATABASE CONNECTED")
	return s
}
