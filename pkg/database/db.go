package database

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

// ValidateDatabase checks if the required files exist
func ValidateRequiredFiles() {
	// Check if the finance.json file exists
	if !DoesDirectoryExist("./finance.json") {
		// If not, get the user's net worth input
		util.NET_WORTH = util.UserInputFloat64("NET_WORTH: ")
		// Save the user's input to the database
		Save(0, "Fresh Start")
	}

	// Check if the history.json file exists
	if !DoesDirectoryExist("./history.json") {
		// If not, create an empty file
		WriteDataToFile("./history.json", []byte("[]"))
	}
}

// FetchFinanceDataFromFile reads the finance data from a file and stores it in variables
func FetchFinanceDataFromFile() {
	// Read the finance data from a file
	byteArray := ReadFile("./finance.json")

	// Convert the byte array to a FinanceJsonObject
	financeJsonObject := util.GetFinanceJson(byteArray)

	// Store the values from the FinanceJsonObject in variables
	util.NET_WORTH = financeJsonObject.NET_WORTH
	util.BALANCE = financeJsonObject.BALANCE
	util.EXPENSES = financeJsonObject.EXPENSES
	util.INCOME = financeJsonObject.INCOME

	// Calculate the perfect save amount
	util.SAVING = util.INCOME * 0.25
}
