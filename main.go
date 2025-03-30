package main

import "crunch03/utils"

func main() {
	utils.ParseFlags()

	if utils.Config.Help {
		utils.PrintHelp()
		return
	}

	if utils.Config.Delay == 0 {
		utils.Config.Delay = 2500
	}

	utils.Input()
	utils.RunGame()
}
