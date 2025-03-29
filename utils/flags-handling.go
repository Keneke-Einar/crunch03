package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var PassedFlag map[string]bool = map[string]bool{
	"help":    false,
	"verbose": false,
	"random":  false,
}

var (
	delay            int    = 2500
	randomDimensions string = ""
)

// CheckFlags: checks the flags passed as arguments and sets the corresponding variables
func CheckFlags() error {
	args := os.Args[1:] // discard the filename argument

	for _, arg := range args {
		if !isFlag(arg) {
			return fmt.Errorf("invalid argument: %s", arg)
		}

		flagName, err := getFlag(arg, "name")
		if err != nil {
			return err
		}

		if _, exists := PassedFlag[flagName]; !exists && flagName != "delay-ms" && flagName != "random" {
			return fmt.Errorf("unknown flag: --%s", flagName)
		}

		if hasValue(arg) {
			// need to rewrite other flags to use this format
			flagValue, err := getFlag(arg, "value")
			if err != nil {
				return err
			}

			switch flagName {
			case "delay-ms":
				newDelay, err := strconv.Atoi(flagValue)
				if err != nil {
					return fmt.Errorf("invalid value for --delay-ms: %s", flagValue)
				}
				delay = newDelay
			case "random":
				if err := validateRandomFlag(flagValue); err != nil {
					return err
				}
				randomDimensions = flagValue
				PassedFlag["random"] = true
			default:
				return fmt.Errorf("unknown flag: --%s", flagName)
			}
		} else {
			if _, exists := PassedFlag[flagName]; !exists {
				return fmt.Errorf("unknown flag: --%s", flagName)
			}
			PassedFlag[flagName] = true
		}
	}
	return nil
}

//=======================Utility functions=======================

// isFlag: returns true if the argument is a flag
func isFlag(s string) bool {
	return len(s) > 2 && s[:2] == "--"
}

// hasValue: returns true if the flag has a value
func hasValue(flag string) bool {
	return strings.Contains(flag, "=")
}

// getFlag: returns the name or value of a flag
func getFlag(flag, target string) (string, error) {
	parts := strings.SplitN(flag[2:], "=", 2)

	if len(parts) == 1 && target == "value" {
		return "", errors.New("flag --%s requires value" + parts[0])
	}

	switch target {
	case "name":
		return parts[0], nil
	case "value":
		return parts[1], nil
	default:
		return "", fmt.Errorf("invalid target: %s", target)
	}
}

//=======================Flag functions=======================

// validateRandomFlag: validates the dimensions for the random flag
func validateRandomFlag(dimensions string) error {
	parts := strings.Split(dimensions, "x")
	if len(parts) != 2 {
		return errors.New("invalid format for --random flag. Use --random=WxH")
	}

	width, err1 := strconv.Atoi(parts[0])
	height, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || width <= 0 || height <= 0 {
		return errors.New("invalid dimensions for --random. Width and height must be positive integers")
	}

	return nil
}
