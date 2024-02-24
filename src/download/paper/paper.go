package paper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pawelk1337/mcsrv/shared"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Download(
	Path string,
	Version string,
	Build string,
) error {
	// Get versions form the paper server
	versions, err := GetVersions()
	if err != nil {
		return err
	}

	// Check if user picked the latest version
	if strings.ToLower(Version) == "latest" {
		Version = versions[len(versions)-1]
	}

	// If version was not found return error
	if !contains(versions, Version) {
		return fmt.Errorf("version %s not found", Version)
	}

	// Get version info
	verInfo, err := GetVersionInfo(Version)
	if err != nil {
		return err
	}

	// Get build info
	var buildInfo papermcBuildInfoJson
	// Check if user picked the latest build
	if strings.ToLower(Build) != "latest" {
		buildInfo, err = GetBuildInfo(Version, strings.ToLower(Build))
		if err != nil {
			return err
		}
	} else {
		build := strconv.Itoa(verInfo.Builds[len(verInfo.Builds)-1])
		buildInfo, err = GetBuildInfo(Version, build)
		if err != nil {
			return err
		}
	}

	buildStr := strconv.Itoa(buildInfo.Build)
	// Download the file
	err = shared.Download(
		"https://api.papermc.io/v2/projects/paper/versions/"+Version+"/builds/"+buildStr+"/downloads/"+"paper-"+Version+"-"+buildStr+".jar",
		Path,
	)
	if err != nil {
		println("https://api.papermc.io/v2/projects/paper/versions/" + Version + "/builds/" + buildStr + "/downloads/" + "paper-" + Version + "-" + buildStr + ".jar")
		return err
	}

	// Checksum the file
	checksum, err := shared.Sha256File(Path)
	if err != nil {
		return err
	}

	// Check if checksum is correct
	if strings.ToLower(checksum) != buildInfo.Downloads.Application.Sha256 {
		return fmt.Errorf("Checksum mismatch (%s != %s)", checksum, buildInfo.Downloads.Application.Sha256)
	}

	return nil
}
