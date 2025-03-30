package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Runs the game simulation loop until no live cells remain
func RunGame() {
	for {
		PrintMap()

		if CountLiveCells() == 0 {
			break
		}

		UpdateMap()

		time.Sleep(time.Duration(Config.Delay) * time.Millisecond)
	}
}

// Clears the console and prints the current game grid
func PrintMap() {
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
func UpdateMap() {
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

// Generates a random game map with specified dimensions
func GenerateRandomMap(dimensions string) error {
	parts := strings.Split(dimensions, "x")
	if len(parts) != 2 {
		return fmt.Errorf("Error: invalid format for --random flag. Use --random=WxH")
	}

	width, errW := strconv.Atoi(parts[0])
	height, errH := strconv.Atoi(parts[1])

	if errW != nil || errH != nil || width <= 0 || height <= 0 {
		return fmt.Errorf("Error: invalid dimensions for --random flag. Width and height must be positive integers.")
	}

	if width < 3 || height < 3 {
		return fmt.Errorf("invalid grid size. Minimum size is 3x3")
	}

	if Config.Fullscreen {
		termWidth, termHeight = GetTerminalSize()

		effectiveHeight := termHeight
		if Config.Verbose {
			effectiveHeight -= 5
		}

		if height < effectiveHeight {
			height = effectiveHeight
		}
		if width < termWidth {
			width = termWidth
		}
	}

	w, h = width, height
	gameMap = make([][]rune, h)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < h; i++ {
		row := make([]rune, w)
		for j := 0; j < w; j++ {
			if rand.Intn(2) == 0 {
				row[j] = '#'
			} else {
				row[j] = '.'
			}
		}
		gameMap[i] = row
	}

	return nil
}
