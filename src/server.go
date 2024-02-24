package mcsrv

import (
	"os"
	"path/filepath"

	"github.com/pawelk1337/mcsrv/download"
	"github.com/pawelk1337/mcsrv/shared"
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

	return shared.Server{*srvcfg}, shared.NotImplemented
}
