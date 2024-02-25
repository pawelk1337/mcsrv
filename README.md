# McSrv: A simple library for creating and managing minecraft servers.
> Hey! We have a [WIKI](https://github.com/pawelk1337/mcsrv/wiki)!

## Features
 - Server Downloading
 - Server **Events** (e.g on Player join, on Player Death, on Chat...)
 - Sending commands to server
 - Easy server creation
 - Multiple Server Engine Support

# TODO
 - ~~Server Downloading~~
 - ~~Server running and Server Wrapping~~
 - ~~Make Servers Importable~~
 - Add Proxy support
 - Custom Log Handling
 - Parsing server files (e.g server.properties...)

## Getting started
To download the library use the command
`go get https://github.com/pawelk1337/mcsrv`

To create a simple server:
```go
import (
    mc "github.com/pawelk1337/mcsrv"
    mcsh "github.com/pawelk1337/mcsrv/shared"
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
	}, nil)
    if err != nil {
		panic(err)
	}

    // Start the server
	go srv.Start()
	println("starting server")
	defer srv.Stop()

	// Wait for the server to start
	<-srv.Wrapper.Loaded()
}
```
## More Examples
### Check the [examples](https://github.com/pawelk1337/mcsrv/tree/main/examples) directory