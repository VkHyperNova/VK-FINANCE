package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/dir"
	"github.com/VkHyperNova/VK-FINANCE/pkg/global"
	"github.com/VkHyperNova/VK-FINANCE/pkg/print"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func CMD() {

	print.ClearScreen()

	dir.ValidateRequiredFiles()

	print.PrintGray("============================================\n")
	print.PrintGray("============== VK FINANCE v1 ===============\n")
	print.PrintGray("============================================\n")

	SetFinanceStats()
	PrintSortedHistory()
	print.PrintStats()

	print.PrintGray("--------------------------------------------\n")
	
	print.PrintCommands()

	var user_input string
	print.PrintGray("\n\n=> ")
	fmt.Scanln(&user_input)

	for {
		switch user_input {
		case "add":
			AddIncome()
			CMD()
		case "spend":
			AddExpenses()
			CMD()
		case "history":
			PrintHistory()
			CMD()
		case "q":
			print.ClearScreen()
			os.Exit(0)
		default:
			CMD()
		}
	}
}

func AddIncome() {

	sum := util.UserInputFloat64("Sum: ")
	comment := util.UserInputString("Comment: ")

	global.LastAdd += sum

	database.SaveDatabase(sum, comment)
}

func AddExpenses() {

	sum := util.UserInputFloat64("Sum: ")
	comment := util.UserInputString("Comment: ")

	global.LastExp += sum

	database.SaveDatabase(-1*sum, comment)
}

func PrintHistory() {

	db := database.OpenDatabase()

	print.PrintCyan("History: \n\n")

	for _, value := range db {

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

	fmt.Scanln() // Press enter to continue
}

func PrintSortedHistory() {

	db := database.OpenDatabase()

	var items []string

	for _, value := range db {
		if !util.Contains(items, value.COMMENT) {
			items = append(items, value.COMMENT)
		}
	}

	myMap := make(map[string]float64)

	for _, itemName := range items {
		for _, value := range db {
			if itemName == value.COMMENT {
				myMap[itemName] += value.VALUE

			}
		}
	}

	pairs := make([][2]interface{}, 0, len(myMap))
	for k, v := range myMap {
		pairs = append(pairs, [2]interface{}{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1].(float64) < pairs[j][1].(float64)
	})

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

func SetFinanceStats() {

	db := database.OpenDatabase()

	income := 0.0
	expenses := 0.0

	for _, item := range db {
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
