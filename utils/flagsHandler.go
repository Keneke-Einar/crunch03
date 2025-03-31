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
	UseUnicode  bool
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
	{Name: "verbose", HasValue: false, Process: processCheckFlag(&Config.Verbose)}, //!
	{Name: "delay-ms", HasValue: true, Process: processDelay},
	{Name: "random", HasValue: true, Process: processRandom},
	{Name: "footprints", HasValue: false, Process: processCheckFlag(&Config.Footprints)}, //!
	{Name: "colored", HasValue: false, Process: processColored},
	{Name: "fullscreen", HasValue: false, Process: processCheckFlag(&Config.Fullscreen)},    //!
	{Name: "edges-portal", HasValue: false, Process: processCheckFlag(&Config.EdgesPortal)}, //!
	{Name: "file", HasValue: true, Process: processFile},
	{Name: "template", HasValue: true, Process: processTemplate},
	{Name: "use-unicode", HasValue: false, Process: processUnicode},
}

// Universal function to process flags that toggle a boolean value
func processCheckFlag(target *bool) func(string) error {
	return func(_ string) error {
		*target = true
		return nil
	}
}

// Sets the Help configuration flag
func processHelp(value string) error {
	if Config.Colored || Config.Fullscreen || Config.Footprints || Config.EdgesPortal || Config.Verbose || Config.Delay != 0 || Config.Random != "" || Config.File != "" {
		return fmt.Errorf("--help cannot be used with other flags")
	}

	Config.Help = true
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

// Sets the Colored configuration flag
func processColored(value string) error {
	if Config.UseUnicode {
		return fmt.Errorf("--colored wasn't intended to be used with unicode characters")
	}

	Config.Colored = true
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

// Sets a template from the templates directory for map generation.
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

func processUnicode(value string) error {
	if Config.Colored {
		return fmt.Errorf("--colored wasn't intended to be used with unicode characters")
	}

	Config.UseUnicode = true
	return nil
}
