package database

import "os/exec"

// Server type to start, stop, and connect mongo db server
type Server struct {
	path string
	cmd  *exec.Cmd
}

// Start invokes the mongod server daemon to make it available
func (srv *Server) Start() error {
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
