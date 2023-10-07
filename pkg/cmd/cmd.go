package cmd

import (
	"fmt"
	"os"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func DisplayAndHandleOptions() {

	util.ClearScreen() // Clear the screen

	database.ValidateRequiredFiles()    // Validate the database
	util.FetchFinanceDataFromFile() // Fetch the finance data

	DayBudget()
	WeekBudget()
	Budget()

	DisplayAllVariables() // Print all the data

	util.PrintCyan("Program Options: \n\n") // Print the program options

	// Display the command names
	AddBrackets("add")
	AddBrackets("spend")
	AddBrackets("grow")
	AddBrackets("reset")
	AddBrackets("q")

	var user_input string
	util.PrintGray("\n\n=> ")
	fmt.Scanln(&user_input) // Get the user input

	// Handle the user input
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
			DisplayAndHandleOptions()
		case "q":
			util.ClearScreen()
			os.Exit(0)
		default:
			DisplayAndHandleOptions()
		}
	}
}

func calculateIncome() {

	sum := util.UserInputFloat64("Sum: ")
	comment := util.UserInputString("Comment: ")

	util.LastAdd += sum
	
	util.INCOME = util.INCOME + sum
	util.BALANCE = util.BALANCE + sum

	database.Save(sum, comment)
	DisplayAndHandleOptions()
}

func calculateExpenses() {

	sum := util.UserInputFloat64("Sum: ")
	comment := util.UserInputString("Comment: ")

	util.LastExp += sum

	util.BALANCE = util.BALANCE - sum
	util.EXPENSES = util.EXPENSES - sum

	database.Save(-1 * sum, comment)
	DisplayAndHandleOptions()
}

func NetWorth() {
	util.NET_WORTH = util.NET_WORTH + util.BALANCE // Increase net worth by balance amount.
	SAVED_BALANCE := util.BALANCE                  // Save balance to a variable before setting it to 0.
	util.BALANCE = 0                               // Reset balance to 0.

	ResetVariables()                                         // Reset other variables.
	database.Save(SAVED_BALANCE, "Update Net Worth") // Save finance data and action history.
	DisplayAndHandleOptions()                                // Back to command line.
}

// Reset all variables to 0
func ResetVariables() {
	util.BALANCE = 0
	util.INCOME = 0
	util.EXPENSES = 0
}
