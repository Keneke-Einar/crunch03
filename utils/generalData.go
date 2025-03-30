package utils

import (
	"fmt"
	"os"
	"os/exec"
)

var charMap map[rune]string = map[rune]string{
	'#': "×",
	'.': "·",
	'o': "∘",
}

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

var (
	h, w       int
	tick       int = 1
	gameMap    [][]rune
	termWidth  int
	termHeight int
	hasVisited [][]bool
)

// Retrieves the current terminal dimensions
func GetTerminalSize() (width, height int) {
	width, height = 80, 24

	if os.Getenv("TERM") != "" {
		cmd := exec.Command("stty", "size")
		cmd.Stdin = os.Stdin
		out, err := cmd.Output()
		if err == nil {
			fmt.Sscanf(string(out), "%d %d", &height, &width)
		}
	}

	return width, height
}
