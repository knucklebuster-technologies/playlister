package database

import "os/exec"
import "gopkg.in/mgo.v2"

// Server type to start, stop, and connect mongo db server
type Server struct {
	path    string
	cmd     *exec.Cmd
	Session *mgo.Session
}

// Start invokes the mongod server daemon to make it available
func (srv *Server) Start() error {
	ensureDBPath(srv.path)
	srv.cmd = exec.Command("mongod", "--dbpath", srv.path)
	err := srv.cmd.Start()
	if err != nil {
		return err
	}
	return nil
}

// Stop send sigterm to the mongod server daemon to end the process
func (srv *Server) Stop() error {
	err := srv.cmd.Process.Kill()
	if err != nil {
		return err
	}
	return nil
}

// Connect sets up a session to the database
func (srv *Server) Connect(uri string) error {
	s, err := mgo.Dial(uri)
	if err != nil {
		return err
	}
	srv.Session = s
	return nil
}

// NewServer returns a value of the Server
func NewServer(dirpath string) *Server {
	return &Server{
		dirpath,
		nil,
		nil,
	}
}
