package database

import (
	"encoding/json"
	"strings"
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

func OpenDatabase() []History {
	OpenFile := util.ReadFile("./history.json")
	JsonArray := []History{}
	err := json.Unmarshal(OpenFile, &JsonArray)
	util.HandleError(err)
	return JsonArray
}

func SaveToDatabase(Value float64, Comment string) {

	db := OpenDatabase()

	now := time.Now()

	NewItem := History{
		DATE:    now.Format("02-01-2006"),
		TIME:    now.Format("15:04:05"),
		COMMENT: Comment,
		VALUE:   Value,
	}

	db = append(db, NewItem)

	byteArray, err := json.MarshalIndent(db, "", " ")

	util.HandleError(err)

	util.WriteDataToFile("./history.json", byteArray)

	util.PrintGreenString("\n<< Success! >>\n")
}

func FindItemInDatabase(db []History, comment string) float64 {

	sumOfItem := 0.0

	for _, value := range db {
		if strings.EqualFold(value.COMMENT, comment) {
			sumOfItem += value.VALUE
		}
	}
	return sumOfItem
}
