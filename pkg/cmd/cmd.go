package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

/* Main Functions */

func CMD() {

	util.ClearScreen()

	util.ValidateRequiredFiles()

	util.PrintGray("============================================\n")
	util.PrintGray("============== VK FINANCE v1.1 =============\n")
	util.PrintGray("============================================\n")

	db := database.OpenDatabase()

	PrintSortedHistory(db)
	PrintFinanceStats(db)

	util.PrintGray("\n\n--------------------------------------------\n")

	PrintCommands([]string{"add", "spend", "history", "backup", "q"})

	var user_input string
	util.PrintGray("\n\n=> ")
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
			PrintHistory(db)
			CMD()
		case "backup":
			Backup(db)
			CMD()
		case "q":
			util.ClearScreen()
			os.Exit(0)
		default:
			CMD()
		}
	}
}


func AddIncome() {

	comment := util.UserInputString("Comment: ")

	if comment == "q" {
		CMD()
	}

	sum := util.UserInputFloat64("Add Sum: ")

	database.SaveDatabase(sum, comment)
}

func AddExpenses() {

	comment := util.UserInputString("Comment: ")

	if comment == "q" {
		CMD()
	}

	sum := util.UserInputFloat64("Spend Sum: ")

	database.SaveDatabase(-1*sum, comment)
}

var RESTART_BALANCE float64

func Backup(db []database.History) {
	currentTime := time.Now()
	previousMonth := currentTime.AddDate(0, -1, 0).Format("January2006")

	byteArray, err := json.MarshalIndent(db, "", " ")
	util.HandleError(err)

	util.WriteDataToFile("./history/history_json/"+previousMonth+".json", byteArray)

	util.RemoveFile("./history.json")
	util.WriteDataToFile("./history.json", []byte("[]"))

	database.SaveDatabase(RESTART_BALANCE, "oldbalance")
}

func PrintHistory(db []database.History) {

	util.PrintCyan("History: \n\n")

	for _, value := range db {

		val, err := json.Marshal(value.VALUE)
		util.HandleError(err)

		if value.VALUE < 0 {
			util.PrintRed(" " + value.DATE + " " + value.TIME + " " + value.COMMENT + " " + string(val) + "\n")
		} else {
			util.PrintGreen(" " + value.DATE + " " + value.TIME + " " + value.COMMENT + " " + string(val) + "\n")
		}

	}
	util.PrintPurple("\n\nPress ENTER to continue!")
	fmt.Scanln() // Press enter to continue
}

func PrintSortedHistory(db []database.History) {

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

	util.PrintCyan("\nINCOME\n")
	for _, k := range keys {
		if myMap[k] > 0 {
			util.PrintGreen(k + ": " + fmt.Sprintf("%f", myMap[k]) + "\n")
		}

	}

	util.PrintCyan("\nEXPENSES\n")
	for _, k := range keys {
		if myMap[k] < 0 {
			util.PrintRed(k + ": " + fmt.Sprintf("%f", myMap[k]) + "\n")
		}

	}
}

func PrintFinanceStats(db []database.History) {

	myStats := SetFinanceStats(db)

	util.PrintCyan("\nNET WORTH: ")
	util.PrintGreen(fmt.Sprintf("%.2f", myStats["NET_WORTH"]) + " EUR\n\n")

	util.PrintCyan("INCOME: ")
	util.PrintGreen("+" + fmt.Sprintf("%.2f", myStats["INCOME"]) + " EUR")

	

	util.PrintCyan("\nEXPENSES: ")
	util.PrintRed(fmt.Sprintf("%.2f", myStats["EXPENSES"]) + " EUR")

	

	util.PrintCyan("\n\nDay Budget: ")
	util.PrintGreen(fmt.Sprintf("%.2f", myStats["DayBudget"]) + " EUR")
	util.PrintCyan(" | ")
	util.PrintRed(fmt.Sprintf("%.2f", myStats["DayBudgetSpent"]) + " EUR")
	
	util.PrintCyan("\nWeek Budget: ")
	util.PrintGreen(fmt.Sprintf("%.2f", myStats["WeekBudget"]) + " EUR")
	util.PrintCyan(" | ")
	util.PrintRed(fmt.Sprintf("%.2f", myStats["WeekBudgetSpent"]) + " EUR")

	

	util.PrintCyan("\nSAVING (25%): ")
	util.PrintGreen(fmt.Sprintf("%.2f", myStats["SAVING"]) + " EUR")


	util.PrintCyan("\n\nBALANCE: ")
	util.PrintYellow(fmt.Sprintf("%.2f", myStats["BALANCE"]) + " EUR")


	util.PrintCyan("\n\nBudget: ")
	if myStats["Budget"] < 0 {
		util.PrintRed(fmt.Sprintf("%.2f", myStats["Budget"]) + " EUR")
	} else {
		util.PrintGreen(fmt.Sprintf("%.2f", myStats["Budget"]) + " EUR")
	}
}

func SetFinanceStats(db []database.History) map[string]float64 {
	income := 0.0
	expenses := 0.0

	for _, item := range db {
		if item.VALUE < 0 {
			expenses += item.VALUE
		} else {
			income += item.VALUE
		}

	}

	myStats := make(map[string]float64)

	NET_WORTH := 1300.0
	myStats["NET_WORTH"] = NET_WORTH

	INCOME := income
	myStats["INCOME"] = INCOME

	EXPENSES := expenses
	myStats["EXPENSES"] = EXPENSES

	BALANCE := income + expenses // income + (-expenses)
	myStats["BALANCE"] = BALANCE
	RESTART_BALANCE = BALANCE

	SAVING := income * 0.25
	myStats["SAVING"] = SAVING

	Budget := BALANCE - SAVING
	myStats["Budget"] = Budget

	DayBudget := (INCOME - SAVING) / 31
	myStats["DayBudget"] = DayBudget

	DayBudgetSpent := EXPENSES / 31
	myStats["DayBudgetSpent"] = DayBudgetSpent

	WeekBudget := ((INCOME - SAVING) / 31) * 7
	myStats["WeekBudget"] = WeekBudget

	WeekBudgetSpent := (EXPENSES / 31) * 7
	myStats["WeekBudgetSpent"] = WeekBudgetSpent

	return myStats
}

func PrintCommands(commands []string) {

	util.PrintCyan("Program Options: \n\n")

	for _, value := range commands {
		util.PrintCyan("[")
		util.PrintYellow(value)
		util.PrintCyan("] ")
	}
}
