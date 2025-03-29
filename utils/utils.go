package utils

import (
	"fmt"
	"time"
)

var charMap map[rune]string = map[rune]string{
	'#': "×", // live cell
	'.': "·", // dead cell
	'o': ".", // trace of a cell. By default, looks the same as the dead cell, but with the footprints flag, it changes to "∘"
}

var (
	h, w    int
	tick    int = 1
	delay   int = 2500
	gameMap [][]rune
)

func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}

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

	// print the statistics
	fmt.Printf(`Tick: %v
Grid Size: %vx%v
Live Cells: %v
DelayMs: %v
`, tick, w, h, CountLiveCells(), delay)

	// print the map
	for _, row := range gameMap {
		for _, char := range row {
			fmt.Print(charMap[char])
		}
		fmt.Println("")
	}

	tick++
}

func CountNeighbors(row, col int) int {
	count := 0

	// upper neighbor
	if row-1 >= 0 && gameMap[row-1][col] == '#' {
		count++
	}

	// lower neighbor
	if row+1 < h && gameMap[row+1][col] == '#' {
		count++
	}

	// neighbor on the right
	if col+1 < w && gameMap[row][col+1] == '#' {
		count++
	}

	// neighbor on the left
	if col-1 >= 0 && gameMap[row][col-1] == '#' {
		count++
	}

	// upper left neighbor
	if row-1 >= 0 && col-1 >= 0 && gameMap[row-1][col-1] == '#' {
		count++
	}

	// upper right neighbor
	if row-1 >= 0 && col+1 < w && gameMap[row-1][col+1] == '#' {
		count++
	}

	// lower left neighbor
	if row+1 < h && col-1 >= 0 && gameMap[row+1][col-1] == '#' {
		count++
	}

	// lower right neighbor
	if row+1 < h && col+1 < w && gameMap[row+1][col+1] == '#' {
		count++
	}

	return count
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
