package utils

import (
	"fmt"
	"os"
	"strings"
)

// Processes command-line arguments into configuration flags
func ParseFlags() error {
	args := os.Args[1:]

	for _, arg := range args {
		flagName, flagValue, valid := extractFlag(arg)
		if !valid {
			return fmt.Errorf("Error: invalid argument '%s'\n", arg)
		}

		flag, found := findFlag(flagName)
		if !found {
			return fmt.Errorf("Error: unknown flag '--%s'\n", flagName)
		}

		if err := processFlag(flag, flagValue); err != nil {
			return fmt.Errorf("Error processing flag '--%s': %s\n", flagName, err)
		}
	}
	return nil
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
