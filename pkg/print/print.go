package print

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"github.com/VkHyperNova/VK-FINANCE/pkg/global"
)

func ClearScreen() {
	if runtime.GOOS == "linux" { // check if the operating system is Linux
		cmd := exec.Command("clear") // execute the clear command
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") // execute the cls command
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func PrintStats() {
	PrintNetWorth()
	PrintIncome()
	PrintExpences()
	PrintEstimatedDaylySpendingAmount()
	PrintEstimatedWeeklySpendingAmount()
	PrintSavingAmount()
	PrintBalance()
	PrintMoneyLeft()
}

func PrintCommands() {
	PrintCyan("Program Options: \n\n")
	PrintWithBrackets("add")
	PrintWithBrackets("spend")
	PrintWithBrackets("history")
	PrintWithBrackets("q")
}

func PrintNetWorth() {
	PrintCyan("\nNET WORTH: ")
	PrintGreen(fmt.Sprintf("%.2f", global.NET_WORTH) + " EUR\n\n")
}

func PrintIncome() {
	PrintCyan("INCOME: ")
	PrintGreen("+" + fmt.Sprintf("%.2f", global.INCOME) + " EUR")

	if global.LastAdd != 0 {
		PrintCyan(" | ")
		PrintYellow("+" + fmt.Sprintf("%.2f", global.LastAdd) + " EUR")
	}
	PrintGray("\n")

}

func PrintExpences() {
	PrintCyan("EXPENCES: ")
	PrintRed(fmt.Sprintf("%.2f", global.EXPENSES) + " EUR")

	if global.LastExp != 0 {
		PrintCyan(" | ")
		PrintYellow("+" + fmt.Sprintf("%.2f", global.LastExp) + " EUR")
	}
	PrintGray("\n\n")
}

func PrintEstimatedDaylySpendingAmount() {
	PrintCyan("ESTIMATED DAY: ")
	PrintGreen(fmt.Sprintf("%.2f", global.DayBudget) + " EUR")
	PrintCyan(" | ")
	PrintRed(fmt.Sprintf("%.2f", global.DayBudgetSpent) + " EUR\n")
}

func PrintEstimatedWeeklySpendingAmount() {
	PrintCyan("ESTIMATED WEEK: ")
	PrintGreen(fmt.Sprintf("%.2f",global.WeekBudget ) + " EUR")
	PrintCyan(" | ")
	PrintRed(fmt.Sprintf("%.2f", global.WeekBudgetSpent) + " EUR\n")
}

func PrintSavingAmount() {
	PrintCyan("SAVING (25%): ")
	PrintGreen(fmt.Sprintf("%.2f", global.SAVING) + " EUR\n\n")
}

func PrintBalance() {
	PrintCyan("BALANCE: ")
	PrintYellow(fmt.Sprintf("%.2f", global.BALANCE) + " EUR\n")
}

func PrintMoneyLeft() {

	PrintCyan("\nUsable Money: ")

	if global.Budget < 0 {
		PrintRed(fmt.Sprintf("%.2f", global.Budget) + " EUR\n\n")
	} else {
		PrintGreen(fmt.Sprintf("%.2f", global.Budget) + " EUR\n\n")
	}
}

func PrintWithBrackets(name string) {
	PrintCyan("[")
	PrintYellow(name)
	PrintCyan("] ")
}

func PrintSeparatorSingleDash() {
	PrintGray("============================================\n")
}

func PrintSeparatorDoubleDash() {
	PrintGray("--------------------------------------------\n")
}
