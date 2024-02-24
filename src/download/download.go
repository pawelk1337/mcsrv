package download

import (
	"github.com/pawelk1337/mcsrv/download/paper"
	"github.com/pawelk1337/mcsrv/download/purpur"
	"github.com/pawelk1337/mcsrv/download/vanilla"
	"github.com/pawelk1337/mcsrv/shared"
)

// Download a minecraft server from the internet
// Path - Where to download (use full path e.g ./servers/test/server.jar)
// Engine - What server engine to use (use VANILLA, SPIGOT, etc. from the mcsrv package)
// Version - What version to download (1.20.2, 1.16.5, 1.8.8, etc.)
// Build - What build to use (use "latest" or "" for the latest build) apples only to nonVanilla engines
func Download(
	Path string,

	Engine shared.ServerEngine,
	Version string,
	Build string,
) error {
	// Check if filepath is empty
	if Path == "" {
		return shared.PathEmpty
	}
	// Check if path already exists
	if exists, err := shared.Exists(Path); err == nil && exists {
		return shared.PathAlreadyExists
	}

	// Download the server
	switch Engine {
	case shared.VANILLA:
		return vanilla.Download(Path, Version)
	case shared.PAPER:
		return paper.Download(Path, Version, Build)
	case shared.PURPUR:
		return purpur.Download(Path, Version, Build)
	default:
		return shared.EngineNotFound
	}
}
