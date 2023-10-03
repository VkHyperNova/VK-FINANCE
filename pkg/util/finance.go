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
func Finance() finance {

	// Get the current time
	now := time.Now()

	// Return a finance struct with financial data rounded to 2 decimal places
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


