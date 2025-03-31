package utils

import (
	"fmt"
	"os"
	"strings"
)

// Processes command-line arguments into configuration flags
func ParseFlags() (int, error) {
	// all passed arguments except the filename
	args := os.Args[1:]

	for _, arg := range args {
		flagName, flagValue, valid := extractFlag(arg)
		if !valid {
			return 0, fmt.Errorf("error: invalid argument '%s'", arg)
		}

		flag, found := findFlag(flagName)
		if !found {
			return 0, fmt.Errorf("error: unknown flag '--%s'", flagName)
		}

		if err := processFlag(flag, flagValue); err != nil {
			return 0, fmt.Errorf("error processing flag '--%s': %s", flagName, err)
		}
	}

	if Config.UseUnicode {
		charMap['#'] = "⬛"
		charMap['.'] = "⬜"
		charMap['o'] = "🟨"
	}

	return len(args), nil
}

// Extracts flag name and value from an argument string
func extractFlag(arg string) (string, string, bool) {
	if !strings.HasPrefix(arg, "--") {
		return "", "", false
	}

	content := arg[2:]
	if strings.Contains(content, "=") {
		parts := strings.SplitN(content, "=", 2)
		return parts[0], parts[1], true
	}
	return content, "", true
}

// Finds a flag definition by its name
func findFlag(flagName string) (Flag, bool) {
	for _, f := range flags {
		if f.Name == flagName {
			return f, true
		}
	}
	return Flag{}, false
}

// Applies the processing function for the found flag
func processFlag(flag Flag, value string) error {
	if flag.HasValue && value == "" {
		return fmt.Errorf("flag '--%s' requires a value", flag.Name)
	}
	if !flag.HasValue && value != "" {
		return fmt.Errorf("flag '--%s' does not accept a value", flag.Name)
	}
	return flag.Process(value)
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
  --file=FILENAME	: Load the initial grid from a specified file
  --edges-portal	: Enable portal edges where cells that exit the
				grid appear on the opposite side
  --random=WxH		: Generate a random grid of the specified width (W)
				and height (H)
  --fullscreen		: Adjust the grid to fit the terminal size with
				empty cells
  --footprints		: Add traces of visited cells, displayed as '∘'
  --colored		: Add color to live cells and traces if footprints
				are enabled
  --use-unicode		: Use unicode characters to display the cells.
				(not intended to be used with --colored)
  --template=TEMPLATE	: Load one of the existing templates:
  				3g-hwss - heavyweight spaceship from 3 gliders,
				3g-mwss - middleweight spaceship from 3 gliders,
				acorn - a small pattern that grows big,
				crab - a crab,
				pentadecathlon - period 15 oscillator,
				pulsar - period 3 oscillator,
				toad - period 2 oscillator`)
}
