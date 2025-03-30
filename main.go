package main

import (
	"crunch03/utils"
	"fmt"
	"os"
)

func main() {
	if err := utils.ParseFlags(); err != nil {
		fmt.Println("Error in ParseFlags:", err)
		os.Exit(1)
	}

	if utils.Config.Help {
		utils.PrintHelp()
		return
	}

	if utils.Config.Delay == 0 {
		utils.Config.Delay = 2500
	}

	if err := utils.Input(); err != nil {
		fmt.Println("Error in Input():", err)
		os.Exit(2)
	}

	utils.RunGame()
}
