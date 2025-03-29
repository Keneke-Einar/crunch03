package utils

import (
	"fmt"
	"time"
)

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

func PrintMap() {
	// clear the console
	ClearConsole()

	if passedFlag["verbose"] {
		// print the statistics
		fmt.Printf(`Tick: %v
Grid Size: %vx%v
Live Cells: %v
DelayMs: %v

`, tick, w, h, CountLiveCells(), delay)
	}

	// print the map
	for _, row := range gameMap {
		for _, char := range row {
			fmt.Print(charMap[char])
		}
		fmt.Println("")
	}

	tick++
}

func UpdateMap() {
	// need a new map for proper count
	newMap := make([][]rune, h)
	for i := range newMap {
		newMap[i] = make([]rune, w)
	}

	for i := range h {
		for j := range w {
			// count the number of neighbors of that cell
			n := CountNeighbors(i, j)

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
