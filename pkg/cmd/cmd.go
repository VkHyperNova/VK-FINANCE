package cmd

import (
	"fmt"
	"os"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
	"github.com/VkHyperNova/VK-FINANCE/pkg/print"
	"github.com/VkHyperNova/VK-FINANCE/pkg/global"
	"github.com/VkHyperNova/VK-FINANCE/pkg/dir"
)

func CMD() {

	print.ClearScreen()

	

	dir.ValidateRequiredFiles() 
	database.GetFinanceJson()        

	DayBudget()
	WeekBudget()
	Budget()

	print.PrintSeparatorSingleDash()
	print.PrintGray("============== VK FINANCE v1 ===============\n")
	print.PrintSeparatorSingleDash()

	database.PrintCurrentMonthHistory()
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
	print.PrintWithBrackets("grow")
	print.PrintWithBrackets("reset")
	print.PrintWithBrackets("q")

	var user_input string
	print.PrintGray("\n\n=> ")
	fmt.Scanln(&user_input) 

	for {
		switch user_input {
		case "add":
			calculateIncome()
		case "spend":
			calculateExpenses()
		case "grow":
			NetWorth()
		case "reset":
			ResetVariables()
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

func calculateIncome() {

	sum := util.UserInputFloat64("Sum: ")
	comment := util.UserInputString("Comment: ")

	global.LastAdd += sum

	global.INCOME = global.INCOME + sum
	global.BALANCE = global.BALANCE + sum

	database.Save(sum, comment)
	CMD()
}

func calculateExpenses() {

	sum := util.UserInputFloat64("Sum: ")
	comment := util.UserInputString("Comment: ")

	global.LastExp += sum

	global.BALANCE = global.BALANCE - sum
	global.EXPENSES = global.EXPENSES - sum

	database.Save(-1*sum, comment)
	CMD()
}

func NetWorth() {
	global.NET_WORTH = global.NET_WORTH + global.BALANCE 
	SAVED_BALANCE := global.BALANCE                  
	global.BALANCE = 0                              

	ResetVariables()                                 
	database.Save(SAVED_BALANCE, "Update Net Worth") 
	CMD()                                            
}
func ResetVariables() {
	global.BALANCE = 0
	global.INCOME = 0
	global.EXPENSES = 0
}
