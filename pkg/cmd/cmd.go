package cmd

import (
	"os"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

/* Main Functions */

func CMD() {
	util.ClearScreen()
	db := database.OpenDatabase()

	input := PrintCLI(db)

	for {
		switch input {
		case "add", "a":
			Add(db, true)
			CMD()
		case "spend", "s":
			Add(db, false)
			CMD()
		case "history", "h":
			PrintHistory(db)
			CMD()
		case "backup":
			Backup(db)
			CMD()
		case "q":
			util.ClearScreen()
			os.Exit(0)
		default:
			CMD()
		}
	}
}
