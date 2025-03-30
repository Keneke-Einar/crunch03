package utils

import (
	"fmt"
)

var charMap map[rune]string = map[rune]string{
	'#': "×", // live cell
	'.': "·", // dead cell
	'o': "∘", // trace of a cell. By default, looks the same as the dead cell, but with the footprints flag, it changes to "∘"
}

var (
	h, w    int
	tick    int = 1
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
