package util

import (
	"encoding/json"
	"time"
)

type history struct {
	DATE    string  `json:"date"`
	TIME    string  `json:"time"`
	COMMENT string  `json:"comment"`
	VALUE   float64 `json:"value"`
}

// SetHistoryJson creates a history object with the current date, time, action, and value.
func SetHistoryJson(value float64, comment string) history {

	now := time.Now()

	return history{
		DATE:    now.Format("02-01-2006"),
		TIME:    now.Format("15:04:05"),
		COMMENT: comment,
		VALUE:   value,
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
