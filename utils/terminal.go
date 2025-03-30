package utils

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// GetTerminalSize returns the width and height of the terminal
func GetTerminalSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 80, 24 // default size if unable to determine
	}

	parts := strings.Split(strings.TrimSpace(string(out)), " ")
	if len(parts) != 2 {
		return 80, 24
	}

	height, err := strconv.Atoi(parts[0])
	if err != nil {
		height = 24
	}

	width, err := strconv.Atoi(parts[1])
	if err != nil {
		width = 80
	}

	// Account for potential status information display (stats line)
	if PassedFlag["verbose"] {
		height = height - 6 // Reduce height to account for verbose information (4 lines + 2 margin)
	}

	return width, height
}

// AdjustGridToTerminal expands the current grid to fit the terminal size
func AdjustGridToTerminal() {
	termWidth, termHeight := GetTerminalSize()

	// Ensure we don't exceed terminal dimensions
	if h > termHeight {
		h = termHeight
		// Trim game map if needed
		if len(gameMap) > h {
			gameMap = gameMap[:h]
		}
	}

	if w > termWidth {
		w = termWidth
		// Trim each row if needed
		for i := range gameMap {
			if len(gameMap[i]) > w {
				gameMap[i] = gameMap[i][:w]
			}
		}
	}

	// Expand grid if terminal is larger than current grid
	if h < termHeight {
		// Add rows until we reach terminal height
		for len(gameMap) < termHeight {
			newRow := make([]rune, w)
			for i := range newRow {
				newRow[i] = '.'
			}
			gameMap = append(gameMap, newRow)
		}
		h = termHeight
	}

	if w < termWidth {
		// Expand each row to terminal width
		for i := range gameMap {
			for len(gameMap[i]) < termWidth {
				gameMap[i] = append(gameMap[i], '.')
			}
		}
		w = termWidth
	}
}
