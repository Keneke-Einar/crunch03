package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateRandomMap(dimensions string) error {
	width, height, err := parseDimensions(dimensions)
	if err != nil {
		return err
	}

	applyFullscreenAdjustments(&width, &height)
	initializeGameMap(width, height, true)

	return nil
}

// Parses dimensions from a "WxH" formatted string
func parseDimensions(dimensions string) (int, int, error) {
	parts := strings.Split(dimensions, "x")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("error: invalid format for --random flag. Use --random=WxH")
	}

	width, errW := strconv.Atoi(parts[0])
	height, errH := strconv.Atoi(parts[1])
	if errW != nil || errH != nil || width <= 0 || height <= 0 {
		return 0, 0, fmt.Errorf("error: invalid dimensions for --random flag. Width and height must be positive integers")
	}

	if width < 3 || height < 3 {
		return 0, 0, fmt.Errorf("error: invalid grid size. Minimum size is 3x3")
	}

	return width, height, nil
}

// Adjusts dimensions for fullscreen mode if enabled
func applyFullscreenAdjustments(width, height *int) {
	if Config.Fullscreen {
		termWidth, termHeight := GetTerminalSize()
		effectiveHeight := termHeight
		if Config.Verbose {
			effectiveHeight -= 5
		}

		if *height < effectiveHeight {
			*height = effectiveHeight
		}
		if *width < termWidth {
			*width = termWidth
		}
	}
}

// Initializes the game map with random '#' and '.' characters
func initializeGameMap(width, height int, randomize bool) {
	w, h = width, height
	gameMap = make([][]rune, h)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < h; i++ {
		gameMap[i] = make([]rune, w)
		for j := 0; j < w; j++ {
			if randomize && rand.Intn(2) == 0 {
				gameMap[i][j] = '#'
			} else {
				gameMap[i][j] = '.'
			}
		}
	}
}
