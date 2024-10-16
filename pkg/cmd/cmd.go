package cmd

import (
	"os"
	"strings"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func CommandLine(db *database.History) {

	db.PrintCLI()

	for {
		cmd := util.Input()

		itemParts := strings.Fields(cmd)

		if len(itemParts) != 2 {
			executeCommand(cmd, db)
		}

		db.Save(itemParts)

		CommandLine(db)
	}
}

func executeCommand(cmd string, db *database.History) {
	switch cmd {
	case "history", "h":
		db.PrintHistory()
		CommandLine(db)
	case "day", "d":
		db.PrintDays()
		CommandLine(db)
	case "stats", "s":
		db.PrintStatistics()
		CommandLine(db)
	case "backup":
		db.Backup()
		CommandLine(db)
	case "quit", "q":
		util.ClearScreen()
		os.Exit(0)
	default:
		CommandLine(db)
	}
}
