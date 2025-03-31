package main

import (
	"crunch03/utils"
	"fmt"
	"os"
)

func main() {
	nArgs, err := utils.ParseFlags()
	if err != nil {
		fmt.Println("Error in ParseFlags:", err)
		os.Exit(1)
	}

	if utils.Config.Help && nArgs == 1 {
		utils.PrintHelp()
		return
	} else if utils.Config.Help && nArgs > 1 {
		fmt.Println("Error: --help cannot be used with other flags")
		os.Exit(3)
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
