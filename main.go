package main

import "crunch03/utils"

func main() {
	utils.CheckFlags()

	if utils.PassedFlag["help"] {
		utils.PrintHelp()
		return
	}

	utils.Input()
	utils.RunGame()
}
