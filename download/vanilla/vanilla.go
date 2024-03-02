package vanilla

import (
	"fmt"
	"strings"

	"github.com/pawelk1337/mcsrv/shared"
)

func Download(
	Path string,
	Version string,
) error {
	// Get versions form the mojang servers
	versions, err := GetVersions()
	if err != nil {
		return err
	}

	// Check if user picked the latest version
	if strings.ToLower(Version) == "latest" {
		Version = versions.Latest.Release
	}

	// Get the download url
	var versionDetailsDownloadURL string
	for i := 0; i < len(versions.Versions); i++ {
		if Version == versions.Versions[i].ID {
			versionDetailsDownloadURL = versions.Versions[i].URL
		}
	}

	// If download url was not found return an error
	if versionDetailsDownloadURL == "" {
		return fmt.Errorf("version \"%s\" not found", Version)
	}

	// Get the version details form the mojang serves
	versionDetails, err := GetVersionDetails(versionDetailsDownloadURL)
	if err != nil {
		return err
	}

	// Download the file
	err = shared.Download(
		versionDetails.Downloads.Server.URL,
		Path,
	)
	if err != nil {
		return err
	}

	// Check if file downloaded successfully
	checksum, err := shared.Sha1File(Path)
	if err != nil {
		return err
	}

	if strings.ToLower(checksum) != strings.ToLower(versionDetails.Downloads.Server.Sha1) {
		return fmt.Errorf("checksum mismatch (%s != %s)", checksum, versionDetails.Downloads.Server.Sha1)
	}

	// if downloaded successfully return
	return nil
}
