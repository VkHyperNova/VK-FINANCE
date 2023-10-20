package database

import (
	"encoding/json"
	"math"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/dir"
	"github.com/VkHyperNova/VK-FINANCE/pkg/global"
	"github.com/VkHyperNova/VK-FINANCE/pkg/print"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

type finance struct {
	NET_WORTH float64    `json:"net_worth"`
	INCOME    float64    `json:"income"`
	BALANCE   float64    `json:"balance"`
	EXPENSES  float64    `json:"expences"`
	MONTH     time.Month `json:"month"`
}

// GetFinanceData returns a finance struct with financial data
func SetFinanceJson() finance {

	now := time.Now()

	return finance{
		NET_WORTH: math.Round(global.NET_WORTH*100) / 100,
		INCOME:    math.Round(global.INCOME*100) / 100,
		BALANCE:   math.Round(global.BALANCE*100) / 100,
		EXPENSES:  math.Round(global.EXPENSES*100) / 100,
		MONTH:     now.Month(),
	}
}

// GetFinanceJson reads the finance data from a file and stores it in variables
func GetFinanceJson() {
	// Read the finance data from a file
	byteArray := dir.ReadFile("./finance.json")

	// Initialize finance JSON object
	financeJsonObject := finance{}

	// Unmarshal byte array to finance JSON object
	err := json.Unmarshal(byteArray, &financeJsonObject)
	print.HandleError(err)

	// Store the values from the FinanceJsonObject in variables
	global.NET_WORTH = financeJsonObject.NET_WORTH
	global.BALANCE = financeJsonObject.BALANCE
	global.EXPENSES = financeJsonObject.EXPENSES
	global.INCOME = financeJsonObject.INCOME

	// Calculate the perfect save amount
	global.SAVING = global.INCOME * 0.25
}

func CalculateIncome() {

	sum := util.UserInputFloat64("Sum: ")
	comment := util.UserInputString("Comment: ")

	global.LastAdd += sum

	global.INCOME = global.INCOME + sum
	global.BALANCE = global.BALANCE + sum

	Save(sum, comment)

}

func CalculateExpenses() {

	sum := util.UserInputFloat64("Sum: ")
	comment := util.UserInputString("Comment: ")

	global.LastExp += sum

	global.BALANCE = global.BALANCE - sum
	global.EXPENSES = global.EXPENSES - sum

	Save(-1*sum, comment)

}

func AddNetWorth() {
	global.NET_WORTH = util.UserInputFloat64("NET_WORTH: ")
	Save(global.NET_WORTH, "Net Worth")
}

func Grow() {
	global.NET_WORTH = global.NET_WORTH + global.BALANCE
	SAVED_BALANCE := global.BALANCE
	global.BALANCE = 0

	ResetVariables()
	Save(SAVED_BALANCE, "Grow")

}
func ResetVariables() {
	global.BALANCE = 0
	global.INCOME = 0
	global.EXPENSES = 0
}

func DayBudget() {
	global.DayBudget = (global.INCOME - global.SAVING) / 31
	global.DayBudgetSpent = global.EXPENSES / 31
}

func WeekBudget() {
	global.WeekBudget = ((global.INCOME - global.SAVING) / 31) * 7
	global.WeekBudgetSpent = (global.EXPENSES / 31) * 7
}

func Budget() {
	global.Budget = global.BALANCE - global.SAVING
}
