package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func PrintCLI(db []database.History) string {
	util.PrintGrayString("============================================\n")
	util.PrintGrayString("============== VK FINANCE v1.1 =============\n")
	util.PrintGrayString("============================================\n")
	PrintIncomeByType(db)
	util.PrintGrayString("============================================\n")
	PrintExpencesByType(db)
	util.PrintGrayString("============================================\n")
	PrintFinanceStats(db)
	util.PrintGrayString("\n============================================\n")

	commands := [5]string{"add", "spend", "history", "backup", "q"}
	for _, value := range commands {
		util.PrintCyanString("[")
		util.PrintYellowString(value)
		util.PrintCyanString("] ")
	}

	var input string
	util.PrintGrayString("\n\n=> ")
	fmt.Scanln(&input)

	return input
}

func PrintFinanceStats(db []database.History) {

	myStats := database.SetFinanceStats(db)

	util.PrintCyanString("INCOME: ")
	util.PrintGreenString("+" + fmt.Sprintf("%.2f", myStats["INCOME"]) + " EUR")

	util.PrintCyanString("\nEXPENSES: ")
	util.PrintRedString(fmt.Sprintf("%.2f", myStats["EXPENSES"]) + " EUR")

	util.PrintCyanString("\nBALANCE: ")
	util.PrintGreenString("+" + fmt.Sprintf("%.2f", myStats["INCOME"] + myStats["EXPENSES"]) + " EUR")	
}

func PrintExpencesByType(db []database.History) {

	importantExpences := []string{"arved", "food", "trenn", "saun", "bensiin", "e-smoke", "vanemad", "weed", "other"}

	for _, item := range importantExpences {
		itemValue := CountItemValue(item, db)
		util.PrintRedString(item + " ")
		util.PrintRedFloat(itemValue)
		fmt.Println()
	}
}

func PrintIncomeByType(db []database.History) {

	importantIncome := []string{"pension", "sotsiaal", "wolt", "bolt", "vanemad", "muu"}

	for _, item := range importantIncome {
		itemValue := CountItemValue(item, db)
		util.PrintGreenString(item + " ")
		util.PrintGreenFloat(itemValue)
		fmt.Println()
	}
}

func CountItemValue(item string, db []database.History) float64 {

	itemValue := 0.0
	for _, value := range db {
		if item == value.COMMENT {
			itemValue += value.VALUE

		}
	}

	return itemValue
}

func PrintHistory(db []database.History) {
	util.PrintCyanString("History: \n\n")
	for _, value := range db {
		val, err := json.Marshal(value.VALUE)
		util.HandleError(err)
		if value.VALUE < 0 {
			util.PrintRedString(" " + value.DATE + " " + value.TIME + " " + value.COMMENT + " " + string(val) + "\n")
		} else {
			util.PrintGreenString(" " + value.DATE + " " + value.TIME + " " + value.COMMENT + " " + string(val) + "\n")
		}
	}
	util.PressAnyKey()
}
