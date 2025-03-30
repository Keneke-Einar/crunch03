package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Generates a random game map with specified dimensions
func GenerateRandomMap(dimensions string) error {
	parts := strings.Split(dimensions, "x")
	if len(parts) != 2 {
		return fmt.Errorf("Error: invalid format for --random flag. Use --random=WxH")
	}

	width, errW := strconv.Atoi(parts[0])
	height, errH := strconv.Atoi(parts[1])

	if errW != nil || errH != nil || width <= 0 || height <= 0 {
		return fmt.Errorf("Error: invalid dimensions for --random flag. Width and height must be positive integers.")
	}

	if w < 2 || h < 2 {
		return fmt.Errorf("invalid grid size: %dx%d. Minimum size is 2x2", w, h)
	}

	if Config.Fullscreen {
		termWidth, termHeight = GetTerminalSize()

		effectiveHeight := termHeight
		if Config.Verbose {
			effectiveHeight -= 5
		}

		if height < effectiveHeight {
			height = effectiveHeight
		}
		if width < termWidth {
			width = termWidth
		}
	}

	w, h = width, height
	gameMap = make([][]rune, h)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < h; i++ {
		row := make([]rune, w)
		for j := 0; j < w; j++ {
			if rand.Intn(2) == 0 {
				row[j] = '#'
			} else {
				row[j] = '.'
			}
		}
		gameMap[i] = row
	}

	return nil
}
