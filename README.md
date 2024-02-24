# McSrv: A simple library for creating and managing minecraft servers.

## Getting started
To download the library use the command
`go get https://github.com/pawelk1337/mcsrv`

### Simple Starting Server Example
```
import (
    "time"

    mc "github.com/pawelk1337/mcsrv"
)

func main() {
    srvcfg := mc.ServerConfig{
        Engine: mc.PAPER,
        Version: "1.20.2
    }

    // Create a server
    srv, err := mc.NewServer(srvcfg)
    if err != nil {
        panic(err)
    }

    // Start the server
    wrapper, err := srv.Start()
    if err != nil {
        panic(err)
    }

    // Wait for the server to start
    wrapper.Wait()

    // Run a command
    wrapper.SendCommand("say Hello World!")

    time.sleep(time.Second * 30)

    // Stop the server
    wrapper.Stop()
}
```