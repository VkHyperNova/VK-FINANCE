package database

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
	"github.com/VkHyperNova/VK-FINANCE/pkg/dir"
)

func Save(Value float64, Comment string) {

	financeJsonObject := SetFinanceJson() // Construct finance financeJsonObject as a JSON object

	financeByteArray := util.InterfaceToByte(financeJsonObject) // Convert finance data to a byte array

	dir.WriteDataToFile("./finance.json", financeByteArray) // Write finance data to a JSON file

	historyByteArray := dir.ReadFile("./history.json") // Read the history file content

	historyJsonArray := GetHistoryJson(historyByteArray) // Convert the file content to history data

	historyJsonArrayObject := SetHistoryJson(Value, Comment) // Construct a new history JSON object

	historyJsonArray = append(historyJsonArray, historyJsonArrayObject) // Append the new data to the history data

	historyByteArrayUpdated := util.InterfaceToByte(historyJsonArray) // Convert the history data to a byte array

	dir.WriteDataToFile("./history.json", historyByteArrayUpdated) // Write the data to the history file
}


