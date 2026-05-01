package db

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/color"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func (f *Finance) PrintDashboard() {

	util.ClearScreen()

	fmt.Print(color.Blue + color.Bold + "\nVK FINANCE v1.6\n\n" + color.Reset)

	// Print current month
	currentMonth := time.Now().AddDate(0, -1, 0).Format("January 2006")
	fmt.Println(color.Yellow + "\t" + currentMonth + "\n" + color.Reset)

	income, expenses, balance := f.calculate()
	fmt.Println(color.Green, "INCOME: "+"+"+strconv.FormatFloat(income, 'f', 2, 64)+" EUR", color.Reset)
	fmt.Println(color.Red, "EXPENSES: "+strconv.FormatFloat(expenses, 'f', 2, 64)+" EUR", color.Reset)
	fmt.Println(color.Bold, "BALANCE: "+strconv.FormatFloat(balance, 'f', 2, 64)+" EUR", color.Reset)
	fmt.Println()

	f.PrintItemsBySum()

	fmt.Print(color.Blue + "\n< history, undo, import, export, restart, unmount, quit >" + color.Reset)
	fmt.Print("\n=> ")
}

func (f *Finance) PrintItemsBySum() {

	// Guard against empty slice panic
    if len(f.Finance) == 0 {
        return
    }

    // Group and Sum up Items
    totals := make(map[string]float64)
    for _, item := range f.Finance {
        totals[item.COMMENT] += item.VALUE
    }

    // Convert to sortable slice, skipping zero-sum entries
    type itemSum struct {
        name string
        sum  float64
    }
    pairs := make([]itemSum, 0, len(totals))
    for name, sum := range totals {
        if sum != 0 {
            pairs = append(pairs, itemSum{name, sum})
        }
    }

    // Sort highest to lowest
    sort.Slice(pairs, func(a, b int) bool {
        return pairs[a].sum > pairs[b].sum
    })

    // Last added item is used for highlight
    lastItem := f.Finance[len(f.Finance)-1]

    for _, p := range pairs {
        line := fmt.Sprintf("\t%s: %.2f EUR", p.name, p.sum)

        if p.name == lastItem.COMMENT {
            line += fmt.Sprintf(" | %.2f", lastItem.VALUE)
        }

        highlight := p.name == lastItem.COMMENT
        fmt.Println(util.Colorize(line, p.sum, highlight))
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
