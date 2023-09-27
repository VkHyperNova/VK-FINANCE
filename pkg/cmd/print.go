package cmd

import (
	"encoding/json"
	"time"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
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
	displayAllExpences()

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
	byteArray := database.ReadFile("./history.json")
	// Convert byte array to historyJson
	historyJson := util.GetHistoryJson(byteArray)

	// Clear the screen
	util.ClearScreen()

	// Print cyan color text
	util.PrintCyan("History: \n")

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

			val, err := json.Marshal(value)
			util.HandleError(err)

			util.PrintGray(string(val) + "\n")
		}
	}
}

func displayNetWorth() {
	util.PrintCyan("NET WORTH: ")
	util.PrintGreen(util.Float64ToStringWithTwoDecimalPoints(util.NET_WORTH) + " EUR\n\n")
}

func displayIncome() {
	util.PrintCyan("INCOME: ")
	util.PrintGreen("+" + util.Float64ToStringWithTwoDecimalPoints(util.INCOME) + " EUR\n")
}

func displayAllExpences() {
	util.PrintCyan("EXPENCES: ")
	util.PrintRed(util.Float64ToStringWithTwoDecimalPoints(util.EXPENSES) + " EUR\n\n")

	util.PrintCyan("-> Bills: ")
	util.PrintRed(util.Float64ToStringWithTwoDecimalPoints(util.BILLS) + " EUR\n")
	util.PrintCyan("-> Gas: ")
	util.PrintRed(util.Float64ToStringWithTwoDecimalPoints(util.GAS) + " EUR\n")
	util.PrintCyan("-> Food: ")
	util.PrintRed(util.Float64ToStringWithTwoDecimalPoints(util.FOOD) + " EUR\n")
	util.PrintCyan("-> Other: ")
	util.PrintRed(util.Float64ToStringWithTwoDecimalPoints(util.OTHER) + " EUR\n\n")
}

func displayEstimatedDaylySpendingAmount() {
	// Print estimated daily savings budget
	util.PrintCyan("ESTIMATED DAY: ")
	util.PrintGreen(util.Float64ToStringWithTwoDecimalPoints(util.MaxSavingsBudgetDay) + " EUR")
	util.PrintCyan(" | ")
	// Print estimated daily spendable amount
	util.PrintRed(util.Float64ToStringWithTwoDecimalPoints(util.MaxSpendableAmountDay) + " EUR\n")
}

func displayEstimatedWeeklySpendingAmount() {
	// Print estimated weekly spending amount
	util.PrintCyan("ESTIMATED WEEK: ")
	util.PrintGreen(util.Float64ToStringWithTwoDecimalPoints(util.MaxSavingsBudgetWeek) + " EUR")
	util.PrintCyan(" | ")
	util.PrintRed(util.Float64ToStringWithTwoDecimalPoints(util.MaxSpendableAmountWeek) + " EUR\n")
}

func displaySavingAmount() {
	// Print cyan text
	util.PrintCyan("SAVING (25%): ")
	// Convert float to string with two decimal points
	savingAmount := util.Float64ToStringWithTwoDecimalPoints(util.PERFECT_SAVE)
	// Print green text
	util.PrintGreen(savingAmount + " EUR\n\n")
}

func displayBalance() {
	util.PrintCyan("BALANCE: ")
	util.PrintYellow(util.Float64ToStringWithTwoDecimalPoints(util.BALANCE) + " EUR\n")
}

func displayMoneyLeft() {
	// Print the text "MONEY: " in cyan color
	util.PrintCyan("MONEY: ")

	// Check if the money left is less than 0
	if util.MONEY < 0 {
		// Print the money left in red color with two decimal points
		util.PrintRed(util.Float64ToStringWithTwoDecimalPoints(util.MONEY) + " EUR\n\n")
	} else {
		// Print the money left in green color with two decimal points
		util.PrintGreen(util.Float64ToStringWithTwoDecimalPoints(util.MONEY) + " EUR\n\n")
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
