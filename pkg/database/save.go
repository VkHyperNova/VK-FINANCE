package database

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func Save(Value float64, Comment string) {

	financeJsonObject := SetFinanceJson() // Construct finance financeJsonObject as a JSON object

	financeByteArray := util.InterfaceToByte(financeJsonObject) // Convert finance data to a byte array

	util.WriteDataToFile("./finance.json", financeByteArray) // Write finance data to a JSON file

	historyByteArray := util.ReadFile("./history.json") // Read the history file content

	historyJsonArray := GetHistoryJson(historyByteArray) // Convert the file content to history data

	historyJsonArrayObject := SetHistoryJson(Value, Comment) // Construct a new history JSON object

	historyJsonArray = append(historyJsonArray, historyJsonArrayObject) // Append the new data to the history data

	historyByteArrayUpdated := util.InterfaceToByte(historyJsonArray) // Convert the history data to a byte array

	util.WriteDataToFile("./history.json", historyByteArrayUpdated) // Write the data to the history file
}

// ValidateDatabase checks if the required files exist
func ValidateRequiredFiles() {
	// Check if the finance.json file exists
	if !util.DoesDirectoryExist("./finance.json") {
		// If not, get the user's net worth input
		util.NET_WORTH = util.UserInputFloat64("NET_WORTH: ")
		// Save the user's input to the database
		Save(0, "Fresh Start")
	}

	// Check if the history.json file exists
	if !util.DoesDirectoryExist("./history.json") {
		// If not, create an empty file
		util.WriteDataToFile("./history.json", []byte("[]"))
	}
}
