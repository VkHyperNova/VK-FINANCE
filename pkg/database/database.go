package database

import (
	"encoding/json"
	"time"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

/* Database Functions */

type History struct {
	DATE    string  `json:"date"`
	TIME    string  `json:"time"`
	COMMENT string  `json:"comment"`
	VALUE   float64 `json:"value"`
}

func NewItem(value float64, comment string) History {

	now := time.Now()

	return History{
		DATE:    now.Format("02-01-2006"),
		TIME:    now.Format("15:04:05"),
		COMMENT: comment,
		VALUE:   value,
	}
}

func OpenDatabase() []History {

	OpenFile := util.ReadFile("./history.json")

	JsonArray := []History{}

	err := json.Unmarshal(OpenFile, &JsonArray)
	util.HandleError(err)

	return JsonArray
}

func SaveDatabase(Value float64, Comment string) {

	db := OpenDatabase()

	NewItem := NewItem(Value, Comment)

	db = append(db, NewItem)

	byteArray, err := json.MarshalIndent(db, "", " ")
	util.HandleError(err)

	util.WriteDataToFile("./history.json", byteArray)

	util.PrintGreen("\n<< Success! >>\n")
}
