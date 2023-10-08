package database

import (
	"encoding/json"
	"math"
	"time"

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
		NET_WORTH: math.Round(util.NET_WORTH*100) / 100,
		INCOME:    math.Round(util.INCOME*100) / 100,
		BALANCE:   math.Round(util.BALANCE*100) / 100,
		EXPENSES:  math.Round(util.EXPENSES*100) / 100,
		MONTH:     now.Month(),
	}
}

// GetFinanceJson reads the finance data from a file and stores it in variables
func GetFinanceJson() {
	// Read the finance data from a file
	byteArray := util.ReadFile("./finance.json")

	// Initialize finance JSON object
	financeJsonObject := finance{}

	// Unmarshal byte array to finance JSON object
	err := json.Unmarshal(byteArray, &financeJsonObject)
	util.HandleError(err)

	// Store the values from the FinanceJsonObject in variables
	util.NET_WORTH = financeJsonObject.NET_WORTH
	util.BALANCE = financeJsonObject.BALANCE
	util.EXPENSES = financeJsonObject.EXPENSES
	util.INCOME = financeJsonObject.INCOME

	// Calculate the perfect save amount
	util.SAVING = util.INCOME * 0.25
}
