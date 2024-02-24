package main

import (
	"github.com/pawelk1337/mcsrv"
	"github.com/pawelk1337/mcsrv/shared"
)

func main() {
	srv, err := mcsrv.NewServer(&shared.ServerConfig{
		Engine:  shared.PAPER,
		Version: "1.20",
		Build:   "latest",
		Port:    "25565",
		Host:    "localhost",
	})

	if err != nil {
		panic(err)
	}

	println(srv.Config.Version)
}
