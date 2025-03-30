package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds the configuration values set by flags.
var Config struct {
	Help    bool
	Verbose bool
	Delay   int
	Random  string
}

// Flag represents a command-line flag.
type Flag struct {
	Name     string
	HasValue bool
	Process  func(value string) error
}

// processHelp sets the Help flag.
func processHelp(value string) error {
	Config.Help = true
	return nil
}

// processVerbose sets the Verbose flag.
func processVerbose(value string) error {
	Config.Verbose = true
	return nil
}

// processDelay sets the Delay value.
func processDelay(value string) error {
	d, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("invalid delay value: %s", value)
	}
	Config.Delay = d
	return nil
}

// processRandom sets the Random flag value.
func processRandom(value string) error {
	// You might add validation for the WxH format here.
	Config.Random = value
	return nil
}

// flags defines the list of supported flags.
var flags = []Flag{
	{Name: "help", HasValue: false, Process: processHelp},
	{Name: "verbose", HasValue: false, Process: processVerbose},
	{Name: "delay-ms", HasValue: true, Process: processDelay},
	{Name: "random", HasValue: true, Process: processRandom},
}

// ParseFlags processes command-line arguments.
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

// extractFlag splits the flag into name and value if applicable.
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

// findFlag searches for a flag definition by name.
func findFlag(flagName string) (Flag, bool) {
	for _, f := range flags {
		if f.Name == flagName {
			return f, true
		}
	}
	return Flag{}, false
}

// processFlag applies the found flag's processing function.
func processFlag(flag Flag, value string) error {
	if flag.HasValue && value == "" {
		return fmt.Errorf("flag '--%s' requires a value", flag.Name)
	}
	if !flag.HasValue && value != "" {
		return fmt.Errorf("flag '--%s' does not accept a value", flag.Name)
	}
	return flag.Process(value)
}
