package utils

import (
	"fmt"
	"time"
)

// Runs the game simulation loop until no live cells remain
func RunGame() {
	for {
		printMap()

		if countLiveCells() == 0 {
			break
		}

		updateMap()

		time.Sleep(time.Duration(Config.Delay) * time.Millisecond)
	}
}

// Clears the console and prints the current game grid
func printMap() {
	clearConsole()

	if Config.Verbose {
		fmt.Printf(`Tick: %v
Grid Size: %vx%v
Live Cells: %v
DelayMs: %v ms

`, tick, w, h, countLiveCells(), Config.Delay)
	}

	for i, row := range gameMap {
		for j, char := range row {
			fmt.Print(getCellDisplay(char, i, j))
		}
		fmt.Println("")
	}

	tick++
}

// Applies game rules to update the grid state
func updateMap() {
	newMap := make([][]rune, h)
	for i := range newMap {
		newMap[i] = make([]rune, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			n := countNeighbors(i, j)

			if gameMap[i][j] == '#' {
				if n > 3 || n < 2 {
					newMap[i][j] = '.'
				} else {
					newMap[i][j] = '#'
				}
			} else {
				if n == 3 {
					newMap[i][j] = '#'
				} else {
					newMap[i][j] = '.'
				}
			}
		}
	}

	gameMap = newMap
}

//================Inner Utils================

// Clears the console screen
func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

// Counts the number of live neighbors for a specific cell
func countNeighbors(row, col int) int {
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

// Determines the display string for a cell based on its state and configuration
func getCellDisplay(cell rune, row, col int) string {
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

// Counts the number of live cells in the game map
func countLiveCells() int {
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
