package shared

import (
	"github.com/pawelk1337/mcsrv/wrapper"
)

type ServerEngine uint8

// Vanilla servers
const (
	VANILLA ServerEngine = 0
	PAPER   ServerEngine = 1
	PURPUR  ServerEngine = 2
)

type ServerConfig struct {
	AcceptEula bool `json:"acceptEula"` // Required to run the server (its the mojangs eula)

	Path string `json:"path"` // Server path (directory)

	Engine  ServerEngine `json:"engine"`  // Server engine use constants defined in the types.go file (VANILLA"  SPIGOT"  PAPER...)
	Version string       `json:"version"` // Server version (1.20.2"  1.16.5) (use "latest" or "" for the latest version)
	Build   string       `json:"build"`   // Server Build (use "latest" or "" for the latest build) apples only to nonVanilla engines

	Port string `json:"port"` // Server port minecraft default is 25565
	Host string `json:"host"` // Server host

	InitialHeapSize int `json:"initHeap"` // Server heap sizes (In MB)
	MaxHeapSize     int `json:"maxHeap"`  // Server heap sizes (In MB)
}

// HandMade"  Based on https://server.properties/
type ServerProperties struct {
	Host                           bool   `properties:"enable-jmx-monitoring" default:"false"`
	RconPort                       int    `properties:"rcon.port" default:"25575"`
	Seed                           int    `properties:"level-seed" default:""`
	Gamemode                       string `properties:"gamemode" default:"survival"`
	EnableCommandBlock             bool   `properties:"enable-command-block" default:"false"`
	EnableQuery                    bool   `properties:"enable-query" default:"false"`
	GeneratorSettings              string `properties:"generator-settings" default:"{}"`
	EnforceSecureProfile           bool   `properties:"enforce-secure-profile" default:"true"`
	LevelName                      string `properties:"level-name" default:"25575"`
	Motd                           string `properties:"motd" default:"A mcsrv Server\nPawelk1337 @ github.com"`
	QueryPort                      int    `properties:"query.port" default:"25565"`
	PvP                            bool   `properties:"pvp" default:"true"`
	GenerateStructures             bool   `properties:"generate-structures" default:"true"`
	MaxChainedNeighborUpdates      int    `properties:"max-chained-neighbor-updates" default:"1000000"`
	Difficulty                     string `properties:"difficulty" default:"easy"`
	NetworkCompressionThreshold    int    `properties:"network-compression-threshold" default:"256"`
	MaxTickTime                    int    `properties:"max-tick-time" default:"60000"`
	RequireResourcePack            bool   `properties:"require-resource-pack" default:"false"`
	UseNativeTransport             bool   `properties:"use-native-transport" default:"true"`
	MaxPlayers                     int    `properties:"max-players" default:"20"`
	OnlineMode                     bool   `properties:"online-mode" default:"true"`
	EnableStatus                   bool   `properties:"enable-status" default:"true"`
	AllowFlight                    bool   `properties:"allow-flight" default:"false"`
	InitialDisabledPacks           string `properties:"initial-disabled-packs" default:""`
	BroadCastRconToOps             bool   `properties:"broadcast-rcon-to-ops" default:"true"`
	ViewDistance                   int    `properties:"view-distance" default:"10"`
	ServerIP                       string `properties:"server-ip" default:""`
	ResourcePackPrompt             string `properties:"resource-pack-prompt" default:""`
	AllowNether                    bool   `properties:"allow-nether" default:"true"`
	ServerPort                     string `properties:"server-port" default:"25565"`
	EnableRcon                     bool   `properties:"enable-rcon" default:"false"`
	SyncChunkWrites                bool   `properties:"sync-chunk-writes" default:"true"`
	OpPermissionLevel              int    `properties:"op-permission-level" default:"4"`
	PreventProxyConnections        bool   `properties:"prevent-proxy-connections" default:"false"`
	HideOnlinePlayers              bool   `properties:"hide-online-players" default:"false"`
	ResourcePack                   string `properties:"resource-pack" default:""`
	EntityBroadCastRangePercentage int    `properties:"entity-broadcast-range-percentage" default:"100"`
	SimulationDistanse             int    `properties:"simulation-distance" default:"10"`
	RconPassword                   string `properties:"rcon.password" default:""`
	PlayerIdleTimeout              int    `properties:"player-idle-timeout" default:"0"`
	Debug                          string `properties:"debug" default:"false"`
	ForceGamemode                  bool   `properties:"force-gamemode" default:"false"`
	RateLimit                      int    `properties:"rate-limit" default:"0"`
	Hardcore                       bool   `properties:"hardcore" default:"false"`
	WhiteList                      bool   `properties:"white-list" default:"VAL"`
	BrodCastConsoleToOps           bool   `properties:"broadcast-console-to-ops" default:"true"`
	SpawnNpcs                      bool   `properties:"spawn-npcs" default:"true"`
	SpawnAnimals                   bool   `properties:"spawn-animals" default:"true"`
	LogIps                         bool   `properties:"log-ips" default:"true"`
	FunctionPermissionLevel        int    `properties:"function-permission-level" default:"2"`
	InitialEnabledPacks            string `properties:"initial-enabled-packs" default:"vanilla"`
	LevelType                      string `properties:"level-type" default:"minecraft\:normal"`
	TextFilteringConfig            string `properties:"text-filtering-config" default:""`
	SpawnMonsters                  bool   `properties:"spawn-monsters" default:"true"`
	EnforceWhitelist               bool   `properties:"enforce-whitelist" default:"false"`
	SpawnProtection                int    `properties:"spawn-protection" default:"16"`
	ResourcePackSha1               string `properties:"resource-pack-sha1" default:""`
	MaxWorldSize                   int    `properties:"max-world-size" default:"29999984"`
}

type Server struct {
	Config  ServerConfig
	Console *wrapper.DefaultConsole
	Wrapper *wrapper.Wrapper
}

func (s *Server) Start() error {
	return s.Wrapper.Start()
}

func (s *Server) Stop() error {
	return s.Wrapper.Stop()
}

func (s *Server) Wait() error {
	<-s.Wrapper.Loaded()
	return nil
}
