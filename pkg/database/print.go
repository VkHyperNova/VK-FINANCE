package database

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/config"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func (h *Finance) PrintCLI() {

	util.ClearScreen()

	fmt.Print("\nVK FINANCE v1.3\n\n")

	income, expenses, balance := h.Calculate()
	fmt.Println(config.Green, "INCOME: "+"+"+strconv.FormatFloat(income, 'f', 2, 64)+" EUR", config.Reset)
	fmt.Println(config.Red, "EXPENSES: "+strconv.FormatFloat(expenses, 'f', 2, 64)+" EUR", config.Reset)
	fmt.Println(config.Bold, "BALANCE: "+strconv.FormatFloat(balance, 'f', 2, 64)+" EUR", config.Reset)
	fmt.Println()

	h.PrintItems(config.IncomeItems)
	h.PrintItems(config.ExpensesItems)

	fmt.Print("\n\nhistory undo backup quit")
}

func (h *Finance) PrintItems(items []string) {
	totals := make(map[string]float64)

	for _, item := range h.Finance {
		totals[item.COMMENT] += item.VALUE
	}

	for _, name := range items {
		value := totals[name]

		// 👇 skip zero values
		if value == 0 {
			continue
		}

		line := fmt.Sprintf("\t%s: %.2f EUR", name, value)

		if name == config.LastAddedItemName {
			line = config.Bold + line + config.Reset
		}

		switch {
		case value > 0:
			line = config.Green + line + config.Reset
		case value < 0:
			line = config.Red + line + config.Reset
		}

		fmt.Println(line)
	}
}


func (h *Finance) PrintHistory() {

	util.ClearScreen()

	fmt.Println(config.Bold+config.Yellow, "\n\t\tHistory: \n", config.Reset)

	now := time.Now()

	for _, value := range h.Finance {
		val, err := json.Marshal(value.VALUE)
		if err != nil {
			log.Fatal(err)
		}

		time := " " + value.TIME + " | "
		date := value.DATE + " "
		sum := fmt.Sprint(config.Bold, value.COMMENT+" ", string(val)+" EUR", config.Reset)

		if value.VALUE > 0 {
			sum = fmt.Sprint(config.Bold+config.Green, value.COMMENT+" ", string(val)+" EUR", config.Reset)
		}

		if value.VALUE < 0 {
			sum = fmt.Sprint(config.Bold+config.Red, value.COMMENT+" ", string(val)+" EUR", config.Reset)
		}

		msg := time + date + sum

		if value.DATE == now.Format("02-01-2006") {
			fmt.Println(config.Bold, msg, config.Reset)
		} else {
			fmt.Println(msg)
		}
	}
	util.PressAnyKey()
}

