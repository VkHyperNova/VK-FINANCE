package cmd

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)


func PrintVKFINANCE(db []database.History) string{

	util.PrintGray("============================================\n")
	util.PrintGray("============== VK FINANCE v1.1 =============\n")
	util.PrintGray("============================================\n")
	PrintSortedHistory(db)
	PrintFinanceStats(db)
	util.PrintGray("\n\n============================================\n")

	commands := [5]string{"add", "spend", "history", "backup", "q"}
	for _, value := range commands {
		util.PrintCyan("[")
		util.PrintYellow(value)
		util.PrintCyan("] ")
	}

	var input string
	util.PrintGray("\n\n=> ")
	fmt.Scanln(&input)

	return input
}
func PrintFinanceStats(db []database.History) {

	myStats := database.SetFinanceStats(db)

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
	util.PressAnyKey()
}
