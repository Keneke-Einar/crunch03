package utils

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Reads grid dimensions and initializes the game map
func Input() {
	if Config.Random != "" {
		GenerateRandomMap(Config.Random)
		if Config.Footprints {
			InitializeFootprints()
		}
		return
	}

	if Config.Delay == 0 {
		Config.Delay = 2500
	}

	if Config.Fullscreen {
		termWidth, termHeight = GetTerminalSize()
	}

	if Config.File != "" {
		readFromFile()
	} else {
		var originalH, originalW int
		fmt.Println("Enter the dimensions (height width):")
		_, err := fmt.Scanf("%d %d\n", &originalH, &originalW)
		if err != nil {
			fmt.Println("Error: invalid dimension format. Please enter two integers separated by space.")
			os.Exit(1)
		}

		h, w = originalH, originalW
		if Config.Fullscreen {
			effectiveHeight := termHeight
			if Config.Verbose {
				effectiveHeight -= 5
			}

			if h < effectiveHeight {
				h = effectiveHeight
			}
			if w < termWidth {
				w = termWidth
			}
		}

		gameMap = make([][]rune, h)
		for i := range gameMap {
			gameMap[i] = make([]rune, w)
			for j := range gameMap[i] {
				gameMap[i][j] = '.'
			}
		}

		inputHeight := originalH
		if inputHeight > h {
			inputHeight = h
		}

		if Config.Footprints {
			InitializeFootprints()
		}

		for i := 0; i < inputHeight; i++ {
			rowInput := ""
			_, err := fmt.Scanf("%s\n", &rowInput)
			if err != nil {
				fmt.Println("Error: failed to read row input")
				os.Exit(1)
			}

			if len(rowInput) != originalW {
				fmt.Println("Error: row length does not match specified width")
				os.Exit(1)
			}

			for j, char := range rowInput {
				if char != '.' && char != '#' {
					fmt.Println("Error: grid can only contain '.' and '#' characters")
					os.Exit(1)
				}
				if j < w {
					gameMap[i][j] = char
				}
			}
		}
	}
}

// Reads the game grid from a specified file
func readFromFile() {
	file, err := os.Open(Config.File)
	if err != nil {
		fmt.Println("Error: cannot open file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var originalH, originalW int
	if scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &originalH, &originalW)
		if err != nil {
			fmt.Println("Error: invalid dimensions in file")
			os.Exit(1)
		}
	}

	if originalH < 3 || originalW < 3 {
		fmt.Println("Error: grid dimensions must be at least 3x3")
		os.Exit(1)
	}

	h, w = originalH, originalW
	if Config.Fullscreen {
		effectiveHeight := termHeight
		if Config.Verbose {
			effectiveHeight -= 5
		}
		if h < effectiveHeight {
			h = effectiveHeight
		}
		if w < termWidth {
			w = termWidth
		}
	}

	gameMap = make([][]rune, h)
	for i := range gameMap {
		gameMap[i] = make([]rune, w)
		for j := range gameMap[i] {
			gameMap[i][j] = '.'
		}
	}

	inputHeight := originalH
	if inputHeight > h {
		inputHeight = h
	}

	if Config.Footprints {
		InitializeFootprints()
	}

	for i := 0; i < inputHeight && scanner.Scan(); i++ {
		rowInput := scanner.Text()

		if len(rowInput) != originalW {
			fmt.Println("Error: row length in file does not match specified width")
			os.Exit(1)
		}

		for j, char := range rowInput {
			if char != '.' && char != '#' {
				fmt.Println("Error: grid in file can only contain '.' and '#' characters")
				os.Exit(1)
			}
			if j < w {
				gameMap[i][j] = char
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error: reading file:", err)
		os.Exit(1)
	}
}

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

// Displays usage instructions for the program
func PrintHelp() {
	fmt.Println(`Usage: go run main.go [options]

Options:
  --help            : Show this message and exit
  --verbose         : Display tick number, grid size, delay time, and live cell count
  --delay-ms=DELAY  : Set the delay time in milliseconds (accepts only integer values). Default is 2500
  --random=WxH      : Generate a random grid of the specified width (W) and height (H)
  --footprints      : Add traces of visited cells, displayed as '∘'
  --colored         : Add color to live cells and traces if footprints are enabled
  --fullscreen      : Adjust the grid to fit the terminal size with empty cells
  --file=X          : Load the initial grid from a specified file
`)
}
