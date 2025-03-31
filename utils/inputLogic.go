package utils

import (
	"bufio"
	"fmt"
	"os"
)

// Reads grid dimensions and initializes the game map
func Input() error {
	if Config.Random != "" {
		if err := GenerateRandomMap(Config.Random); err != nil {
			return err
		}
		if Config.Footprints {
			initializeFootprints()
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
		return readFromFile()
	}

	return readFromUserInput()
}

// Reads dimensions and initializes the game map
func readFromUserInput() error {
	var originalH, originalW int
	fmt.Println("Enter the dimensions (height width):")
	if _, err := fmt.Scanf("%d %d\n", &originalH, &originalW); err != nil {
		return fmt.Errorf("error: invalid dimension format. Please enter two integers separated by space")
	}

	if err := validateGridSize(originalH, originalW); err != nil {
		return err
	}

	initializeGrid(originalH, originalW)
	return populateGridFromUser(originalH, originalW)
}

// Reads the game grid from a specified file
func readFromFile() error {
	file, err := os.Open(Config.File)
	if err != nil {
		return fmt.Errorf("error: cannot open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var originalH, originalW int
	if scanner.Scan() {
		if _, err := fmt.Sscanf(scanner.Text(), "%d %d", &originalH, &originalW); err != nil {
			return fmt.Errorf("error: invalid dimensions in file")
		}
	}

	if err := validateGridSize(originalH, originalW); err != nil {
		return err
	}

	initializeGrid(originalH, originalW)
	return populateGridFromFile(scanner, originalH, originalW)
}

// Initializes the game grid based on dimensions
func initializeGrid(originalH, originalW int) {
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

	if Config.Footprints {
		initializeFootprints()
	}
}

// Reads grid input from a file
func populateGridFromFile(scanner *bufio.Scanner, originalH, originalW int) error {
	for i := 0; i < originalH && scanner.Scan(); i++ {
		rowInput := scanner.Text()
		if err := validateRowInput(rowInput, originalW); err != nil {
			return err
		}

		for j, char := range rowInput {
			gameMap[i][j] = char
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error: reading file: %w", err)
	}
	return nil
}

// Validates a row of grid input
func validateRowInput(rowInput string, expectedWidth int) error {
	if len(rowInput) != expectedWidth {
		return fmt.Errorf("error: row length does not match specified width")
	}
	for _, char := range rowInput {
		if char != '.' && char != '#' {
			return fmt.Errorf("error: grid can only contain '.' and '#' characters")
		}
	}
	return nil
}

// Reads grid input from the user
func populateGridFromUser(originalH, originalW int) error {
	fmt.Println("Enter the grid; use only '#' and '.' for live and dead cells, respectively:")

	for i := 0; i < originalH; i++ {
		var rowInput string
		if _, err := fmt.Scanf("%s\n", &rowInput); err != nil {
			return fmt.Errorf("error: failed to read row input")
		}

		if err := validateRowInput(rowInput, originalW); err != nil {
			return err
		}

		for j, char := range rowInput {
			gameMap[i][j] = char
		}
	}
	return nil
}

// Validates the grid size
func validateGridSize(h, w int) error {
	if h < 3 || w < 3 {
		return fmt.Errorf("error: invalid grid size %dx%d. Minimum size is 3x3", h, w)
	}
	return nil
}

//======================Inner Utils======================

// Initializes the tracking of visited cells for footprints
func initializeFootprints() {
	hasVisited = make([][]bool, h)
	for i := range hasVisited {
		hasVisited[i] = make([]bool, w)
	}
}
