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

	if PassedFlag["verbose"] {
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
			if PassedFlag["colored"] {
				switch char {
				case '#':
					fmt.Print("\033[31m" + charMap[char] + "\033[0m") // Red for live cells
				case 'o':
					fmt.Print("\033[34m" + charMap[char] + "\033[0m") // Blue for footprints
				default:
					fmt.Print(charMap[char])
				}
			} else {
				fmt.Print(charMap[char])
			}
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

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			// count the number of neighbors of that cell
			n := CountNeighbors(i, j)

			if gameMap[i][j] == '#' { // if it's a live cell
				if n > 3 || n < 2 {
					newMap[i][j] = '.' // kill
				} else {
					newMap[i][j] = '#' // lives further
				}
			} else if gameMap[i][j] == '.' || gameMap[i][j] == 'o' { // if it's a dead cell or footprint
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

	if PassedFlag["footprints"] {
		newMap = SaveFootprints(gameMap, newMap)
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

func PrintHelp() {
	fmt.Println(`Usage: go run main.go [options]

Options:
  --help			: Show this message and exit
  --verbose			: Display the tick number, grid size, delay time, and the number of living cells
  --delay-ms=DELAY	: Set the delay time in milliseconds (accepts only integer values). Default is 2500
  --footprints  	: Add traces of visited cells, displayed as '∘'
  --colored     : Add color to live cells and traces if footprints are enabled`)
}
