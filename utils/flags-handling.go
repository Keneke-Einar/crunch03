package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Config struct {
	Help        bool
	Verbose     bool
	Delay       int
	Random      string
	Footprints  bool
	Colored     bool
	Fullscreen  bool
	EdgesPortal bool
	File        string
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
	if err != nil {
		return fmt.Errorf("invalid delay value: %s", value)
	}
	Config.Delay = d
	return nil
}

// Sets the Random configuration value and generates a random map
func processRandom(value string) error {
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
	Config.File = value
	return nil
}

// Enables portal behavior for map edges
func processEdgesPortal(value string) error {
	Config.EdgesPortal = true
	return nil
}

// Processes command-line arguments into configuration flags
func ParseFlags() {
	args := os.Args[1:]

	for _, arg := range args {
		flagName, flagValue, valid := extractFlag(arg)
		if !valid {
			fmt.Printf("Error: invalid argument '%s'\n", arg)
			os.Exit(1)
		}

		flag, found := findFlag(flagName)
		if !found {
			fmt.Printf("Error: unknown flag '--%s'\n", flagName)
			os.Exit(1)
		}

		if err := processFlag(flag, flagValue); err != nil {
			fmt.Printf("Error processing flag '--%s': %s\n", flagName, err.Error())
			os.Exit(1)
		}
	}
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

// Applies the processing function for a found flag
func processFlag(flag Flag, value string) error {
	if flag.HasValue && value == "" {
		return fmt.Errorf("flag '--%s' requires a value", flag.Name)
	}
	if !flag.HasValue && value != "" {
		return fmt.Errorf("flag '--%s' does not accept a value", flag.Name)
	}
	return flag.Process(value)
}
