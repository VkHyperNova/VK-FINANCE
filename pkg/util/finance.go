package util

import (
	"encoding/json"
	"math"
	"time"
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
		NET_WORTH: math.Round(NET_WORTH*100) / 100,
		INCOME:    math.Round(INCOME*100) / 100,
		BALANCE:   math.Round(BALANCE*100) / 100,
		EXPENSES:  math.Round(EXPENSES*100) / 100,
		MONTH:     now.Month(),
	}
}

// Convert byte array to finance JSON object
func GetFinanceJson(byteArray []byte) finance {

	// Initialize finance JSON object
	financeJsonObject := finance{}

	// Unmarshal byte array to finance JSON object
	err := json.Unmarshal(byteArray, &financeJsonObject)
	HandleError(err)

	// Return finance JSON object
	return financeJsonObject
}


// FetchFinanceDataFromFile reads the finance data from a file and stores it in variables
func FetchFinanceDataFromFile() {
	// Read the finance data from a file
	byteArray := ReadFile("./finance.json")

	// Convert the byte array to a FinanceJsonObject
	financeJsonObject := GetFinanceJson(byteArray)

	// Store the values from the FinanceJsonObject in variables
	NET_WORTH = financeJsonObject.NET_WORTH
	BALANCE = financeJsonObject.BALANCE
	EXPENSES = financeJsonObject.EXPENSES
	INCOME = financeJsonObject.INCOME

	// Calculate the perfect save amount
	SAVING = INCOME * 0.25
}