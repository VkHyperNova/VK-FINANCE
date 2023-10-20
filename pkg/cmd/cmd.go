package cmd

import (
	"fmt"
	"os"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/dir"
	"github.com/VkHyperNova/VK-FINANCE/pkg/print"
)

func CMD() {

	print.ClearScreen()

	dir.ValidateRequiredFiles()
	database.GetFinanceJson()

	database.DayBudget()
	database.WeekBudget()
	database.Budget()

	print.PrintSeparatorSingleDash()
	print.PrintGray("============== VK FINANCE v1 ===============\n")
	print.PrintSeparatorSingleDash()

	database.PrintHistory()
	database.CountAndPrintHistoryItems()

	print.PrintNetWorth()
	print.PrintIncome()
	print.PrintExpences()
	print.PrintEstimatedDaylySpendingAmount()
	print.PrintEstimatedWeeklySpendingAmount()
	print.PrintSavingAmount()
	print.PrintBalance()
	print.PrintMoneyLeft()
	print.PrintSeparatorDoubleDash()

	print.PrintCyan("Program Options: \n\n")

	print.PrintWithBrackets("add")
	print.PrintWithBrackets("spend")
	print.PrintWithBrackets("networth")
	print.PrintWithBrackets("reset")
	print.PrintWithBrackets("q")

	var user_input string
	print.PrintGray("\n\n=> ")
	fmt.Scanln(&user_input)

	for {
		switch user_input {
		case "add":
			database.CalculateIncome()
			CMD()
		case "spend":
			database.CalculateExpenses()
			CMD()
		case "networth":
			database.AddNetWorth()
			CMD()
		case "reset":
			database.ResetVariables()
			database.Save(0, "Reset Income and Expenses")
			CMD()
		case "q":
			print.ClearScreen()
			os.Exit(0)
		default:
			CMD()
		}
	}
}
