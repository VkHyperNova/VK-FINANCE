package database

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/colors"
	"github.com/VkHyperNova/VK-FINANCE/pkg/config"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func (h *History) PrintCLI() {

	util.ClearScreen()

	fmt.Print("\nVK FINANCE v1.2\n\n")

	h.PrintSummary()

	fmt.Print("\n\nhistory day stats backup quit")
}

func (h *History) PrintItems(items []string, highlightName string) {

	for _, name := range items {

		itemValue := 0.0

		for _, item := range h.History {
			if item.COMMENT == name {
				itemValue += item.VALUE
			}
		}

		item := "\t" + name + ": " + fmt.Sprintf("%.2f", itemValue) + " EUR"

		var pMsg string

		// Highlight the specified name
		if name == highlightName {
			pMsg = colors.Bold + colors.Yellow + item + colors.Reset

			// Positive values
		} else if itemValue > 0 {
			pMsg = colors.Green + item + colors.Reset

			// Negative values
		} else if itemValue < 0 {
			pMsg = colors.Red + item + colors.Reset
		} else {
			pMsg = item
		}

		fmt.Println(pMsg)
	}
}

func (h *History) PrintSummary() {

	income, expenses, balance := h.Calculate()

	// PRINT INCOME
	fmt.Println(colors.Green, "\tINCOME: "+"+"+strconv.FormatFloat(income, 'f', 2, 64)+" EUR", colors.Reset)

	// PRINT EXPENSES
	fmt.Println(colors.Red, "\tEXPENSES: "+strconv.FormatFloat(expenses, 'f', 2, 64)+" EUR", colors.Reset)

	// PRINT BALANCE
	fmt.Println(colors.Bold, "\tBALANCE: "+strconv.FormatFloat(balance, 'f', 2, 64)+" EUR", colors.Reset)

	// PRINT LAST ADDED ITEMS
	now := time.Now()
	currentDate := now.Format("02-01-2006")

	for _, value := range h.History {
		if value.DATE == currentDate {
			fmt.Print(colors.Green+ "\n=> " +value.COMMENT+": ", strconv.FormatFloat(value.VALUE, 'f', 2, 64), colors.Reset)
		}
	}
}

func (h *History) PrintHistory() {

	util.ClearScreen()

	fmt.Println(colors.Bold+colors.Yellow, "\n\t\tHistory: \n", colors.Reset)

	now := time.Now()

	for _, value := range h.History {
		val, err := json.Marshal(value.VALUE)
		if err != nil {
			log.Fatal(err)
		}

		time := " " + value.TIME + " | "
		date := value.DATE + " "
		sum := fmt.Sprint(colors.Bold, value.COMMENT+" ", string(val)+" EUR", colors.Reset)

		if value.VALUE > 0 {
			sum = fmt.Sprint(colors.Bold+colors.Green, value.COMMENT+" ", string(val)+" EUR", colors.Reset)
		}

		if value.VALUE < 0 {
			sum = fmt.Sprint(colors.Bold+colors.Red, value.COMMENT+" ", string(val)+" EUR", colors.Reset)
		}

		msg := time + date + sum

		if value.DATE == now.Format("02-01-2006") {
			fmt.Println(colors.Bold, msg, colors.Reset)
		} else {
			fmt.Println(msg)
		}
	}
	util.PressAnyKey()
}

func (h *History) PrintDays() {

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

	util.ClearScreen()

	// Print the sorted map
	fmt.Println(colors.Yellow+colors.Bold, "\n\tDaily Spending\n", colors.Reset)

	for _, kv := range keyValueSlice {
		fmt.Print(colors.Bold+colors.Cyan, kv.Key.Format("02-01-2006"), colors.Reset)
		fmt.Print(colors.Bold+colors.Cyan, " "+kv.Key.Weekday().String()+": ", colors.Reset)
		fmt.Println(colors.Bold+colors.Red, fmt.Sprintf("%.2f", kv.Value), colors.Reset)
	}

	util.PressAnyKey()
}

func (h *History) PrintStatistics() {

	util.ClearScreen()

	fmt.Println(colors.Bold+colors.Yellow, "\n\tStatistics\n", colors.Reset)

	h.PrintItems(config.IncomeItems, "")
	fmt.Println()

	h.PrintItems(config.ExpensesItems, "")
	fmt.Println()

	h.PrintSummary()

	util.PressAnyKey()
}

func (h *History) PrintMessage() {

	util.ClearScreen()

	lastAddedItem := h.History[len(h.History)-1]

	if util.Contains(config.IncomeItems, lastAddedItem.COMMENT) {
		h.PrintItems(config.IncomeItems, lastAddedItem.COMMENT)
	}

	if util.Contains(config.ExpensesItems, lastAddedItem.COMMENT) {
		h.PrintItems(config.ExpensesItems, lastAddedItem.COMMENT)
	}

	util.PressAnyKey()
}
