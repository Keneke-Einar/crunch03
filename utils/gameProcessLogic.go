package utils

import (
	"fmt"
	"time"
)

// Runs the game simulation loop until no live cells remain
func RunGame() {
	for {
		printMap()

		if CountLiveCells() == 0 {
			break
		}

		updateMap()

		time.Sleep(time.Duration(Config.Delay) * time.Millisecond)
	}
}

// Clears the console and prints the current game grid
func printMap() {
	ClearConsole()

	if Config.Verbose {
		fmt.Printf(`Tick: %v
Grid Size: %vx%v
Live Cells: %v
DelayMs: %vms

`, tick, w, h, CountLiveCells(), Config.Delay)
	}

	for i, row := range gameMap {
		for j, char := range row {
			fmt.Print(GetCellDisplay(char, i, j))
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
			n := CountNeighbors(i, j)

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
