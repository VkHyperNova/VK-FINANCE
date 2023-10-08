package cmd

import (
	"encoding/json"
	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
	"time"
)

func DisplayAllVariables() {
	// Display separator line
	displaySeparatorSingleDash()

	// Display program title and version
	util.PrintGray("============== VK FINANCE v1 ===============\n")

	// Display separator line
	displaySeparatorSingleDash()

	// Display current month history
	displayCurrentMonthHistory()

	// Display net worth
	displayNetWorth()

	// Display income
	displayIncome()

	// Display all expenses
	displayExpences()

	// Calculate estimated daily spending amount
	displayEstimatedDaylySpendingAmount() // split this function!!!

	// Calculate estimated weekly spending amount
	displayEstimatedWeeklySpendingAmount()

	// Display saving amount
	displaySavingAmount()

	// Display balance
	displayBalance()

	// Calculate money left
	displayMoneyLeft()

	displaySeparatorDoubleDash()
}

func displayCurrentMonthHistory() {
	// Get current date and time
	now := time.Now()
	// Get current month
	CurrentMonth := now.Month()

	// Read history.json file and convert it to byte array
	byteArray := util.ReadFile("./history.json")
	// Convert byte array to historyJson
	historyJson := database.GetHistoryJson(byteArray)

	// Clear the screen
	util.ClearScreen()

	// Print cyan color text
	util.PrintCyan("History: \n\n")

	// Loop through historyJsonArray
	for _, value := range historyJson {
		// Define date layout format
		layout := "02-01-2006"

		// Parse date string to time.Time object
		t, err := time.Parse(layout, value.DATE)
		// Handle error if any
		util.HandleError(err)

		// Check if the month of the current date is equal to the current month
		if t.Month() == CurrentMonth {
			// Print the value

			val, err := json.Marshal(value.VALUE)
			util.HandleError(err)

			if value.VALUE < 0 {
				util.PrintRed(" ")
				util.PrintRed(value.DATE)
				util.PrintRed(" ")
				util.PrintRed(value.TIME)
				util.PrintRed(" ")
				util.PrintRed(value.COMMENT)
				util.PrintRed(" ==> ")
				util.PrintRed(string(val) + "\n")
			} else {
				util.PrintGreen(" ")
				util.PrintGreen(value.DATE)
				util.PrintGreen(" ")
				util.PrintGreen(value.TIME)
				util.PrintGreen(" ")
				util.PrintGreen(value.COMMENT)
				util.PrintGreen(" ==> ")
				util.PrintGreen(string(val) + "\n")
			}

		}
	}
}

func displayNetWorth() {
	util.PrintCyan("\nNET WORTH: ")
	util.PrintGreen(util.FloatToString(util.NET_WORTH) + " EUR\n\n")
}

func displayIncome() {
	util.PrintCyan("INCOME: ")
	util.PrintGreen("+" + util.FloatToString(util.INCOME) + " EUR")

	if util.LastAdd != 0 {
		util.PrintCyan(" | ")
		util.PrintYellow("+" + util.FloatToString(util.LastAdd) + " EUR")
	}
	util.PrintGray("\n")

}

func displayExpences() {
	util.PrintCyan("EXPENCES: ")
	util.PrintRed(util.FloatToString(util.EXPENSES) + " EUR")

	if util.LastExp != 0 {
		util.PrintCyan(" | ")
		util.PrintYellow("+" + util.FloatToString(util.LastExp) + " EUR")
	}
	util.PrintGray("\n\n")
}

func displayEstimatedDaylySpendingAmount() {
	// Print estimated daily savings budget
	util.PrintCyan("ESTIMATED DAY: ")
	util.PrintGreen(util.FloatToString(util.DayBudget) + " EUR")
	util.PrintCyan(" | ")
	// Print estimated daily spendable amount
	util.PrintRed(util.FloatToString(util.DayBudgetSpent) + " EUR\n")
}

func displayEstimatedWeeklySpendingAmount() {
	// Print estimated weekly spending amount
	util.PrintCyan("ESTIMATED WEEK: ")
	util.PrintGreen(util.FloatToString(util.WeekBudget) + " EUR")
	util.PrintCyan(" | ")
	util.PrintRed(util.FloatToString(util.WeekBudgetSpent) + " EUR\n")
}

func displaySavingAmount() {
	// Print cyan text
	util.PrintCyan("SAVING (25%): ")
	// Convert float to string with two decimal points
	savingAmount := util.FloatToString(util.SAVING)
	// Print green text
	util.PrintGreen(savingAmount + " EUR\n\n")
}

func displayBalance() {
	util.PrintCyan("BALANCE: ")
	util.PrintYellow(util.FloatToString(util.BALANCE) + " EUR\n")
}

func displayMoneyLeft() {
	// Print the text "MONEY: " in cyan color
	util.PrintCyan("MONEY: ")

	// Check if the money left is less than 0
	if util.Budget < 0 {
		// Print the money left in red color with two decimal points
		util.PrintRed(util.FloatToString(util.Budget) + " EUR\n\n")
	} else {
		// Print the money left in green color with two decimal points
		util.PrintGreen(util.FloatToString(util.Budget) + " EUR\n\n")
	}
}

func AddBrackets(name string) {
	util.PrintCyan("[")
	util.PrintYellow(name)
	util.PrintCyan("] ")
}

// Display separator with single dash
func displaySeparatorSingleDash() {
	util.PrintGray("============================================\n")
}

// Display separator with double dashes
func displaySeparatorDoubleDash() {
	util.PrintGray("--------------------------------------------\n")
}
