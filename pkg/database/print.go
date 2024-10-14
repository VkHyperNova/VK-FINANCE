package database

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/colors"
	"github.com/VkHyperNova/VK-FINANCE/pkg/config"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func (h *History) PrintCLI() {

	util.ClearScreen()

	fmt.Println(colors.Bold, "\nVK FINANCE v1.2\n", colors.Reset)

	h.PrintSummary()

	fmt.Println(colors.Bold, "\nhistory day stats backup quit", colors.Reset)
}

func (h *History) PrintItems(items []string, highlightName string) {

	for _, name := range items {

		sum := h.Calculate(name)

		itemSum := sum[0]

		item := "\t" + name + ": " + fmt.Sprintf("%.2f", itemSum) + " EUR"

		var pMsg string

		// Highlight the specified name
		if name == highlightName {
			pMsg = colors.Bold + colors.Yellow + item + colors.Reset

			// Positive values
		} else if itemSum > 0 {
			pMsg = colors.Green + item + colors.Reset

			// Negative values
		} else if itemSum < 0 {
			pMsg = colors.Red + item + colors.Reset
		} else {
			pMsg = item
		}

		fmt.Println(pMsg)
	}
}

func (h *History) PrintSummary() {

	values := h.Calculate("")

	income := values[1]
	expenses := values[2]

	// PRINT INCOME
	fmt.Print(colors.Green, "\tINCOME: ", colors.Reset)
	fmt.Println(colors.Green, "+"+fmt.Sprintf("%.2f", income)+" EUR", colors.Reset)

	// PRINT EXPENSES
	fmt.Print(colors.Red, "\tEXPENSES: ", colors.Reset)
	fmt.Println(colors.Red, fmt.Sprintf("%.2f", expenses)+" EUR", colors.Reset)

	// PRINT BALANCE
	name := "\tBALANCE: "
	sum := fmt.Sprintf("%.2f", income + expenses) + " EUR"
	fmt.Println(colors.Bold,name + sum, colors.Reset)
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

	fmt.Println(colors.Bold + colors.Yellow, "\n\tStatistics\n", colors.Reset)

	h.PrintItems(config.IncomeItems, "")
	fmt.Println()

	h.PrintItems(config.ExpensesItems, "")
	fmt.Println()

	h.PrintSummary()

	util.PressAnyKey()
}

func (h *History) PrintMessage(name string) {

	util.ClearScreen()

	if util.ArrayContainsString(config.IncomeItems, name) {
		h.PrintItems(config.IncomeItems, name)
	}

	if util.ArrayContainsString(config.ExpensesItems, name) {
		h.PrintItems(config.ExpensesItems, name)
	}

	util.PressAnyKey()
}
