package update

import (
	"fmt"
	"log/slog"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

func CheckForUpdate() (bool, string) {
	latest, found, err := selfupdate.DetectLatest(REPO)
	if err != nil {
		slog.Error(fmt.Sprintf("error occurred while finding latest version: %v", err))
		return false, ""
	}

	v := semver.MustParse(version)
	if !found || latest.Version.LTE(v) {
		slog.Info(fmt.Sprintf("current version %s is the latest", v))
		return false, ""
	}
	return true, latest.Version.String()
}
