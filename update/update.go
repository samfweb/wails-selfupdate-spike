package update

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const (
	REPO = "samfweb/wails-selfupdate-spike"
)

var CurrentVersion string

func init() {
	var ok bool
	CurrentVersion, ok = os.LookupEnv("version")
	if !ok {
		CurrentVersion = "0.0.0"
	}
}

func UpdateSelf(ctx context.Context) error {
	withModule := context.Background()
	if runtime.GOOS == "darwin" {
		return doSelfUpdateMac(withModule)
	}
	return doSelfUpdate(withModule)
}

func doSelfUpdate(ctx context.Context) error {
	v := semver.MustParse(CurrentVersion)
	selfupdate.EnableLog()
	latest, err := selfupdate.UpdateSelf(v, REPO)
	if err != nil {
		return fmt.Errorf("self update failed: %s", err)
	}
	if latest.Version.Equals(v) {
		slog.InfoContext(ctx, fmt.Sprintf("current version %s is the latest", v))
		return nil
	}

	slog.InfoContext(ctx, fmt.Sprintf("successfully updated to version %s", latest.Version))
	slog.InfoContext(ctx, fmt.Sprintf("release notes: %s", latest.ReleaseNotes))
	slog.InfoContext(ctx, "restarting...")
	restartSelf()
	return nil
}

func doSelfUpdateMac(ctx context.Context) error {
	latest, found, err := selfupdate.DetectLatest(REPO)
	if err != nil {
		return err
	}
	if !found {
		slog.ErrorContext(ctx, "update server not found")
		return nil
	}
	if latest.Version.Equals(semver.MustParse(CurrentVersion)) {
		slog.InfoContext(ctx, fmt.Sprintf("current version %s is the latest", CurrentVersion))
		return nil
	}
	homeDir, _ := os.UserHomeDir()
	downloadPath := filepath.Join(homeDir, "Downloads", "wails-selfupdate-spike.zip")
	slog.InfoContext(ctx, fmt.Sprintf("downloading %s to %s", latest.Version, downloadPath))
	err = exec.Command("curl", "-L", latest.AssetURL, "-o", downloadPath).Run()
	if err != nil {
		return err
	}
	slog.InfoContext(ctx, "download successfu!")
	cmdPath, err := os.Executable()
	appPath := strings.TrimSuffix(cmdPath, "wails-selfupdate-spike.app/Contents/MacOS/wails-selfupdate-spike")
	if err != nil {
		appPath = "/Applications/"
	}
	slog.InfoContext(ctx, "overwriting binary at path "+appPath)
	err = exec.Command("ditto", "-xk", downloadPath, appPath).Run()
	if err != nil {
		log.Println("ditto error:", err)
		return err
	}
	slog.InfoContext(ctx, "removing downloaded file from "+downloadPath)
	err = exec.Command("rm", downloadPath).Run()
	if err != nil {
		log.Println("removing error:", err)
		return err
	}
	slog.InfoContext(ctx, fmt.Sprintf("successfully updated to version %s", latest.Version))
	slog.InfoContext(ctx, fmt.Sprintf("release notes: %s", latest.ReleaseNotes))
	slog.InfoContext(ctx, "restarting...")

	restartSelf()
	return nil
}
