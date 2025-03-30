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

// Clears the console screen
func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}

// Counts the number of live cells in the game map
func CountLiveCells() int {
	count := 0

	for _, row := range gameMap {
		for _, char := range row {
			if char == '#' {
				count++
			}
		}
	}

	return count
}

// Counts the number of live neighbors for a specific cell
func CountNeighbors(row, col int) int {
	count := 0

	directions := [8][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 1},
		{1, -1},
		{1, 1},
	}

	for _, offset := range directions {
		ni, nj := row+offset[0], col+offset[1]

		if Config.EdgesPortal {
			if ni < 0 {
				ni = h - 1
			} else if ni >= h {
				ni = 0
			}

			if nj < 0 {
				nj = w - 1
			} else if nj >= w {
				nj = 0
			}
		} else {
			if ni < 0 || ni >= h || nj < 0 || nj >= w {
				continue
			}
		}

		if gameMap[ni][nj] == '#' {
			count++
		}
	}
	return count
}

// Initializes the tracking of visited cells for footprints
func InitializeFootprints() {
	hasVisited = make([][]bool, h)
	for i := range hasVisited {
		hasVisited[i] = make([]bool, w)
	}
}

// Determines the display string for a cell based on its state and configuration
func GetCellDisplay(cell rune, row, col int) string {
	var display string

	if cell == '#' {
		display = charMap[cell]
		if Config.Footprints {
			hasVisited[row][col] = true
		}

		if Config.Colored {
			return Cyan + display + Reset
		}
	} else if cell == '.' && Config.Footprints && hasVisited[row][col] {
		display = charMap['o']

		if Config.Colored {
			return Yellow + display + Reset
		}
	} else {
		display = charMap[cell]
	}

	return display
}
