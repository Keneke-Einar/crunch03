package main

import (
	"crunch03/utils"
	"fmt"
	"os"
)

func main() {
	utils.ParseFlags()

	if utils.Config.Help {
		utils.PrintHelp()
		return
	}

	if utils.Config.Delay == 0 {
		utils.Config.Delay = 2500
	}

	err := utils.Input()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	utils.RunGame()
}
