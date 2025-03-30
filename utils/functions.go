package utils

import (
	"fmt"
	"time"
)

// Input: reads grid dimensions and initializes game map.
func Input() {
	if Config.Random != "" {
		fmt.Println("Random grid generation not implemented yet.")
		return
	}

	fmt.Println("Enter the dimensions (height width):")
	fmt.Scanf("%d %d\n", &h, &w)

	// Read rows until the grid is filled.
	for len(gameMap) < h {
		row := []rune{}
		rowInput := ""
		fmt.Scanf("%s\n", &rowInput)
		for _, char := range rowInput {
			row = append(row, char)
		}
		gameMap = append(gameMap, row)
	}
}

// RunGame: runs the simulation loop until no live cells remain.
func RunGame() {
	for {
		PrintMap()

		if CountLiveCells() == 0 {
			break
		}

		UpdateMap()

		// Use Config.Delay instead of a global delay variable.
		time.Sleep(time.Duration(Config.Delay) * time.Millisecond)
	}
}

// PrintMap: clears the console and prints the game grid.
func PrintMap() {
	ClearConsole()

	// Use the structured Config to check if verbose output is enabled.
	if Config.Verbose {
		fmt.Printf(`Tick: %v
Grid Size: %vx%v
Live Cells: %v
DelayMs: %v

`, tick, w, h, CountLiveCells(), Config.Delay)
	}
	for _, row := range gameMap {
		for _, char := range row {
			fmt.Print(charMap[char])
		}
		fmt.Println("")
	}

	tick++
}

// UpdateMap: applies game rules and updates the grid.
func UpdateMap() {
	// Create a new map to update cell states.
	newMap := make([][]rune, h)
	for i := range newMap {
		newMap[i] = make([]rune, w)
	}

	// Iterate over each cell to apply game rules.
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			n := CountNeighbors(i, j) // count neighbors of cell [i][j]

			if gameMap[i][j] == '#' { // live cell
				if n > 3 || n < 2 {
					newMap[i][j] = '.' // cell dies
				} else {
					newMap[i][j] = '#' // cell lives on
				}
			} else if gameMap[i][j] == '.' { // dead cell
				if n == 3 {
					newMap[i][j] = '#' // cell becomes alive
				} else {
					newMap[i][j] = '.' // remains dead
				}
			} else {
				newMap[i][j] = gameMap[i][j] // for trace or other states
			}
		}
	}

	gameMap = newMap // update the global grid state
}

// PrintHelp: displays usage instructions for the program.
func PrintHelp() {
	fmt.Println(`Usage: go run main.go [options]

Options:
  --help            : Show this message and exit
  --verbose         : Display tick number, grid size, delay time, and live cell count
  --delay-ms=DELAY  : Set the delay time in milliseconds (accepts only integer values). Default is 2500
  --random=WxH      : Generate a random grid of the specified width (W) and height (H)`)
}
