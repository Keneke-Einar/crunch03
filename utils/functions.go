package utils

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Reads grid dimensions and initializes the game map
func Input() error {
	if Config.Random != "" {
		err := GenerateRandomMap(Config.Random)
		if err != nil {
			return err
		}

		if Config.Footprints {
			InitializeFootprints()
		}

		return nil
	}

	if Config.Delay == 0 {
		Config.Delay = 2500
	}

	if Config.Fullscreen {
		termWidth, termHeight = GetTerminalSize()
	}

	if Config.File != "" {
		err := readFromFile()
		if err != nil {
			return err
		}
	} else {
		var originalH, originalW int
		fmt.Println("Enter the dimensions (height width):")
		_, err := fmt.Scanf("%d %d\n", &originalH, &originalW)
		if err != nil {
			return fmt.Errorf("Error: invalid dimension format. Please enter two integers separated by space.")
		}

		if originalW < 3 || originalH < 3 {
			return fmt.Errorf("invalid grid size. Minimum size is 3x3")
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
			if _, err := fmt.Scanf("%s\n", &rowInput); err != nil {
				return fmt.Errorf("Error: failed to read row input")
			}

			if len(rowInput) != originalW {
				return fmt.Errorf("Error: row length does not match specified width")
			}

			for j, char := range rowInput {
				if char != '.' && char != '#' {
					return fmt.Errorf("Error: grid can only contain '.' and '#' characters")
				}
				if j < w {
					gameMap[i][j] = char
				}
			}
		}
	}

	return nil
}

// Reads the game grid from a specified file
func readFromFile() error {
	file, err := os.Open(Config.File)
	if err != nil {
		return fmt.Errorf("Error: cannot open file: %w", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var originalH, originalW int
	if scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &originalH, &originalW)
		if err != nil {
			return fmt.Errorf("Error: invalid dimensions in file")
		}
	}

	if originalH < 2 || originalW < 2 {
		return fmt.Errorf("Error: invalid grid size %dx%d. Minimum size is 2x2", originalH, originalW)
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
			return fmt.Errorf("Error: row length in file does not match specified width")
		}

		for j, char := range rowInput {
			if char != '.' && char != '#' {
				return fmt.Errorf("Error: grid in file can only contain '.' and '#' characters")
			}
			if j < w {
				gameMap[i][j] = char
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error: reading file: %w", err)
	}

	return nil
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
  --help        : Show the help message and exit
  --verbose     : Display detailed information about the simulation, including grid size, number of ticks, speed, and map name
  --delay-ms=X  : Set the animation speed in milliseconds. Default is 2500 milliseconds
  --file=X      : Load the initial grid from a specified file
  --edges-portal: Enable portal edges where cells that exit the grid appear on the opposite side
  --random=WxH  : Generate a random grid of the specified width (W) and height (H)
  --fullscreen  : Adjust the grid to fit the terminal size with empty cells
  --footprints  : Add traces of visited cells, displayed as '∘'
  --colored     : Add color to live cells and traces if footprints are enabled
`)
}
