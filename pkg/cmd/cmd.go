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

	db := database.OpenDatabase()

	PrintSortedHistory(db)
	PrintFinanceStats(db)

	print.PrintGray("\n\n--------------------------------------------\n")

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
			PrintHistory(db)
			CMD()
		case "backup":
			Backup(db)
			CMD()
		case "q":
			print.ClearScreen()
			os.Exit(0)
		default:
			CMD()
		}
	}
}

var Last float64

func AddIncome() {

	comment := util.UserInputString("Comment: ")

	if comment == "q" {
		CMD()
	}

	sum := util.UserInputFloat64("Add Sum: ")

	Last += sum

	database.SaveDatabase(sum, comment)
}

func AddExpenses() {

	comment := util.UserInputString("Comment: ")

	if comment == "q" {
		CMD()
	}

	sum := util.UserInputFloat64("Spend Sum: ")

	Last -= sum

	database.SaveDatabase(-1*sum, comment)
}

var RESTART_BALANCE float64

func Backup(db []database.History) {
	currentTime := time.Now()
	previousMonth := currentTime.AddDate(0, -1, 0).Format("January2006")

	byteArray, err := json.MarshalIndent(db, "", " ")
	print.HandleError(err)

	dir.WriteDataToFile("./history/history_json/"+previousMonth+".json", byteArray)

	dir.RemoveFile("./history.json")
	dir.WriteDataToFile("./history.json", []byte("[]"))

	database.SaveDatabase(RESTART_BALANCE, "oldbalance")
}

func PrintHistory(db []database.History) {

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
	print.PrintPurple("\n\nPress ENTER to continue!")
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

func PrintFinanceStats(db []database.History) {

	myStats := SetFinanceStats(db)
	fmt.Println("Last ", Last)

	print.PrintCyan("\nNET WORTH: ")
	print.PrintGreen(fmt.Sprintf("%.2f", myStats["NET_WORTH"]) + " EUR\n\n")

	print.PrintCyan("INCOME: ")
	print.PrintGreen("+" + fmt.Sprintf("%.2f", myStats["INCOME"]) + " EUR")

	if Last > 0 {
		print.PrintYellow(" (")
		print.PrintYellow("+" + fmt.Sprintf("%.2f", Last) + " EUR")
		print.PrintYellow(")")

	}

	print.PrintCyan("\nEXPENSES: ")
	print.PrintRed(fmt.Sprintf("%.2f", myStats["EXPENSES"]) + " EUR")

	if Last < 0 {
		print.PrintYellow(" (")
		print.PrintYellow(fmt.Sprintf("%.2f", Last) + " EUR")
		print.PrintYellow(")")

	}

	print.PrintCyan("\n\nDay Budget: ")
	print.PrintGreen(fmt.Sprintf("%.2f", myStats["DayBudget"]) + " EUR")
	print.PrintCyan(" | ")
	print.PrintRed(fmt.Sprintf("%.2f", myStats["DayBudgetSpent"]) + " EUR")
	if Last != 0 {
		NewIncome := myStats["INCOME"] + Last
		NewSaving := NewIncome * 0.25
		NewDayBudget := (NewIncome - NewSaving) / 31
		difference := NewDayBudget-myStats["DayBudget"] // new budget minus old budget! Important order

		print.PrintYellow(" (")
		if difference > 0 {
			print.PrintYellow(fmt.Sprintf("+%.2f", difference)) 
		} else {
			print.PrintYellow(fmt.Sprintf("%.2f", difference)) 
		}
		
		print.PrintYellow(")")

	}
	

	

	print.PrintCyan("\nWeek Budget: ")
	print.PrintGreen(fmt.Sprintf("%.2f", myStats["WeekBudget"]) + " EUR")
	print.PrintCyan(" | ")
	print.PrintRed(fmt.Sprintf("%.2f", myStats["WeekBudgetSpent"]) + " EUR")
	if Last != 0 {
		NewIncome := myStats["INCOME"] + Last
		NewSaving := NewIncome * 0.25
		NewWeekBudget := ((NewIncome - NewSaving) / 31) * 7
		difference := NewWeekBudget-myStats["DayBudget"] // new budget minus old budget! Important order

		print.PrintYellow(" (")
		if difference > 0 {
			print.PrintYellow(fmt.Sprintf("+%.2f", difference)) 
		} else {
			print.PrintYellow(fmt.Sprintf("%.2f", difference)) 
		}
		
		print.PrintYellow(")")
	}
	

	

	print.PrintCyan("\nSAVING (25%): ")
	print.PrintGreen(fmt.Sprintf("%.2f", myStats["SAVING"]) + " EUR")
	if Last != 0 {
		NewIncome := myStats["INCOME"] + Last
		NewSaving := NewIncome * 0.25
		difference := NewSaving - myStats["SAVING"]

		print.PrintYellow(" (")
		if difference > 0 {
			print.PrintYellow(fmt.Sprintf("+%.2f", difference)) 
		} else {
			print.PrintYellow(fmt.Sprintf("%.2f", difference)) 
		}
		
		print.PrintYellow(")")
	}

	print.PrintCyan("\n\nBALANCE: ")
	print.PrintYellow(fmt.Sprintf("%.2f", myStats["BALANCE"]) + " EUR")
	if Last != 0 {
		NewBalance := myStats["BALANCE"] + Last
		difference := NewBalance - myStats["BALANCE"]

		print.PrintYellow(" (")
		if difference > 0 {
			print.PrintYellow(fmt.Sprintf("+%.2f", difference)) 
		} else {
			print.PrintYellow(fmt.Sprintf("%.2f", difference)) 
		}
		
		print.PrintYellow(")")
	}

	print.PrintCyan("\n\nBudget: ")
	if myStats["Budget"] < 0 {
		print.PrintRed(fmt.Sprintf("%.2f", myStats["Budget"]) + " EUR")
	} else {
		print.PrintGreen(fmt.Sprintf("%.2f", myStats["Budget"]) + " EUR")
	}

	if Last != 0 {
		NewIncome := myStats["INCOME"] + Last
		NewSaving := NewIncome * 0.25
		NewBalance := (myStats["BALANCE"] + Last) - NewSaving
		difference := NewBalance - myStats["BALANCE"]

		print.PrintYellow(" (")
		if difference > 0 {
			print.PrintYellow(fmt.Sprintf("+%.2f", difference)) 
		} else {
			print.PrintYellow(fmt.Sprintf("%.2f", difference)) 
		}
		
		print.PrintYellow(")")
	}

	Last = 0
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

	print.PrintCyan("Program Options: \n\n")

	for _, value := range commands {
		print.PrintCyan("[")
		print.PrintYellow(value)
		print.PrintCyan("] ")
	}
}
