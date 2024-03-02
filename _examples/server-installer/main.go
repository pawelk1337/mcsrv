package main

import (
	"os"
	"strings"

	mc "github.com/pawelk1337/mcsrv"
	mcsh "github.com/pawelk1337/mcsrv/shared"
	cli "github.com/urfave/cli/v2"
)

func newServer(cCtx *cli.Context) error {
	println("Downloading Server...")
	eng := mcsh.PAPER
	switch strings.ToLower(cCtx.String("engine")) {
	case "vanilla":
		eng = mcsh.VANILLA
	case "paper":
		eng = mcsh.PAPER
	case "purpur":
		eng = mcsh.PURPUR
	default:
		eng = mcsh.VANILLA
	}
	srv, err := mc.NewServer(&mcsh.ServerConfig{
		AcceptEula: cCtx.Bool("acceptEula"),

		Path: cCtx.String("path"),

		Engine:  eng,
		Version: cCtx.String("version"),
		Build:   "latest",

		Port: cCtx.String("port"),
		Host: cCtx.String("host"),

		InitialHeapSize: 2048, // 2 GB
		MaxHeapSize:     2048, // 2 GB
	}, func(line string, tick int) {
		print(line)
	})
	if err != nil {
		return err
	}

	srv.Console.Cmd.GetCmd().Stdin = os.Stdin

	srv.Start()
	println("starting server")
	println("press Ctrl+C to exit")

	<-srv.Wrapper.Loaded()

	// Wait for server to stop
	<-srv.Wrapper.Stopped()
	println("server stopped")

	return nil
}

func importServer(cCtx *cli.Context) error {
	srv, err := mc.ImportServer(cCtx.String("path"), func(line string, tick int) {
		print(line)
	})
	if err != nil {
		return err
	}

	srv.Console.Cmd.GetCmd().Stdin = os.Stdin

	srv.Start()
	println("starting server")
	println("press Ctrl+C to exit")

	<-srv.Wrapper.Loaded()
	println("server started")

	// Wait for server to stop
	<-srv.Wrapper.Stopped()
	println("server stopped")

	return nil
}

func main() {
	app := &cli.App{
		Name:  "Server Installer",
		Usage: "install minecraft servers with ease",
		Commands: []*cli.Command{{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create a new server",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "acceptEula",
					Value: false,
					Usage: "Accept Mojang's Eula",
				},
				&cli.StringFlag{
					Name:  "path",
					Value: "./server",
					Usage: "Path to the server",
				},
				&cli.StringFlag{
					Name:  "engine",
					Value: "VANILLA",
					Usage: "Server engine (VANILLA, SPIGOT, PAPER)",
				},
				&cli.StringFlag{
					Name:  "version",
					Value: "latest",
					Usage: "Server version (e.g 1.20.2, 1.18) (use \"latest\" for the latest version)",
				},
				&cli.StringFlag{
					Name:  "port",
					Value: "25565",
					Usage: "Server port",
				},
				&cli.StringFlag{
					Name:  "host",
					Value: "127.0.0.1",
					Usage: "Server host",
				},
				&cli.StringFlag{
					Name:  "initHeap",
					Value: "2048",
					Usage: "Server heap sizes (In MB)",
				},
				&cli.StringFlag{
					Name:  "maxHeap",
					Value: "2048",
					Usage: "Server heap sizes (In MB)",
				},
			},
			Action: newServer,
		},
			{
				Name:    "import",
				Aliases: []string{"i"},
				Usage:   "import an existing server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "path",
						Value: "./server",
						Usage: "Path to the server",
					},
				},
				Action: importServer,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
