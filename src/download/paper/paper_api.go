package paper

import (
	"time"

	"github.com/pawelk1337/mcsrv/shared"
)

// https://api.papermc.io/v2/projects/papermc
type papermcJson struct {
	Project  string   `json:"project"`
	Versions []string `json:"versions"`
}

type papermcVerInfoJson struct {
	ProjectID   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	Version     string `json:"version"`
	Builds      []int  `json:"builds"`
}

type papermcBuildInfoJson struct {
	ProjectID   string    `json:"project_id"`
	ProjectName string    `json:"project_name"`
	Version     string    `json:"version"`
	Build       int       `json:"build"`
	Time        time.Time `json:"time"`
	Channel     string    `json:"channel"`
	Promoted    bool      `json:"promoted"`
	Changes     []struct {
		Commit  string `json:"commit"`
		Summary string `json:"summary"`
		Message string `json:"message"`
	} `json:"changes"`
	Downloads struct {
		Application struct {
			Name   string `json:"name"`
			Sha256 string `json:"sha256"`
		} `json:"application"`
		MojangMappings struct {
			Name   string `json:"name"`
			Sha256 string `json:"sha256"`
		} `json:"mojang-mappings"`
	} `json:"downloads"`
}

func GetVersions() ([]string, error) {
	var parsed papermcJson
	err := shared.FetchJson("https://api.papermc.io/v2/projects/paper", &parsed)

	return parsed.Versions, err
}

func GetVersionInfo(version string) (papermcVerInfoJson, error) {
	var parsed papermcVerInfoJson
	err := shared.FetchJson("https://api.papermc.io/v2/projects/paper/versions/"+version, &parsed)

	return parsed, err
}

func GetBuildInfo(version string, build string) (papermcBuildInfoJson, error) {
	var parsed papermcBuildInfoJson
	err := shared.FetchJson("https://api.papermc.io/v2/projects/paper/versions/"+version+"/builds/"+build, &parsed)

	return parsed, err
}
