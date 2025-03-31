package utils

import (
	"fmt"
	"strconv"
	"strings"
)

var Config struct {
	Colored     bool
	Fullscreen  bool
	Footprints  bool
	EdgesPortal bool
	Help        bool
	Verbose     bool
	Delay       int
	File        string
	Random      string
}

type Flag struct {
	Name     string
	HasValue bool
	Process  func(value string) error
}

var flags = []Flag{
	{Name: "help", HasValue: false, Process: processHelp},
	{Name: "verbose", HasValue: false, Process: processVerbose},
	{Name: "delay-ms", HasValue: true, Process: processDelay},
	{Name: "random", HasValue: true, Process: processRandom},
	{Name: "footprints", HasValue: false, Process: processFootprints},
	{Name: "colored", HasValue: false, Process: processColored},
	{Name: "fullscreen", HasValue: false, Process: processFullscreen},
	{Name: "edges-portal", HasValue: false, Process: processEdgesPortal},
	{Name: "file", HasValue: true, Process: processFile},
	{Name: "template", HasValue: true, Process: processTemplate},
}

// Sets the Help configuration flag
func processHelp(value string) error {
	Config.Help = true
	return nil
}

// Sets the Verbose configuration flag
func processVerbose(value string) error {
	Config.Verbose = true
	return nil
}

// Sets the Delay configuration value
func processDelay(value string) error {
	d, err := strconv.Atoi(value)
	if err != nil || d <= 0 {
		return fmt.Errorf("invalid delay value: %s, expected a positive integer", value)
	}
	Config.Delay = d
	return nil
}

// Sets the Random configuration value and generates a random map
func processRandom(value string) error {
	if Config.File != "" {
		return nil // don't implement if --file or --template is already implemented
	}

	parts := strings.Split(value, "x")
	if len(parts) != 2 {
		return fmt.Errorf("invalid format for --random, expected WxH (e.g., 5x5)")
	}

	width, err1 := strconv.Atoi(parts[0])
	height, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil || width <= 0 || height <= 0 {
		return fmt.Errorf("invalid dimensions for --random, expected positive integers (e.g., 5x5)")
	}

	Config.Random = value
	GenerateRandomMap(value)
	return nil
}

// Sets the Footprints configuration flag
func processFootprints(value string) error {
	Config.Footprints = true
	return nil
}

// Sets the Colored configuration flag
func processColored(value string) error {
	Config.Colored = true
	return nil
}

// Sets the Fullscreen configuration flag
func processFullscreen(value string) error {
	Config.Fullscreen = true
	return nil
}

// Sets the File configuration value
func processFile(value string) error {
	if Config.Random != "" || Config.File != "" {
		return nil // don't implement if --random or --template is already implemented
	}

	Config.File = value
	return nil
}

// Enables portal behavior for map edges
func processEdgesPortal(value string) error {
	Config.EdgesPortal = true
	return nil
}

func processTemplate(value string) error {
	if Config.Random != "" || Config.File != "" {
		return nil // don't implement if --random or --file is already implemented
	}

	if err := findTemplate(value); err != nil {
		return err
	}

	Config.File = "utils/templates/" + value + ".txt"
	return nil
}

// Checks if the template exists in the library. If it does, returns nil. If it doesn't, returns an error.
func findTemplate(value string) error {
	templates := []string{"3g-hwss", "3g-mwss", "acorn", "crab", "pentadecathlon", "pulsar", "toad"}

	for _, template := range templates {
		if template == value {
			return nil
		}
	}

	return fmt.Errorf("template %s doesn't exist", value)
}

// Displays usage instructions for the program
func PrintHelp() {
	fmt.Println(`Usage: go run main.go [options]

Options:
  --help		: Show the help message and exit
  --verbose		: Display detailed information about the simulation,
				including grid size, number of ticks, speed,
				and map name
  --delay-ms=X		: Set the animation speed in milliseconds.
				Default is 2500 milliseconds
  --file=X		: Load the initial grid from a specified file
  --edges-portal	: Enable portal edges where cells that exit the
				grid appear on the opposite side
  --random=WxH		: Generate a random grid of the specified width (W)
				and height (H)
  --fullscreen		: Adjust the grid to fit the terminal size with
				empty cells
  --footprints		: Add traces of visited cells, displayed as '∘'
  --colored		: Add color to live cells and traces if footprints
				are enabled`)
}
