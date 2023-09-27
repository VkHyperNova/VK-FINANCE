package database

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func Save(Action string, Value float64) {
	// Construct finance financeJsonObject as a JSON object
	financeJsonObject := util.Finance()

	// Convert finance data to a byte array
	financeByteArray := util.InterfaceToByteArray(financeJsonObject)

	// Write finance data to a JSON file
	WriteDataToFile("./finance.json", financeByteArray)

	// Read the history file content
	historyByteArray := ReadFile("./history.json")

	// Convert the file content to history data
	historyJsonArray := util.GetHistoryJson(historyByteArray)

	// Construct a new history JSON object
	historyJsonArrayObject := util.History(Action, Value)

	// Append the new data to the history data
	historyJsonArray = append(historyJsonArray, historyJsonArrayObject)

	// Convert the history data to a byte array
	historyByteArrayUpdated := util.InterfaceToByteArray(historyJsonArray)

	// Write the data to the history file
	WriteDataToFile("./history.json", historyByteArrayUpdated)
}

