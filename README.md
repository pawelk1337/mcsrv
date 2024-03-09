# McSrv: A simple library for creating and managing minecraft servers
>
> Hey! We have a [WIKI](https://github.com/pawelk1337/mcsrv/wiki)!

![https://www.codefactor.io/repository/github/pawelk1337/mcsrv/badge/main](https://www.codefactor.io/repository/github/pawelk1337/mcsrv)

Table of Contents
---

- [McSrv: A simple library for creating and managing minecraft servers](#mcsrv-a-simple-library-for-creating-and-managing-minecraft-servers)
	- [Table of Contents](#table-of-contents)
	- [Features](#features)
	- [Getting started](#getting-started)
	- [More Examples](#more-examples)
		- [Check the examples directory](#check-the-examples-directory)
	- [Recommended projects](#recommended-projects)
	- [TODO](#todo)
	- [Contact](#contact)

Features
---

- Server Downloading
- Server **Events** (e.g on Player join, on Player Death, on Chat...)
- Sending commands to server
- Easy server creation
- Multiple Server Engine Support

Getting started
---

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
    }, log) // log can be nil
    if err != nil {
        panic(err) // Replace with actual error handling
    }

    // Start the server
    go srv.Start()
    println("starting server")
    defer srv.Stop()

    // Wait for the server to start
    <-srv.Wrapper.Loaded()

    // Wait for the server to stop
    <-srv.Wrapper.Stopped()
}
```

More Examples
---

### Check the [examples](https://github.com/pawelk1337/mcsrv/tree/main/examples) directory

Recommended projects
---

- [Gate](https://github.com/minekube/gate) - High-Performance, Low-Memory, Lightweight, Extensible Minecraft Reverse Proxy with Excellent Multi-Protocol Version Support - Velocity/Bungee Replacement - Ready for dev and large deploy!
- [Minecraft Router](https://github.com/AbandonTech/minecraftrouter) - Route Minecraft traffic from a configuration file or api.
- [go-Liter](https://github.com/LiterMC/go-liter) - A proxy for Minecraft that can be extended with JavaScript

TODO
---

- ~~Server Downloading~~
- ~~Server running and Server Wrapping~~
- ~~Make Servers Importable~~
- ~~Custom Log Handling~~
- ~~Fix wrapper.Loaded()~~
- ~~add wrapper.Stopped()~~
- Parsing server files (e.g server.properties...)

Contact
---

You can contact me on discord: `pawelk1337` (id: 967830338116153496)

Feel free to contact me about **anything** i will try my hardest to anwser any of your questions.
