package db

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/color"
	"github.com/VkHyperNova/VK-FINANCE/pkg/config"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func (f *Finance) PrintDashboard() {

	util.ClearScreen()

	fmt.Print("\nVK FINANCE v1.5\n\n")

	// Print month
	currentMonth := time.Now().AddDate(0, -1, 0).Format("January 2006")
	fmt.Println("\t" + currentMonth + "\n")

	income, expenses, balance := f.calculate()
	fmt.Println(color.Green, "INCOME: "+"+"+strconv.FormatFloat(income, 'f', 2, 64)+" EUR", color.Reset)
	fmt.Println(color.Red, "EXPENSES: "+strconv.FormatFloat(expenses, 'f', 2, 64)+" EUR", color.Reset)
	fmt.Println(color.Bold, "BALANCE: "+strconv.FormatFloat(balance, 'f', 2, 64)+" EUR", color.Reset)
	fmt.Println()

	f.PrintItems(config.AllItems)

	fmt.Print("\n\n< history, undo, import, backup, unmount, quit >")
	fmt.Print("\n=> ")
}

func (f *Finance) PrintItems(items []string) {
	totals := make(map[string]float64)
	for _, item := range f.Finance {
		totals[item.COMMENT] += item.VALUE
	}

	for _, name := range items {
		value := totals[name]
		if value == 0 {
			continue
		}

		line := fmt.Sprintf("\t%s: %.2f EUR", name, value)
		highlight := name == config.LastAddedItemName
		if highlight {
			line += " | " + strconv.FormatFloat(config.LastAddedItemSum, 'f', 2, 64)
		}

		fmt.Println(util.Colorize(line, value, highlight))
	}

}

func (f *Finance) PrintHistory() {

	util.ClearScreen()

	fmt.Println(color.Bold+color.Yellow, "\n\t\tHistory: \n", color.Reset)

	now := time.Now()

	for _, value := range f.Finance {
		val, err := json.Marshal(value.VALUE)
		if err != nil {
			log.Fatal(err)
		}

		time := " " + value.TIME + " | "
		date := value.DATE + " "
		sum := fmt.Sprint(color.Bold, value.COMMENT+" ", string(val)+" EUR", color.Reset)

		if value.VALUE > 0 {
			sum = fmt.Sprint(color.Bold+color.Green, value.COMMENT+" ", string(val)+" EUR", color.Reset)
		}

		if value.VALUE < 0 {
			sum = fmt.Sprint(color.Bold+color.Red, value.COMMENT+" ", string(val)+" EUR", color.Reset)
		}

		msg := time + date + sum

		if value.DATE == now.Format("02-01-2006") {
			fmt.Println(color.Bold, msg, color.Reset)
		} else {
			fmt.Println(msg)
		}
	}
	util.PressAnyKey()
}
