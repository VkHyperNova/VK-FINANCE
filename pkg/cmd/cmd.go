package cmd

import (
	"fmt"
	"os"

	"github.com/VkHyperNova/VK-FINANCE/pkg/config"
	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func CommandLine(db *database.Finance) {

	db.PrintCLI()

	for {
		var name string = ""
		var sum float64 = 0.0

		fmt.Print("\n=> ")

		fmt.Scanln(&name, &sum)

		if sum == 0.0 {
			executeCommand(name, db)
		}

		config.LastAddedItemName = name

		db.Insert(name, sum)

		CommandLine(db)
	}
}

func executeCommand(cmd string, db *database.Finance) {
	switch cmd {
	case "history", "h":
		db.PrintHistory()
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
