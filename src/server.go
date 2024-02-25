package mcsrv

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pawelk1337/mcsrv/download"
	"github.com/pawelk1337/mcsrv/shared"
	"github.com/pawelk1337/mcsrv/wrapper"
	"github.com/pawelk1337/mcsrv/wrapper/events"
)

func NewServer(srvcfg *shared.ServerConfig) (srv shared.Server, err error) {
	// Set the default server path if it doesn't exist
	if srvcfg.Path == "" {
		srvcfg.Path = "./server"
	}
	// Check if path already exists
	if exists, err := shared.Exists(srvcfg.Path); err == nil && exists {
		return shared.Server{}, shared.PathAlreadyExists
	}

	err = os.MkdirAll(srvcfg.Path, 0777)
	if err != nil {
		return shared.Server{}, err
	}

	err = SaveServer(srvcfg)
	if err != nil {
		return shared.Server{}, err
	}

	// Download The Server
	err = download.Download(
		filepath.Join(srvcfg.Path, "server.jar"),

		srvcfg.Engine,
		srvcfg.Version,
		srvcfg.Build,
	)

	if err != nil {
		return shared.Server{}, err
	}

	if srvcfg.AcceptEula {
		eulaFile, err := os.Create(filepath.Join(srvcfg.Path, "eula.txt"))
		if err != nil {
			return shared.Server{}, err
		}
		defer eulaFile.Close()
		_, err = eulaFile.WriteString("eula=true")
		if err != nil {
			return shared.Server{}, err
		}
	}

	console := wrapper.NewConsole(wrapper.JavaExecCmd(
		filepath.Join(srvcfg.Path),
		srvcfg.InitialHeapSize,
		srvcfg.MaxHeapSize,
	))

	logHandler := func(line string, tick int) (events.Event, events.EventType) {
		return wrapper.LogParserFunc(line, tick)
	}
	wrapper := wrapper.NewWrapper(
		console,
		logHandler,
	)

	srv = shared.Server{
		Config:  *srvcfg,
		Console: console,
		Wrapper: wrapper,
	}

	return srv, nil
}

func ImportServer(path string) (shared.Server, error) {
	content, err := os.ReadFile(filepath.Join(path, "server.json"))
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var srvcfg shared.ServerConfig
	err = json.Unmarshal(content, &srvcfg)
	if err != nil {
		return shared.Server{}, err
	}

	console := wrapper.NewConsole(wrapper.JavaExecCmd(
		filepath.Join(srvcfg.Path),
		srvcfg.InitialHeapSize,
		srvcfg.MaxHeapSize,
	))

	logHandler := func(line string, tick int) (events.Event, events.EventType) {
		return wrapper.LogParserFunc(line, tick)
	}
	wrapper := wrapper.NewWrapper(
		console,
		logHandler,
	)

	return shared.Server{
		Config:  srvcfg,
		Console: console,
		Wrapper: wrapper,
	}, nil
}

func SaveServer(srvcfg *shared.ServerConfig) error {
	// Check if the server exists
	exists, err := shared.Exists(srvcfg.Path)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("server does not exist")
	}

	// Save the server
	data, err := json.MarshalIndent(srvcfg, "", "\t")
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(srvcfg.Path, "server.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}
