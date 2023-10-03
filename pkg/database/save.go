package database

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func Save(Value float64, Comment string) {
	// Construct finance financeJsonObject as a JSON object
	financeJsonObject := util.Finance()

	// Convert finance data to a byte array
	financeByteArray := util.InterfaceToByte(financeJsonObject)

	// Write finance data to a JSON file
	WriteDataToFile("./finance.json", financeByteArray)

	// Read the history file content
	historyByteArray := ReadFile("./history.json")

	// Convert the file content to history data
	historyJsonArray := util.GetHistoryJson(historyByteArray)

	// Construct a new history JSON object
	historyJsonArrayObject := util.History(Value, Comment)

	// Append the new data to the history data
	historyJsonArray = append(historyJsonArray, historyJsonArrayObject)

	// Convert the history data to a byte array
	historyByteArrayUpdated := util.InterfaceToByte(historyJsonArray)

	// Write the data to the history file
	WriteDataToFile("./history.json", historyByteArrayUpdated)
}
