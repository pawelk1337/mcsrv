package purpur

import (
	"fmt"
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
	// Get the versions from the Purpur servers
	versions, err := GetVersions()
	if err != nil {
		return err
	}

	// Check if the requested version is available
	if strings.ToLower(Version) == "latest" {
		Version = versions[len(versions)-1]
	}

	// Check if the version is found
	if !contains(versions, Version) {
		return fmt.Errorf("version %s not found", Version)
	}

	// Get version info
	verInfo, err := GetVersionInfo(Version)
	if err != nil {
		return err
	}

	// Get build info
	var buildInfo purpurBuildInfoJson
	// Check if user picked the latest build
	if strings.ToLower(Build) != "latest" {
		buildInfo, err = GetBuildInfo(Version, strings.ToLower(Build))
		if err != nil {
			return err
		}
	} else {
		buildInfo, err = GetBuildInfo(Version, verInfo.Builds.Latest)
		if err != nil {
			return err
		}
	}

	err = shared.Download(
		"https://api.purpurmc.org/v2/purpur/"+Version+"/"+buildInfo.Build+"/download",
		Path,
	)
	if err != nil {
		return err
	}

	checksum, err := shared.Md5File(Path)
	if err != nil {
		return err
	}

	if strings.ToLower(checksum) != buildInfo.Md5 {
		return fmt.Errorf("Checksum mismatch (%s != %s)", checksum, buildInfo.Md5)
	}

	return nil
}
