package database

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/colors"
	"github.com/VkHyperNova/VK-FINANCE/pkg/config"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func (h *History) PrintCLI() {

	util.ClearScreen()

	fmt.Println(colors.Cyan, "<============= VK FINANCE v1.1 =============>\n", colors.Reset)

	fmt.Println(colors.Yellow, "\nIncome", colors.Reset)
	h.PrintIncome()

	fmt.Println(colors.Yellow, "\nExpences", colors.Reset)
	h.PrintExpences()

	fmt.Println(colors.Yellow, "\nSummary", colors.Reset)
	h.PrintStats()

	fmt.Println(colors.Yellow, "\n[history, day, backup, q]", colors.Reset)
}

func (h *History) PrintIncome() {

	for _, item := range config.INCOME {

		itemValue := h.CountItemValue(item)
		itemValueTwoDecimalPlaces := fmt.Sprintf("%.2f", itemValue)

		fmt.Print(colors.Cyan, strings.ToUpper(item)+": ", colors.Reset)
		fmt.Print(colors.Green, "+"+itemValueTwoDecimalPlaces+" EUR\n", colors.Reset)
	}
}

func (h *History) PrintExpences() {

	for _, item := range config.EXPENCES {

		itemValue := h.CountItemValue(item)
		itemValueTwoDecimalPlaces := fmt.Sprintf("%.2f", itemValue)

		if itemValue > 0 {
			fmt.Print(colors.Cyan, strings.ToUpper(item)+": ", colors.Reset)
			fmt.Print(colors.Green, "+"+itemValueTwoDecimalPlaces+" EUR\n", colors.Reset)
		} else {
			fmt.Print(colors.Cyan, strings.ToUpper(item)+": ", colors.Reset)
			fmt.Print(colors.Green, itemValueTwoDecimalPlaces+" EUR\n", colors.Reset)
		}
	}
}

func (h *History) PrintStats() {

	income := 0.0
	expenses := 0.0

	for _, item := range h.History {
		if item.VALUE < 0 {
			expenses += item.VALUE
		} else {
			income += item.VALUE
		}

	}

	config.OLDBALANCE = income + expenses // income + (-expenses)

	fmt.Print(colors.Cyan, "INCOME: ", colors.Reset)
	fmt.Println(colors.Green, "+"+fmt.Sprintf("%.2f", income)+" EUR", colors.Reset)

	fmt.Print(colors.Cyan, "EXPENSES: ", colors.Reset)
	fmt.Println(colors.Red, fmt.Sprintf("%.2f", expenses)+" EUR", colors.Reset)

	fmt.Print(colors.Cyan, "BALANCE: ", colors.Reset)

	msg := fmt.Sprintf("%.2f", income+expenses) + " EUR"
	if income+expenses < 0 {
		fmt.Println(colors.Red, msg, colors.Reset)
	} else {
		fmt.Println(colors.Green, msg, colors.Reset)
	}
}

func (h *History) PrintHistory() {

	fmt.Println("History: ")

	now := time.Now()

	for _, value := range h.History {
		val, err := json.Marshal(value.VALUE)
		if err != nil {
			log.Fatal(err)
		}

		msg := " " + value.DATE + " " + value.TIME + " " + value.COMMENT + " " + string(val)

		if value.DATE == now.Format("02-01-2006") {
			fmt.Println(colors.Green, msg, colors.Reset)
		} else {
			fmt.Println(msg)
		}

	}

	util.PressAnyKey()
}

func (h *History) PrintDaySummary() {

	DaySpent := make(map[time.Time]float64)

	for _, item := range h.History {
		date, err := time.Parse("02-01-2006", item.DATE)
		if err != nil {
			fmt.Println("Error parsing date:", err)
		}
		DaySpent[date] += item.VALUE
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
	fmt.Println(colors.Cyan, "DAY SUMMARY", colors.Reset)

	for _, kv := range keyValueSlice {
		fmt.Print(colors.Purple, "("+kv.Key.Format("02-01-2006")+") ", colors.Reset)
		fmt.Print(kv.Key.Weekday().String() + ": ")
		fmt.Println(colors.Red, fmt.Sprintf("%.2f", kv.Value), colors.Reset)
	}
}

func (h *History) PrintItemSummary() {

	item := ""

	searchID := len(h.History) - 1
	for key, value := range h.History {
		if key == searchID {
			item = value.COMMENT
		}
	}

	sumOfItem := 0.0
	for _, value := range h.History {
		if strings.EqualFold(value.COMMENT, item) {
			sumOfItem += value.VALUE
		}
	}

	msg := fmt.Sprintf("\n"+item+": %.2f", sumOfItem) + " EUR"

	fmt.Println(colors.Yellow, msg, colors.Reset)
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
