package utils

import (
	"fmt"
	"os"
	"os/exec"
)

var charMap map[rune]string = map[rune]string{
	'#': "×", // live cell
	'.': "·", // dead cell
	'o': "∘", // trace of a cell when footprints enabled
}

// ANSI color codes
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
	hasVisited [][]bool // To track which cells have been alive
)

// GetTerminalSize: gets the current terminal dimensions
func GetTerminalSize() (width, height int) {
	// Default fallback values
	width, height = 80, 24

	// Only implemented on Unix-like systems
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

// ClearConsole: clears the console
func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}

// CountLiveCells: counts the number of live cells
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

// CountNeighbors: counts the number of live neighbors of a cell
func CountNeighbors(row, col int) int {
	count := 0

	directions := [8][2]int{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, -1}, // Upper-left diagonal
		{-1, 1},  // Upper-right diagonal
		{1, -1},  // Lower-left diagonal
		{1, 1},   // Lower-right diagonal
	}

	for _, d := range directions {
		nRow, nCol := row+d[0], col+d[1]
		if nRow >= 0 && nRow < h && nCol >= 0 && nCol < w && gameMap[nRow][nCol] == '#' {
			count++
		}
	}

	return count
}

// InitializeFootprints: initializes the visited cells tracking
func InitializeFootprints() {
	hasVisited = make([][]bool, h)
	for i := range hasVisited {
		hasVisited[i] = make([]bool, w)
	}
}

// GetCellDisplay: returns the appropriate display for a cell based on its state and config
func GetCellDisplay(cell rune, row, col int) string {
	var display string

	// If the cell is currently alive
	if cell == '#' {
		display = charMap[cell]
		// Mark as visited for future footprints
		if Config.Footprints {
			hasVisited[row][col] = true
		}

		// Apply color if enabled
		if Config.Colored {
			return Cyan + display + Reset
		}
	} else if cell == '.' && Config.Footprints && hasVisited[row][col] {
		// Cell is dead but was previously alive (footprint)
		display = charMap['o']

		// Apply color if both footprints and colored are enabled
		if Config.Colored {
			return Yellow + display + Reset
		}
	} else {
		// Normal dead cell
		display = charMap[cell]
	}

	return display
}
