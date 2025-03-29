package utils

import (
	"fmt"
	"os"
	"strconv"
)

var passedFlag map[string]bool = map[string]bool{
	"help":    false,
	"verbose": false,
}

var (
	delay int = 2500
	// sourceFilename string = ""
	// randomGridSize [2]int // stores width and height in that order
)

// Returns true if the string is a flag (starts with "--").
func isFlag(s string) bool {
	if len(s) < 3 {
		return false
	}

	return s[:2] == "--"
}

// Returns true if the flag has an assigned value (provided the argument is a flag).
func hasValue(flag string) bool {
	for _, char := range flag {
		if char == '=' {
			return true
		}
	}

	return false
}

// Returns the name or the value (target) of the flag that has an assigned value.
func getFlag(flag, target string) string {
	// target should be either "name" or "value"
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

func CheckFlags() {
	args := os.Args[1:] // discard the filename argument

	for _, arg := range args {
		if isFlag(arg) {
			if hasValue(arg) {
				if getFlag(arg, "name") == "delay-ms" {
					newDelay, err := strconv.Atoi(getFlag(arg, "value"))

					if err != nil {
						fmt.Println("Error: invalid value for delay")
						return
					}

					delay = newDelay
				} else {
					fmt.Println("Error: unknown flag.")
					return
				}
			} else {
				_, ok := passedFlag[arg[2:]]

				if !ok {
					fmt.Println("Error: unknown flag.")
					return
				}

				passedFlag[arg[2:]] = true
			}
		} else {
			fmt.Println("Error: invalid argument.")
			return
		}
	}
}
