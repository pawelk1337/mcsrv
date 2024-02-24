package main

import (
	mc "github.com/pawelk1337/mcsrv"
	mcsh "github.com/pawelk1337/mcsrv/shared"
	mcevents "github.com/pawelk1337/mcsrv/wrapper/events"
)

func main() {
	srv, err := mc.NewServer(&mcsh.ServerConfig{
		AcceptEula: true,
		Path:       "./server",

		Engine:  mcsh.PAPER,
		Version: "latest", // Use the latest version
		Build:   "latest",

		Port: "25565",
		Host: "127.0.0.1",

		InitialHeapSize: 2048, // 2 GB
		MaxHeapSize:     2048, // 2 GB
	})

	// srv, err := mc.ImportServer("./server")
	if err != nil {
		panic(err)
	}

	// Start the server
	go srv.Start()
	println("starting server")
	defer srv.Stop()

	// Wait for the server to start
	<-srv.Wrapper.Loaded()

	wrp := srv.Wrapper

	// List all players
	players := wrp.List()
	for _, player := range players {
		println(player.Name)
	}

	wrp.GameEvents()

	// Listen for events
	for {
		select {
		case ev := <-wrp.GameEvents():
			if ev.String() == mcevents.PlayerJoined {
				wrp.Say("Hello " + ev.Data["player_name"])
			}
		}
	}
}
