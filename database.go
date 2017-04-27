package main

import "os/exec"

type dbserver struct {
	path string
	cmd  *exec.Cmd
}

func (srv *dbserver) Start() error {
	srv.cmd = exec.Command("mongod", "--dbpath", srv.path)
	err := srv.cmd.Start()
	if err != nil {
		return err
	}
	return nil
}

func (srv *dbserver) Stop() {
	srv.cmd.Process.Kill()
}

func newDBServer(dirpath string) *dbserver {
	return &dbserver{
		dirpath,
		nil,
	}
}
