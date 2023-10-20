package database

import (
	"encoding/json"
	"fmt"
	"github.com/VkHyperNova/VK-FINANCE/pkg/dir"
	"github.com/VkHyperNova/VK-FINANCE/pkg/print"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
	"sort"
	"time"
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

func CountAndPrintHistoryItems() {
	byteArray := dir.ReadFile("./history.json")
	historyJson := GetHistoryJson(byteArray)

	// Get all the names
	var items []string

	for _, value := range historyJson {
		if !util.Contains(items, value.COMMENT) {
			items = append(items, value.COMMENT)
		}
	}

	/* Count */
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
		stringValue := fmt.Sprintf("%f", myMap[k]) // convert to string

		if myMap[k] > 0 {
			print.PrintGreen(k + ": " + stringValue + "\n")
		}

	}

	// Print sorted map Expenses
	print.PrintCyan("\nEXPENSES\n")
	for _, k := range keys {
		stringValue := fmt.Sprintf("%f", myMap[k]) // convert to string

		if myMap[k] < 0 {
			print.PrintRed(k + ": " + stringValue + "\n")
		}

	}

}

func PrintCurrentMonthHistory() {
	now := time.Now()
	CurrentMonth := now.Month()

	byteArray := dir.ReadFile("./history.json")
	historyJson := GetHistoryJson(byteArray)

	print.ClearScreen()

	print.PrintCyan("History: \n\n")

	for _, value := range historyJson {
		// Define date layout format
		layout := "02-01-2006"

		// Parse date string to time.Time object
		t, err := time.Parse(layout, value.DATE)
		// Handle error if any
		print.HandleError(err)

		// Check if the month of the current date is equal to the current month
		if t.Month() == CurrentMonth {
			// Print the value

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

}
