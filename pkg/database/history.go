package database

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/dir"
	"github.com/VkHyperNova/VK-FINANCE/pkg/global"
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
func CountIncomeAndExpenses() {
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

func GetFinances() {

	byteArray := dir.ReadFile("./history.json")
	historyJson := GetHistoryJson(byteArray)

	income := 0.0
	expenses := 0.0

	for _, item := range historyJson {
		if item.VALUE < 0 {
			expenses += item.VALUE
		} else {
			income += item.VALUE
		}

	}

	global.INCOME = income
	global.EXPENSES = expenses
	global.BALANCE = income + expenses // income + (-expenses)
	global.SAVING = income * 0.25
	global.Budget = global.BALANCE - global.SAVING
	global.DayBudget = (global.INCOME - global.SAVING) / 31
	global.DayBudgetSpent = global.EXPENSES / 31
	global.WeekBudget = ((global.INCOME - global.SAVING) / 31) * 7
	global.WeekBudgetSpent = (global.EXPENSES / 31) * 7
}
