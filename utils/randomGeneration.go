package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// GenerateRandomMap: generates a random game map with given dimensions
func GenerateRandomMap(dimensions string) {
	parts := strings.Split(dimensions, "x")
	if len(parts) != 2 {
		fmt.Println("Error: invalid format for --random flag. Use --random=WxH")
		return
	}

	width, errW := strconv.Atoi(parts[0])
	height, errH := strconv.Atoi(parts[1])

	if errW != nil || errH != nil || width <= 0 || height <= 0 {
		fmt.Println("Error: invalid dimensions for --random flag. Width and height must be positive integers.")
		return
	}

	// If fullscreen is enabled, adjust dimensions to terminal size
	if Config.Fullscreen {
		termWidth, termHeight = GetTerminalSize()

		// Leave room for status info if verbose is enabled
		effectiveHeight := termHeight
		if Config.Verbose {
			effectiveHeight -= 5 // Reserve space for verbose output
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
				row[j] = '#' // live cell
			} else {
				row[j] = '.' // dead cell
			}
		}
		gameMap[i] = row
	}
}
