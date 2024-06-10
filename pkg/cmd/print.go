package cmd

import (
	"fmt"
	"strings"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func PrintCLI(db []database.History) string {

	util.PrintCyanString("<============= VK FINANCE v1.1 ============>\n\n")

	PrintIncomeItems(db)
	util.PrintCyanString("\n============== Expences ====================\n\n")
	PrintExpencesItems(db)
	util.PrintCyanString("\n============== Summary =====================\n\n")
	PrintFinanceStats(db)

	commands := [6]string{"add", "spend", "history", "day", "backup", "q"}
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
	util.PrintGreenString(fmt.Sprintf("%.2f", income + expenses) + " EUR\n")
}

func PrintExpencesItems(db []database.History) {

	importantExpences := []string{"arved", "food", "saun", "bensiin", "e-smoke", "weed", "other", "oldbalance", "correction"}

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

func PrintIncomeItems(db []database.History) {

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


