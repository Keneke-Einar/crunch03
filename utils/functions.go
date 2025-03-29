package utils

import (
	"fmt"
	"time"
)

// Input: reads grid dimensions and initializes game map
func Input() {
	fmt.Println("Enter the dimensions (height width):")
	fmt.Scanf("%d %d\n", &h, &w)

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

// RunGame: runs the simulation loop until no live cells remain
func RunGame() {
	for {
		PrintMap()

		if CountLiveCells() == 0 {
			break
		}

		UpdateMap()

		time.Sleep(time.Duration(delay * int(time.Millisecond)))
	}
}

// PrintMap: clears the console and prints the game grid
func PrintMap() {
	ClearConsole()

	if PassedFlag["verbose"] {
		fmt.Printf(`Tick: %v
Grid Size: %vx%v
Live Cells: %v
DelayMs: %v

`, tick, w, h, CountLiveCells(), delay)
	}
	for _, row := range gameMap {
		for _, char := range row {
			fmt.Print(charMap[char])
		}
		fmt.Println("")
	}

	tick++
}

// UpdateMap: applies game rules and updates the grid
func UpdateMap() {
	// need a new map for proper count
	newMap := make([][]rune, h)
	for i := range newMap {
		newMap[i] = make([]rune, w)
	}

	// Iterate over each cell
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			n := CountNeighbors(i, j) // count the number of neighbors of that cell

			if gameMap[i][j] == '#' { // if it's a live cell
				if n > 3 || n < 2 {
					newMap[i][j] = '.' // kill
				} else {
					newMap[i][j] = '#' // lives further
				}
			} else if gameMap[i][j] == '.' { // if it's a dead cell
				if n == 3 {
					newMap[i][j] = '#' // make it live
				} else {
					newMap[i][j] = '.' // still dead
				}
			} else {
				newMap[i][j] = gameMap[i][j] // if it's a trace
			}
		}
	}

	gameMap = newMap // update the global map
}

// PrintHelp: displays usage instructions for the program
func PrintHelp() {
	fmt.Println(`Usage: go run main.go [options]

Options:
  --help			: Show this message and exit
  --verbose			: Display the tick number, grid size, delay time, and the number of living cells
  --delay-ms=DELAY              : Set the delay time in milliseconds (accepts only integer values). Default is 2500`)
}
