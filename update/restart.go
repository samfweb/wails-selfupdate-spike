package update

import (
	"log"
	"os"
	"os/exec"
	"runtime"

	"syscall"
)

// Courtesy of https://github.com/wailsapp/wails/discussions/2223

func restartSelf() {
	self, _ := os.Executable()
	args := os.Args
	env := os.Environ()

	// Windows does not support exec syscall.
	if runtime.GOOS == "windows" {
		cmd := exec.Command(self, args[1:]...)
		cmd.Env = env
		if err := cmd.Start(); err == nil {
			os.Exit(0)
		}
	} else {
		err := syscall.Exec(self, args, env)
		if err != nil {
			log.Panic(err)
			return
		}
	}
}
