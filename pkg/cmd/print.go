package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func PrintCLI(db []database.History) string {

	util.PrintCyanString("<============= VK FINANCE v1.1 ============>\n\n")

	PrintIncomeByType(db)
	util.PrintCyanString("\n============== Expences ====================\n\n")
	PrintExpencesByType(db)
	util.PrintCyanString("\n============== Summary =====================\n\n")
	PrintFinanceStats(db)
	fmt.Println("\n")

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

var BACKUP_BALANCE float64

func PrintFinanceStats(db []database.History) {
	
	income := 0.0
	expenses := 0.0

	for _, item := range db {
		if item.VALUE < 0 {
			expenses += item.VALUE
		} else {
			income += item.VALUE
		}

	}

	BACKUP_BALANCE = income + expenses // income + (-expenses)

	util.PrintCyanString("INCOME: ")
	util.PrintGreenString("+" + fmt.Sprintf("%.2f", income) + " EUR")

	util.PrintCyanString("\nEXPENSES: ")
	util.PrintRedString(fmt.Sprintf("%.2f", expenses) + " EUR")

	util.PrintCyanString("\nBALANCE: ")
	util.PrintGreenString(fmt.Sprintf("%.2f", income + expenses) + " EUR")
}

func PrintExpencesByType(db []database.History) {

	importantExpences := []string{"arved", "food", "saun", "bensiin", "e-smoke", "weed", "other", "oldbalance"}

	for _, item := range importantExpences {
		itemValue := CountItemValue(item, db)
		itemValueTwoDecimalPlaces := fmt.Sprintf("%.2f", itemValue)
		if itemValue > 0 {
			util.PrintCyanString(strings.ToUpper(item) + ": ")
			util.PrintGreenString("+" + itemValueTwoDecimalPlaces + " EUR\n")
		} else {
			util.PrintCyanString(strings.ToUpper(item) + ": ")
			util.PrintRedString(itemValueTwoDecimalPlaces + " EUR\n")
		}
	}
}

func PrintIncomeByType(db []database.History) {

	importantIncome := []string{"pension", "sotsiaal", "wolt", "bolt", "muu"}

	for _, item := range importantIncome {
		itemValue := CountItemValue(item, db)
		itemValueTwoDecimalPlaces := fmt.Sprintf("%.2f", itemValue)
		util.PrintCyanString(strings.ToUpper(item) + ": ")
		util.PrintGreenString("+" + itemValueTwoDecimalPlaces + " EUR\n")
		
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
