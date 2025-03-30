package utils

import (
	"fmt"
)

var charMap map[rune]string = map[rune]string{
	'#': "×", // live cell
	'.': "·", // dead cell
	'o': ".", // trace of a cell. By default, looks the same as the dead cell, but with the footprints flag, it changes to "∘"
}

var (
	h, w    int
	tick    int = 1
	gameMap [][]rune
)

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
