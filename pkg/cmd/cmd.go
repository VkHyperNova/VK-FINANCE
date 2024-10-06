package cmd

import (
	"os"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func CommandLine(history *database.History) {

	history.PrintCLI()

	cmd := util.Input("=> ")

	for {
		switch cmd {
		case "add", "a":
			userInput := util.Input("add: ")
			if history.SplitInput(userInput) {
				history.SaveToFile()
				history.Summary()
			}
			util.PressAnyKey()
			CommandLine(history)
		case "history", "h":
			history.PrintHistory()
			CommandLine(history)
		case "day", "d":
			history.PrintDailySpending()
			CommandLine(history)
		case "backup":
			history.Backup()
			CommandLine(history)
		case "quit", "q":
			util.ClearScreen()
			os.Exit(0)
		default:
			CommandLine(history)
		}
	}
}
