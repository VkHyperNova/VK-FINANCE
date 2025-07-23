package cmd

import (
	"fmt"
	"os"

	"github.com/VkHyperNova/VK-FINANCE/pkg/config"
	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func CommandLine(db *database.History) {

	db.PrintCLI()

	for {
		var cmd string = ""
		var sum float64 = 0.0

		fmt.Print("\n=> ")

		fmt.Scanln(&cmd, &sum)

		if sum == 0.0 {
			executeCommand(cmd, db)
		}

		db.Insert(cmd, sum)

		CommandLine(db)
	}
}

func executeCommand(cmd string, db *database.History) {
	switch cmd {
	case "history", "h":
		db.PrintHistory()
		CommandLine(db)
	case "stats", "s":
		db.PrintStatistics()
		CommandLine(db)
	case "undo":
		db.Undo()
		CommandLine(db)
	case "backup":
		db.Backup()
		CommandLine(db)
	case "quit", "q":
		util.ClearScreen()
		os.Exit(0)
	default:
		fmt.Println(config.Red, "Command does not exist", config.Reset)
		util.PressAnyKey()
		CommandLine(db)
	}
}
