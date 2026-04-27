package cmd

import (
	"fmt"

	"github.com/VkHyperNova/VK-FINANCE/pkg/config"
	"github.com/VkHyperNova/VK-FINANCE/pkg/db"
)

func Start(f *db.Finance) {

	for {

		f.PrintDashboard()

		var command string
		var sum float64

		fmt.Scanln(&command, &sum)

		switch command {
		case "history", "h":
			f.PrintHistory()
		case "undo":
			f.Undo()
		case "backup":
			err := f.Backup()
			if err != nil {
				fmt.Println(err)
			}
		case "import", "i":
			if err := f.ImportDB(config.BackupFile); err != nil {
				fmt.Println(err)
			}
		case "quit", "q":
			return
		default:
			if command == "" || sum == 0.0 {
				continue
			}
			config.LastAddedItemName = command
			config.LastAddedItemSum = sum
			f.Add(command, sum)
		}
	}
}
