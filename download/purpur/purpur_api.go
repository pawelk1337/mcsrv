package purpur

import "github.com/pawelk1337/mcsrv/shared"

// https://api.purpurmc.org/v2/purpur
type purpurJson struct {
	Project  string   `json:"project"`
	Versions []string `json:"versions"`
}

type purpurVerInfoJson struct {
	Project string `json:"project"`
	Version string `json:"version"`
	Builds  struct {
		Latest string   `json:"latest"`
		All    []string `json:"all"`
	} `json:"builds"`
}

type purpurBuildInfoJson struct {
	Project   string `json:"project"`
	Version   string `json:"version"`
	Build     string `json:"build"`
	Result    string `json:"result"`
	Timestamp int64  `json:"timestamp"`
	Duration  int    `json:"duration"`
	Commits   []struct {
		Author      string `json:"author"`
		Email       string `json:"email"`
		Description string `json:"description"`
		Hash        string `json:"hash"`
		Timestamp   int64  `json:"timestamp"`
	} `json:"commits"`
	Md5 string `json:"md5"`
}

func GetVersions() ([]string, error) {
	var parsed purpurJson
	err := shared.FetchJson("https://api.purpurmc.org/v2/purpur", &parsed)

	return parsed.Versions, err
}

func GetVersionInfo(version string) (purpurVerInfoJson, error) {
	var parsed purpurVerInfoJson
	err := shared.FetchJson("https://api.purpurmc.org/v2/purpur/"+version, &parsed)

	return parsed, err
}

func GetBuildInfo(version string, build string) (purpurBuildInfoJson, error) {
	var parsed purpurBuildInfoJson
	err := shared.FetchJson("https://api.purpurmc.org/v2/purpur/"+version+"/"+build, &parsed)

	return parsed, err
}
