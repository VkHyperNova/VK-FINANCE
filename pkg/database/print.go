package database

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

var INCOMESOURCES = []string{"pension", "sotsiaal", "wolt", "bolt", "muu"}
var MAINEXPENCES = []string{"arved", "food", "saun", "bensiin", "e-smoke", "weed", "other", "oldbalance", "correction"}
var OLDBALANCE float64

func (h *History) PrintCLI() {

	util.ClearScreen()

	util.PrintCyanString("<============= VK FINANCE v1.1 ============>\n\n")

	h.PrintIncomeItems()
	util.PrintCyanString("\n============== Expences ====================\n\n")
	h.PrintExpencesItems()
	util.PrintCyanString("\n============== Summary =====================\n\n")
	h.PrintFinanceStats()
	util.PrintYellowString("\nadd spend history day backup q")
}

func (h *History) PrintIncomeItems() {

	for _, item := range INCOMESOURCES {

		itemValue := h.CountItemValue(item)
		itemValueTwoDecimalPlaces := fmt.Sprintf("%.2f", itemValue)

		util.PrintCyanString(strings.ToUpper(item) + ": ")
		util.PrintGreenString("+" + itemValueTwoDecimalPlaces + " EUR\n")
	}
}

func (h *History) CountItemValue(item string) float64 {

	itemValue := 0.0

	for _, value := range h.History {
		if item == value.COMMENT {
			itemValue += value.VALUE

		}
	}

	return itemValue
}

func (h *History) PrintExpencesItems() {

	for _, item := range MAINEXPENCES {

		itemValue := h.CountItemValue(item)
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

func (h *History) PrintFinanceStats() {

	income := 0.0
	expenses := 0.0

	for _, item := range h.History {
		if item.VALUE < 0 {
			expenses += item.VALUE
		} else {
			income += item.VALUE
		}

	}

	OLDBALANCE = income + expenses // income + (-expenses)

	util.PrintCyanString("INCOME: ")
	util.PrintGreenString("+" + fmt.Sprintf("%.2f", income) + " EUR")

	util.PrintCyanString("\nEXPENSES: ")
	util.PrintRedString(fmt.Sprintf("%.2f", expenses) + " EUR")

	util.PrintCyanString("\nBALANCE: ")
	if income+expenses < 0 {
		util.PrintRedString(fmt.Sprintf("%.2f", income+expenses) + " EUR\n")
	} else {
		util.PrintGreenString(fmt.Sprintf("%.2f", income+expenses) + " EUR\n")
	}
}

func (h *History) PrintHistory() {
	util.PrintCyanString("History: \n\n")
	for _, value := range h.History {
		val, err := json.Marshal(value.VALUE)
		if err != nil {
			log.Fatal(err)
		}
		if value.VALUE < 0 {
			util.PrintRedString(" " + value.DATE + " " + value.TIME + " " + value.COMMENT + " " + string(val) + "\n")
		} else {
			util.PrintGreenString(" " + value.DATE + " " + value.TIME + " " + value.COMMENT + " " + string(val) + "\n")
		}
	}
	util.PressAnyKey()
}

func (h *History) PrintDailySpending() {

	DaySpent := make(map[time.Time]float64)
	fmt.Println()
	for _, item := range h.History {
		DaySpent[util.GetDayFromString(item.DATE)] += item.VALUE
	}

	type KeyValue struct {
		Key   time.Time
		Value float64
	}

	// Convert the map to a slice of key-value pairs
	var keyValueSlice []KeyValue
	for k, v := range DaySpent {
		keyValueSlice = append(keyValueSlice, KeyValue{k, v})
	}

	// Sort the slice by keys
	sort.Slice(keyValueSlice, func(i, j int) bool {
		return keyValueSlice[i].Key.Before(keyValueSlice[j].Key)
	})

	// Print the sorted map
	util.PrintCyanString("DAY SUMMARY\n")
	for _, kv := range keyValueSlice {
		util.PrintPurpleString("(" + kv.Key.Format("02-01-2006") + ") ")
		util.PrintGrayString(kv.Key.Weekday().String() + ": ")
		util.PrintRedString(fmt.Sprintf("%.2f", kv.Value) + "\n")
	}
	util.PressAnyKey()
}
