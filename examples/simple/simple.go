package main

import (
	mc "github.com/pawelk1337/mcsrv"
	mcsh "github.com/pawelk1337/mcsrv/shared"
	mcevents "github.com/pawelk1337/mcsrv/wrapper/events"
)

// Line includes \n
func log(line string, tick int) {
	print(line)
}

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
	}, log)
	// LogFunc can be nil

	if err != nil {
		panic(err)
	}

	// Start the server
	go srv.Start()
	println("starting server")
	defer srv.Stop()

	// Wait for the server to start
	<-srv.Wrapper.Loaded()
	println("server started")

	wrp := srv.Wrapper

	// List all players
	players := wrp.List()
	for _, player := range players {
		println(player.Name)
	}

	// Listen for events
	for {
		ev := <-wrp.GameEvents()

		switch ev.String() {
		case mcevents.PlayerJoined:
			wrp.Say("Hello " + ev.Data["player_name"])
		}
	}
}
