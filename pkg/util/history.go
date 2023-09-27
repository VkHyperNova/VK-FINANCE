package util

import (
	"encoding/json"
	"time"
)

type history struct {
	DATE   string  `json:"date"`
	TIME   string  `json:"time"`
	ACTION string  `json:"action"`
	VALUE  float64 `json:"value"`
}

// History creates a history object with the current date, time, action, and value.
func History(action string, value float64) history {

	// Get the current time.
	now := time.Now()

	// Format the current time as a string.
	formattedTime := now.Format("15:04:05")

	// Format the current date as a string.
	formattedDate := now.Format("02-01-2006")

	// Return the history object with the current date, time, action, and value.
	return history{
		DATE:   formattedDate,
		TIME:   formattedTime,
		ACTION: action,
		VALUE:  value,
	}
}

// Converts a byte array to a history json array.
func GetHistoryJson(byteArray []byte) []history {

	// initialize historyJsonArray as an empty slice of history
	historyJsonArray := []history{}

	// unmarshal byteArray to historyJsonArray
	err := json.Unmarshal(byteArray, &historyJsonArray)
	HandleError(err)

	// return historyJsonArray
	return historyJsonArray
}
