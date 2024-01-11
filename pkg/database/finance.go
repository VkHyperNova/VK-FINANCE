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

func CompileHistoryEntry(value float64, comment string) History {
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
	NewItem := CompileHistoryEntry(Value, Comment)
	db = append(db, NewItem)
	byteArray, err := json.MarshalIndent(db, "", " ")
	util.HandleError(err)
	util.WriteDataToFile("./history.json", byteArray)
	util.PrintGreen("\n<< Success! >>\n")
}

var RESTART_BALANCE float64
func SetFinanceStats(db []History) map[string]float64 {
	income := 0.0
	expenses := 0.0
	for _, item := range db {
		if item.VALUE < 0 {
			expenses += item.VALUE
		} else {
			income += item.VALUE
		}

	}
	myStats := make(map[string]float64)

	NET_WORTH := 1300.0
	myStats["NET_WORTH"] = NET_WORTH

	INCOME := income
	myStats["INCOME"] = INCOME

	EXPENSES := expenses
	myStats["EXPENSES"] = EXPENSES

	BALANCE := income + expenses // income + (-expenses)
	myStats["BALANCE"] = BALANCE
	RESTART_BALANCE = BALANCE

	SAVING := income * 0.25
	myStats["SAVING"] = SAVING

	Budget := BALANCE - SAVING
	myStats["Budget"] = Budget

	DayBudget := (INCOME - SAVING) / 31
	myStats["DayBudget"] = DayBudget

	DayBudgetSpent := EXPENSES / 31
	myStats["DayBudgetSpent"] = DayBudgetSpent

	WeekBudget := ((INCOME - SAVING) / 31) * 7
	myStats["WeekBudget"] = WeekBudget

	WeekBudgetSpent := (EXPENSES / 31) * 7
	myStats["WeekBudgetSpent"] = WeekBudgetSpent

	return myStats
}
