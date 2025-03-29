package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// Input: reads grid dimensions and initializes game map
func Input() {
	if PassedFlag["random"] {
		GenerateRandomMap()
		return
	}

	fmt.Println("Enter the dimensions (height width):")
	fmt.Scanf("%d %d\n", &h, &w)
	fmt.Scanln()

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

func GenerateRandomMap() {
	// Parse `randomDimensions` (which is stored as WxH)
	var wVal, hVal int
	fmt.Sscanf(randomDimensions, "%dx%d", &wVal, &hVal)

	// Assign parsed values to global dimensions
	h, w = hVal, wVal
	gameMap = make([][]rune, h)

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < h; i++ {
		gameMap[i] = make([]rune, w)
		for j := 0; j < w; j++ {
			if rand.Float64() < 0.4 { // 40% chance of being alive
				gameMap[i][j] = '#'
			} else {
				gameMap[i][j] = '.'
			}
		}
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
  --delay-ms=DELAY              : Set the delay time in milliseconds (accepts only integer values). Default is 2500
  --random=WxH  : Generate a random grid of the specified width (W) and height (H)`)
}
