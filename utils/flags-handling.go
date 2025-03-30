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

// flags defines the list of supported flags.
var flags = []Flag{
	{
		Name:     "help",
		HasValue: false,
		Process: func(value string) error {
			Config.Help = true
			return nil
		},
	},
	{
		Name:     "verbose",
		HasValue: false,
		Process: func(value string) error {
			Config.Verbose = true
			return nil
		},
	},
	{
		Name:     "delay-ms",
		HasValue: true,
		Process: func(value string) error {
			d, err := strconv.Atoi(value)
			if err != nil {
				return fmt.Errorf("invalid delay value: %s", value)
			}
			Config.Delay = d
			return nil
		},
	},
	{
		Name:     "random",
		HasValue: true,
		Process: func(value string) error {
			Config.Random = value
			return nil
		},
	},
}

// ParseFlags checks the command-line arguments and processes any supported flags.
func ParseFlags() {
	args := os.Args[1:]
	for _, arg := range args {
		if !strings.HasPrefix(arg, "--") {
			fmt.Printf("Error: invalid argument '%s'\n", arg)
			continue
		}

		content := arg[2:]
		var flagName, flagValue string

		if strings.Contains(content, "=") {
			parts := strings.SplitN(content, "=", 2)
			flagName = parts[0]
			flagValue = parts[1]
		} else {
			flagName = content
		}

		found := false
		for _, f := range flags {
			if f.Name == flagName {
				found = true
				if f.HasValue && flagValue == "" {
					fmt.Printf("Error: flag '--%s' requires a value\n", flagName)
					break
				}
				if !f.HasValue && flagValue != "" {
					fmt.Printf("Error: flag '--%s' does not accept a value\n", flagName)
					break
				}
				if err := f.Process(flagValue); err != nil {
					fmt.Printf("Error processing flag '--%s': %s\n", flagName, err.Error())
				}
				break
			}
		}
		if !found {
			fmt.Printf("Error: unknown flag '--%s'\n", flagName)
		}
	}
}
