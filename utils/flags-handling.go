package utils

import (
	"fmt"
	"os"
	"strconv"
)

var PassedFlag map[string]bool = map[string]bool{
	"help":    false,
	"verbose": false,
	"random":  false,
}

var delay int = 2500

// isFlag: returns true if the argument is a flag
func isFlag(s string) bool {
	if len(s) < 3 {
		return false
	}

	return s[:2] == "--"
}

// hasValue: returns true if the flag has a value
func hasValue(flag string) bool {
	for _, char := range flag {
		if char == '=' {
			return true
		}
	}

	return false
}

// getFlag: returns the name or value of a flag
func getFlag(flag, target string) string { // target should be either "name" or "value"
	end := 0

	for i, char := range flag {
		if char == '=' {
			end = i
			break
		}
	}

	if target == "name" {
		return flag[2:end]
	} else if target == "value" {
		return flag[end+1:]
	} else {
		return ""
	}
}

// CheckFlags: checks the flags passed as arguments and sets the corresponding variables
func CheckFlags() {
	args := os.Args[1:] // discard the filename argument

	for _, arg := range args {
		if isFlag(arg) {
			if hasValue(arg) {
				if getFlag(arg, "name") == "delay-ms" {

					newDelay, err := strconv.Atoi(getFlag(arg, "value"))
					if err != nil {
						fmt.Println("Error: invalid value for delay")
					}

					delay = newDelay
				} else {
					fmt.Println("Error: unknown flag.")
				}
			} else {

				_, ok := PassedFlag[arg[2:]]

				if !ok {
					fmt.Println("Error: unknown flag.")
				}

				PassedFlag[arg[2:]] = true
			}
		} else {
			fmt.Println("Error: invalid argument.")
		}
	}
}
