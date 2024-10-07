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

func (h *History) PrintCLI() {

	util.ClearScreen()

	fmt.Println(Cyan, "<============= VK FINANCE v1.1 =============>\n", Reset)

	fmt.Println(Yellow, "\nIncome", Reset)
	h.PrintIncome()

	fmt.Println(Yellow, "\nExpences", Reset)
	h.PrintExpences()

	fmt.Println(Yellow, "\nSummary", Reset)
	h.PrintStats()

	fmt.Println(Yellow, "\n[history, day, backup, q]", Reset)
}

func (h *History) PrintIncome() {

	for _, item := range INCOME {

		itemValue := h.CountItemValue(item)
		itemValueTwoDecimalPlaces := fmt.Sprintf("%.2f", itemValue)

		fmt.Print(Cyan, strings.ToUpper(item)+": ", Reset)
		fmt.Print(Green, "+"+itemValueTwoDecimalPlaces+" EUR\n", Reset)
	}
}

func (h *History) PrintExpences() {

	for _, item := range EXPENCES {

		itemValue := h.CountItemValue(item)
		itemValueTwoDecimalPlaces := fmt.Sprintf("%.2f", itemValue)

		if itemValue > 0 {
			fmt.Print(Cyan, strings.ToUpper(item)+": ", Reset)
			fmt.Print(Green, "+"+itemValueTwoDecimalPlaces+" EUR\n", Reset)
		} else {
			fmt.Print(Cyan, strings.ToUpper(item)+": ", Reset)
			fmt.Print(Green, itemValueTwoDecimalPlaces+" EUR\n", Reset)
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

	OLDBALANCE = income + expenses // income + (-expenses)

	fmt.Print(Cyan, "INCOME: ", Reset)
	fmt.Println(Green, "+"+fmt.Sprintf("%.2f", income)+" EUR", Reset)

	fmt.Print(Cyan, "EXPENSES: ", Reset)
	fmt.Println(Red, fmt.Sprintf("%.2f", expenses)+" EUR", Reset)

	fmt.Print(Cyan, "BALANCE: ", Reset)

	msg := fmt.Sprintf("%.2f", income+expenses) + " EUR"
	if income+expenses < 0 {
		fmt.Println(Red, msg, Reset)
	} else {
		fmt.Println(Green, msg, Reset)
	}
}

func (h *History) PrintHistory() {
	fmt.Println("History: ")
	for _, value := range h.History {
		val, err := json.Marshal(value.VALUE)
		if err != nil {
			log.Fatal(err)
		}

		msg := " " + value.DATE + " " + value.TIME + " " + value.COMMENT + " " + string(val)
		if value.VALUE < 0 {
			fmt.Println(Red, msg, Reset)
		} else {
			fmt.Println(Green, msg, Reset)
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
	fmt.Println(Cyan, "DAY SUMMARY", Reset)

	for _, kv := range keyValueSlice {
		fmt.Print(Purple, "("+kv.Key.Format("02-01-2006")+") ", Reset)
		fmt.Print(kv.Key.Weekday().String() + ": ")
		fmt.Println(Red, fmt.Sprintf("%.2f", kv.Value), Reset)
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

	msg := fmt.Sprintf("\n\n"+item+": %.2f", sumOfItem) + " EUR"

	if sumOfItem > 0 {
		fmt.Println(Green, msg, Reset)
	} else {
		fmt.Println(Red, msg, Reset)
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
