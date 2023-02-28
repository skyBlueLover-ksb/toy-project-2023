package menu

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	cmd := exec.Command("clear") // for Unix-like systems
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // for Windows
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}