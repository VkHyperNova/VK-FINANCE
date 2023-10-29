package database

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/dir"
	"github.com/VkHyperNova/VK-FINANCE/pkg/print"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

type history struct {
	DATE    string  `json:"date"`
	TIME    string  `json:"time"`
	COMMENT string  `json:"comment"`
	VALUE   float64 `json:"value"`
}

// SetHistoryJson creates a history object with the current date, time, action, and value.
func SetHistoryJson(value float64, comment string) history {

	now := time.Now()

	return history{
		DATE:    now.Format("02-01-2006"),
		TIME:    now.Format("15:04:05"),
		COMMENT: comment,
		VALUE:   value,
	}
}

// Converts a byte array to a history json array.
func GetHistoryJson(byteArray []byte) []history {

	// initialize historyJsonArray as an empty slice of history
	historyJsonArray := []history{}

	// unmarshal byteArray to historyJsonArray
	err := json.Unmarshal(byteArray, &historyJsonArray)
	print.HandleError(err)

	// return historyJsonArray
	return historyJsonArray
}
// Rethink this <------------------------------------------------
func CountAndPrintHistoryItems() {
	byteArray := dir.ReadFile("./history.json")
	historyJson := GetHistoryJson(byteArray)

	// Get all the names from history json
	var items []string

	for _, value := range historyJson {
		if !util.Contains(items, value.COMMENT) {
			items = append(items, value.COMMENT)
		}
	}

	/* Count the value by name */
	myMap := make(map[string]float64)

	for _, itemName := range items {
		for _, value := range historyJson {
			if itemName == value.COMMENT {
				myMap[itemName] += value.VALUE

			}
		}
	}

	// Create slice of key-value pairs
	pairs := make([][2]interface{}, 0, len(myMap))
	for k, v := range myMap {
		pairs = append(pairs, [2]interface{}{k, v})
	}

	// Sort slice based on values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1].(float64) < pairs[j][1].(float64)
	})

	// Extract sorted keys
	keys := make([]string, len(pairs))
	for i, p := range pairs {
		keys[i] = p[0].(string)
	}

	print.PrintCyan("\nINCOME\n")
	for _, k := range keys {
		if myMap[k] > 0 {
			print.PrintGreen(k + ": " + fmt.Sprintf("%f", myMap[k]) + "\n")
		}

	}

	print.PrintCyan("\nEXPENSES\n")
	for _, k := range keys {
		if myMap[k] < 0 {
			print.PrintRed(k + ": " + fmt.Sprintf("%f", myMap[k]) + "\n")
		}

	}

}

func PrintHistory() {

	byteArray := dir.ReadFile("./history.json")
	historyJson := GetHistoryJson(byteArray)

	print.ClearScreen()

	print.PrintCyan("History: \n\n")

	for _, value := range historyJson {

		val, err := json.Marshal(value.VALUE)
		print.HandleError(err)

		if value.VALUE < 0 {
			print.PrintRed(" ")
			print.PrintRed(value.DATE)
			print.PrintRed(" ")
			print.PrintRed(value.TIME)
			print.PrintRed(" ")
			print.PrintRed(value.COMMENT)
			print.PrintRed(" ==> ")
			print.PrintRed(string(val) + "\n")
		} else {
			print.PrintGreen(" ")
			print.PrintGreen(value.DATE)
			print.PrintGreen(" ")
			print.PrintGreen(value.TIME)
			print.PrintGreen(" ")
			print.PrintGreen(value.COMMENT)
			print.PrintGreen(" ==> ")
			print.PrintGreen(string(val) + "\n")
		}

	}

}
