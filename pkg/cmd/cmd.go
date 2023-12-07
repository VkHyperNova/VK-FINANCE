package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/dir"
	"github.com/VkHyperNova/VK-FINANCE/pkg/print"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

/* Main Functions */

func CMD() {

	print.ClearScreen()
	
	dir.ValidateRequiredFiles()

	print.PrintGray("============================================\n")
	print.PrintGray("============== VK FINANCE v1.1 ===============\n")
	print.PrintGray("============================================\n")

	PrintSortedHistory()
	PrintFinanceStats()
	
	print.PrintGray("--------------------------------------------\n")

	PrintCommands([]string{"add", "spend", "history", "backup", "q"})

	var user_input string
	print.PrintGray("\n\n=> ")
	fmt.Scanln(&user_input)

	for {
		switch user_input {
		case "add", "a":
			AddIncome()
			CMD()
		case "spend", "s":
			AddExpenses()
			CMD()
		case "history", "h":
			PrintHistory()
			CMD()
		case "backup":
			Backup()
			CMD()
		case "q":
			print.ClearScreen()
			os.Exit(0)
		default:
			CMD()
		}
	}
}

var LastAdd float64
func AddIncome() {

	sum := util.UserInputFloat64("Add Sum: ")
	comment := util.UserInputString("Comment: ")

	LastAdd += sum

	database.SaveDatabase(sum, comment)
}

var LastExp float64
func AddExpenses() {

	sum := util.UserInputFloat64("Spend Sum: ")
	comment := util.UserInputString("Comment: ")

	LastExp += sum

	database.SaveDatabase(-1*sum, comment)
}

var RESTART_BALANCE float64
func Backup() {
	currentTime := time.Now()
	previousMonth := currentTime.AddDate(0, -1, 0).Format("January2006")

	db := database.OpenDatabase()
	byteArray, err := json.MarshalIndent(db, "", " ")
	print.HandleError(err)

	dir.WriteDataToFile("./history/history_json/" + previousMonth + ".json", byteArray)

	dir.RemoveFile("./history.json")
	dir.WriteDataToFile("./history.json", []byte("[]"))

	database.SaveDatabase(RESTART_BALANCE, "oldbalance")
}

func PrintHistory() {

	db := database.OpenDatabase()

	print.PrintCyan("History: \n\n")

	for _, value := range db {

		val, err := json.Marshal(value.VALUE)
		print.HandleError(err)

		if value.VALUE < 0 {
			print.PrintRed(" " + value.DATE + " " + value.TIME + " " + value.COMMENT + " " + string(val) + "\n")
		} else {
			print.PrintGreen(" " + value.DATE + " " + value.TIME + " " + value.COMMENT + " " + string(val) + "\n")
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

func PrintFinanceStats() {

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

	NET_WORTH := 1300.0
	INCOME := income
	EXPENSES := expenses
	
	BALANCE := income + expenses // income + (-expenses)
	RESTART_BALANCE = BALANCE

	SAVING := income * 0.25
	Budget := BALANCE - SAVING
	DayBudget := (INCOME - SAVING) / 31
	DayBudgetSpent := EXPENSES / 31
	WeekBudget := ((INCOME - SAVING) / 31) * 7
	WeekBudgetSpent := (EXPENSES / 31) * 7

	print.PrintCyan("\nNET WORTH: ")
	print.PrintGreen(fmt.Sprintf("%.2f", NET_WORTH) + " EUR\n\n")

	print.PrintCyan("INCOME: ")
	print.PrintGreen("+" + fmt.Sprintf("%.2f", INCOME) + " EUR")

	if LastAdd != 0 {
		print.PrintCyan(" | ")
		print.PrintYellow("+" + fmt.Sprintf("%.2f", LastAdd) + " EUR")
	}
	print.PrintGray("\n")

	print.PrintCyan("EXPENSES: ")
	print.PrintRed(fmt.Sprintf("%.2f", EXPENSES) + " EUR")

	if LastExp != 0 {
		print.PrintCyan(" | ")
		print.PrintYellow("+" + fmt.Sprintf("%.2f", LastExp) + " EUR")
	}
	print.PrintGray("\n\n")

	print.PrintCyan("Day Budget: ")
	print.PrintGreen(fmt.Sprintf("%.2f", DayBudget) + " EUR")
	print.PrintCyan(" | ")
	print.PrintRed(fmt.Sprintf("%.2f", DayBudgetSpent) + " EUR\n")

	print.PrintCyan("Week Budget: ")
	print.PrintGreen(fmt.Sprintf("%.2f", WeekBudget) + " EUR")
	print.PrintCyan(" | ")
	print.PrintRed(fmt.Sprintf("%.2f", WeekBudgetSpent) + " EUR\n")

	print.PrintCyan("SAVING (25%): ")
	print.PrintGreen(fmt.Sprintf("%.2f", SAVING) + " EUR\n\n")

	print.PrintCyan("BALANCE: ")
	print.PrintYellow(fmt.Sprintf("%.2f", BALANCE) + " EUR\n")

	print.PrintCyan("\nUsable Money: ")

	if Budget < 0 {
		print.PrintRed(fmt.Sprintf("%.2f", Budget) + " EUR\n\n")
	} else {
		print.PrintGreen(fmt.Sprintf("%.2f", Budget) + " EUR\n\n")
	}
}

func PrintCommands(commands []string) {

	print.PrintCyan("Program Options: \n\n")

	for _, value := range commands {
		print.PrintCyan("[")
		print.PrintYellow(value)
		print.PrintCyan("] ")
	}
}
