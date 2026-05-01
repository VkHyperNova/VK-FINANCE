package cmd

import (
	"fmt"

	"github.com/VkHyperNova/VK-FINANCE/pkg/db"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
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
		case "restart":
			err := f.Restart()
			if err != nil {
				fmt.Println(err)
			}
			util.PressAnyKey()
		case "import", "i":
			if err := f.Import(); err != nil {
				fmt.Println(err)
			}
			util.PressAnyKey()
		case "export", "e":
			if err := f.Export(); err != nil {
				fmt.Println(err)
			}
			util.PressAnyKey()
		case "unmount":
			if err := util.UnmountDrive(); err != nil {
				fmt.Println(err)
			}
			util.PressAnyKey()
		case "quit", "q":
			return
		default:
			if command == "" || sum == 0.0 {
				fmt.Println("Add Expence: food -10")
				fmt.Println("Add Income: wolt 10")
				util.PressAnyKey()
				continue
			}
			f.Add(command, sum)
		}
	}
}
