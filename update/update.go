package update

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const CURRENT_VERSION = "0.0.1"

// Based on schhabra2's self-update implementation for the file-sharing app RiftShare
// https://github.com/achhabra2/riftshare/blob/main/internal/update/selfupdate.go
func DoSelfUpdate() error {
	v := semver.MustParse(CURRENT_VERSION)
	selfupdate.EnableLog()
	latest, err := selfupdate.UpdateSelf(v, "samfweb/mqtt-viewer")
	if err != nil {
		return fmt.Errorf("self update failed: %s", err)
	}
	if latest.Version.Equals(v) {
		log.Println("Current binary is the latest version", CURRENT_VERSION)
		return nil
	} else {
		log.Println("Successfully updated to version", latest.Version)
		log.Println("Release note:\n", latest.ReleaseNotes)
		return nil
	}
}

func DoSelfUpdateMac() bool {
	latest, found, _ := selfupdate.DetectLatest("samfweb/wails-selfupdate-spike")
	if found {
		homeDir, _ := os.UserHomeDir()
		downloadPath := filepath.Join(homeDir, "Downloads", "SelfUpdateTest.zip")
		err := exec.Command("curl", "-L", latest.AssetURL, "-o", downloadPath).Run()
		if err != nil {
			log.Println("curl error:", err)
			return false
		}
		appPath := "/Applications/"
		if cmdPath, err := os.Executable(); err == nil {
			appPath = strings.TrimSuffix(cmdPath, "SelfUpdateTest.app/Contents/MacOS/SelfUpdateTest")
		}
		err = exec.Command("ditto", "-xk", downloadPath, appPath).Run()
		if err != nil {
			log.Println("ditto error:", err)
			return false
		}
		err = exec.Command("rm", downloadPath).Run()
		if err != nil {
			log.Println("removing error:", err)
			return false
		}
		return true
	} else {
		return false
	}
}

func CheckForUpdate() (bool, string) {
	latest, found, err := selfupdate.DetectLatest("samfweb/wails-selfupdate-spike")
	if err != nil {
		log.Println("Error occurred while detecting version:", err)
		return false, ""
	}

	v := semver.MustParse(CURRENT_VERSION)
	if !found || latest.Version.LTE(v) {
		log.Println("Current version is the latest")
		return false, ""
	}

	return true, latest.Version.String()
}
