package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/dir"
	"github.com/VkHyperNova/VK-FINANCE/pkg/global"
	"github.com/VkHyperNova/VK-FINANCE/pkg/print"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func CMD() {

	print.ClearScreen()

	dir.ValidateRequiredFiles()

	database.GetFinances()

	print.PrintSeparatorSingleDash()
	print.PrintGray("============== VK FINANCE v1 ===============\n")
	print.PrintSeparatorSingleDash()

	database.CountIncomeAndExpenses()

	print.PrintStats()
	print.PrintSeparatorDoubleDash()
	print.PrintCommands()
	

	var user_input string
	print.PrintGray("\n\n=> ")
	fmt.Scanln(&user_input)

	for {
		switch user_input {
		case "add":
			AddIncome()
			CMD()
		case "spend":
			AddExpenses()
			CMD()
		case "history":
			PrintHistory()
			CMD()
		case "q":
			print.ClearScreen()
			os.Exit(0)
		default:
			CMD()
		}
	}
}

func AddIncome() {

	sum := util.UserInputFloat64("Sum: ")
	comment := util.UserInputString("Comment: ")

	global.LastAdd += sum

	database.SaveHistory(sum, comment)
}

func AddExpenses() {

	sum := util.UserInputFloat64("Sum: ")
	comment := util.UserInputString("Comment: ")

	global.LastExp += sum

	database.SaveHistory(-1*sum, comment)
}

func PrintHistory() {

	byteArray := dir.ReadFile("./history.json")
	historyJson := database.GetHistoryJson(byteArray)

	print.PrintCyan("History: \n\n")

	for _, value := range historyJson {

		val, err := json.Marshal(value.VALUE)
		print.HandleError(err)

		if value.VALUE < 0 {
			print.PrintRed(" ")
			print.PrintRed(value.DATE)
			print.PrintRed(" ")
			print.PrintRed(value.TIME)
			print.PrintRed(" ")
			print.PrintRed(value.COMMENT)
			print.PrintRed(" ==> ")
			print.PrintRed(string(val) + "\n")
		} else {
			print.PrintGreen(" ")
			print.PrintGreen(value.DATE)
			print.PrintGreen(" ")
			print.PrintGreen(value.TIME)
			print.PrintGreen(" ")
			print.PrintGreen(value.COMMENT)
			print.PrintGreen(" ==> ")
			print.PrintGreen(string(val) + "\n")
		}

	}

	fmt.Scanln() // Press enter to continue
}
