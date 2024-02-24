package shared

type ServerEngine uint8

// Vanilla servers
const (
	VANILLA ServerEngine = 0
	PAPER   ServerEngine = 1
	PURPUR  ServerEngine = 2
)

type ServerConfig struct {
	Path string // Server path (directory)

	Engine  ServerEngine // Server engine use constants defined in the types.go file (VANILLA, SPIGOT, PAPER...)
	Version string       // Server version (1.20.2, 1.16.5) (use "latest" or "" for the latest version)
	Build   string       // Server Build (use "latest" or "" for the latest build) apples only to nonVanilla engines

	Port string // Server port minecraft default is 25565
	Host string // Server host
}

type Server struct {
	Config ServerConfig
}
