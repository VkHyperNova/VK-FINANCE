package database

import (
	"encoding/json"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/dir"
	"github.com/VkHyperNova/VK-FINANCE/pkg/print"
)

/* Database Functions */

type history struct {
	DATE    string  `json:"date"`
	TIME    string  `json:"time"`
	COMMENT string  `json:"comment"`
	VALUE   float64 `json:"value"`
}

func NewItem(value float64, comment string) history {

	now := time.Now()

	return history{
		DATE:    now.Format("02-01-2006"),
		TIME:    now.Format("15:04:05"),
		COMMENT: comment,
		VALUE:   value,
	}
}

func OpenDatabase() []history {

	OpenFile := dir.ReadFile("./history.json")

	JsonArray := []history{}

	err := json.Unmarshal(OpenFile, &JsonArray)
	print.HandleError(err)

	return JsonArray
}

func SaveDatabase(Value float64, Comment string) {

	db := OpenDatabase()

	NewItem := NewItem(Value, Comment)

	db = append(db, NewItem)

	byteArray, err := json.MarshalIndent(db, "", " ")
	print.HandleError(err)

	dir.WriteDataToFile("./history.json", byteArray)
}
