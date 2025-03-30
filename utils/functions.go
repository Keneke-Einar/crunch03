package utils

import (
	"bufio"
	"fmt"
	"os" // Added for file handling
	"time"
)

// Input: reads grid dimensions and initializes game map.
func Input() {
	if Config.Random != "" {
		GenerateRandomMap(Config.Random)
		if Config.Footprints {
			InitializeFootprints()
		}
		return
	}

	// Set default delay if not specified
	if Config.Delay == 0 {
		Config.Delay = 2500
	}

	// Initialize terminal size if fullscreen is enabled
	if Config.Fullscreen {
		termWidth, termHeight = GetTerminalSize()
	}

	if Config.File != "" {
		readFromFile()
	} else {
		var originalH, originalW int
		fmt.Println("Enter the dimensions (height width):")
		fmt.Scanf("%d %d\n", &originalH, &originalW)

		// Check if fullscreen is enabled and adjust dimensions
		h, w = originalH, originalW
		if Config.Fullscreen {
			effectiveHeight := termHeight
			if Config.Verbose {
				effectiveHeight -= 5 // Reserve space for verbose output
			}

			// Adjust dimensions to be at least the terminal size
			if h < effectiveHeight {
				h = effectiveHeight
			}
			if w < termWidth {
				w = termWidth
			}
		}

		// Create game map with adjusted dimensions
		gameMap = make([][]rune, h)
		for i := range gameMap {
			gameMap[i] = make([]rune, w)
			// Initialize all cells as dead
			for j := range gameMap[i] {
				gameMap[i][j] = '.'
			}
		}

		// Determine the number of rows to read from input (originalH, up to adjusted h)
		inputHeight := originalH
		if inputHeight > h {
			inputHeight = h
		}

		// Initialize visited cells tracking if footprints enabled
		if Config.Footprints {
			InitializeFootprints()
		}

		// Read the actual input grid up to inputHeight rows
		for i := 0; i < inputHeight; i++ {
			rowInput := ""
			fmt.Scanf("%s\n", &rowInput)

			// Copy input to game map, up to adjusted width (w)
			for j, char := range rowInput {
				if j < w {
					gameMap[i][j] = char
				}
			}
		}
	}
}

// readFromFile: reads the grid from the file specified in Config.File
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

	// Check if fullscreen is enabled and adjust dimensions
	h, w = originalH, originalW
	if Config.Fullscreen {
		effectiveHeight := termHeight
		if Config.Verbose {
			effectiveHeight -= 5 // Reserve space for verbose output
		}
		if h < effectiveHeight {
			h = effectiveHeight
		}
		if w < termWidth {
			w = termWidth
		}
	}

	// Create game map with adjusted dimensions
	gameMap = make([][]rune, h)
	for i := range gameMap {
		gameMap[i] = make([]rune, w)
		for j := range gameMap[i] {
			gameMap[i][j] = '.' // Initialize all cells as dead
		}
	}

	// Determine the number of rows to read from file
	inputHeight := originalH
	if inputHeight > h {
		inputHeight = h
	}

	// Initialize visited cells tracking if footprints enabled
	if Config.Footprints {
		InitializeFootprints()
	}

	// Read the grid from the file
	for i := 0; i < inputHeight && scanner.Scan(); i++ {
		rowInput := scanner.Text()
		for j, char := range rowInput {
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

// RunGame: runs the simulation loop until no live cells remain.
func RunGame() {
	for {
		PrintMap()

		if CountLiveCells() == 0 {
			break
		}

		UpdateMap()

		// Use Config.Delay instead of a global delay variable.
		time.Sleep(time.Duration(Config.Delay) * time.Millisecond)
	}
}

// PrintMap: clears the console and prints the game grid.
func PrintMap() {
	ClearConsole()

	// Use the structured Config to check if verbose output is enabled.
	if Config.Verbose {
		fmt.Printf(`Tick: %v
Grid Size: %vx%v
Live Cells: %v
DelayMs: %vms

`, tick, w, h, CountLiveCells(), Config.Delay)
	}

	// Print the game map with appropriate formatting
	for i, row := range gameMap {
		for j, char := range row {
			fmt.Print(GetCellDisplay(char, i, j))
		}
		fmt.Println("")
	}

	tick++
}

// UpdateMap: applies game rules and updates the grid.
func UpdateMap() {
	// Create a new map to update cell states.
	newMap := make([][]rune, h)
	for i := range newMap {
		newMap[i] = make([]rune, w)
	}

	// Iterate over each cell to apply game rules.
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			n := CountNeighbors(i, j) // count neighbors of cell [i][j]

			if gameMap[i][j] == '#' { // live cell
				if n > 3 || n < 2 {
					newMap[i][j] = '.' // cell dies
				} else {
					newMap[i][j] = '#' // cell lives on
				}
			} else if gameMap[i][j] == '.' { // dead cell
				if n == 3 {
					newMap[i][j] = '#' // cell becomes alive
				} else {
					newMap[i][j] = '.' // remains dead
				}
			} else {
				newMap[i][j] = gameMap[i][j] // for trace or other states
			}
		}
	}

	gameMap = newMap // update the global grid state
}

// PrintHelp: displays usage instructions for the program.
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
