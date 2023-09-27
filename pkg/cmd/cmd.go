package cmd

import (
	"fmt"
	"os"
	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)
func DisplayAndHandleCommandLineCommands() {
	// Clear the screen
	util.ClearScreen()

	// Validate the database
	database.ValidateRequiredFiles()

	// Fetch the finance data
	database.FetchFinanceDataFromFile()

	// Calculate
	Calculate()

	// Print all the data
	DisplayAllVariables()

	// Print the program options
	util.PrintCyan("Program Options: \n\n")

	// Display the command names
	AddBrackets("add")
	AddBrackets("bills")
	AddBrackets("gas")
	AddBrackets("food")
	AddBrackets("other")
	AddBrackets("grow")
	AddBrackets("reset")
	AddBrackets("q")

	// Get the user input
	var user_input string

	util.PrintGray("\n\n=> ")
	fmt.Scanln(&user_input)

	// Handle the user input
	for {
		switch user_input {
		case "add":
			ReCalculateCashFlow("Add")
		case "bills":
			ReCalculateCashFlow("Bills")
		case "gas":
			ReCalculateCashFlow("Gas")
		case "food":
			ReCalculateCashFlow("Food")
		case "other":
			ReCalculateCashFlow("Other")
		case "grow":
			NetWorth()
		case "reset":
			ResetVariables()
			database.Save("Reset", 0)
			DisplayAndHandleCommandLineCommands()
			
		case "q":
			util.ClearScreen()
			os.Exit(0)
		default:
			DisplayAndHandleCommandLineCommands()
		}
	}
}

func calculateIncome(amount float64) {
	// Add amount to income
	util.INCOME = util.INCOME + amount
	// Add amount to balance
	util.BALANCE = util.BALANCE + amount
}

func calculateExpenses(amount float64) {
	// subtract amount from balance
	util.BALANCE = util.BALANCE - amount
	// subtract amount from expenses
	util.EXPENSES = util.EXPENSES - amount
}

// NetWorth function increases the net worth by the balance amount
func NetWorth() {
	// increase net worth by balance amount
	util.NET_WORTH = util.NET_WORTH + util.BALANCE
	// set saved amount to balance amount
	SAVED_AMOUNT := util.BALANCE
	// reset balance to 0
	util.BALANCE = 0
	// reset other variables
	ResetVariables()
	// save data to file
	database.Save("Grow", SAVED_AMOUNT)
	// display and handle command line commands
	DisplayAndHandleCommandLineCommands()



}

// Reset all variables to 0
func ResetVariables() {
	util.INCOME = 0
	util.EXPENSES = 0
	util.BILLS = 0
	util.GAS = 0
	util.FOOD = 0
	util.OTHER = 0
}

// ReCalculateCashFlow calculates cash flow transactions
func ReCalculateCashFlow(name string) {
	// getUserInput retrieves user input for a specific transaction
	sum_of_money := util.GetUserInput(name + ": ")

	// switch statement handles different transaction types
	switch name {
	case "Bills":
		// subtract transaction amount from BILLS
		util.BILLS -= sum_of_money
		// calculate expenses and balance
		calculateExpenses(sum_of_money)
	case "Gas":
		// subtract transaction amount from GAS
		util.GAS -= sum_of_money
		// calculate expenses and balance
		calculateExpenses(sum_of_money)
	case "Food":
		// subtract transaction amount from FOOD
		util.FOOD -= sum_of_money
		// calculate expenses and balance
		calculateExpenses(sum_of_money)
	case "Other":
		// subtract transaction amount from OTHER
		util.OTHER -= sum_of_money
		// calculate expenses and balance
		calculateExpenses(sum_of_money)
	case "Add":
		// calculate income and balance
		calculateIncome(sum_of_money)
	}

	// SaveData saves transaction data to a file
	database.Save(name, sum_of_money)

	DisplayAndHandleCommandLineCommands()
}
