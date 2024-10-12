package cmd

import (
	"os"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func CommandLine(history *database.History) {

	history.PrintCLI()

	cmd := util.Input("=> ")

	// Add
	if history.Split(cmd) {
		history.Save()
		history.PrintItemSummary()
		util.PressAnyKey()
		CommandLine(history)
	}

	for {
		switch cmd {
		case "history", "h":
			history.PrintHistory()
			CommandLine(history)
		case "day", "d":
			history.PrintDaySummary()
			util.PressAnyKey()
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
