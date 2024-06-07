package cmd

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func Add(db []database.History, income bool) {

	item := GetComment()
	sum := GetSum()

	sumOfItem := 0.0

	if income {
		/* Add */
		database.SaveToDatabase(sum, item)
		sumOfItem = database.FindItemInDatabase(db, item) + sum
	} else {
		/* Spend */
		database.SaveToDatabase(-1*sum, item)
		sumOfItem = database.FindItemInDatabase(db, item) - sum
	}

	util.PrintCyanString("\n" + item + " = ")

	if income {
		/* Print Green */
		util.PrintGreenString(fmt.Sprintf("%.2f", sumOfItem) + " EUR")
	} else {
		/* Print Red */
		util.PrintRedString(fmt.Sprintf("%.2f", sumOfItem) + " EUR")
	}

	util.PressAnyKey()
}

func ShowHistory(db []database.History) {
	util.PrintCyanString("History: \n\n")
	for _, value := range db {
		val, err := json.Marshal(value.VALUE)
		util.HandleError(err)
		if value.VALUE < 0 {
			util.PrintRedString(" " + value.DATE + " " + value.TIME + " " + value.COMMENT + " " + string(val) + "\n")
		} else {
			util.PrintGreenString(" " + value.DATE + " " + value.TIME + " " + value.COMMENT + " " + string(val) + "\n")
		}
	}
	util.PressAnyKey()
}

func ShowDaySpending(db []database.History) {

	DaySpent := make(map[time.Time]float64)
	fmt.Println()
	for _, item := range db {
		DaySpent[util.GetDayFromString(item.DATE)] += item.VALUE
	}

	type KeyValue struct {
		Key   time.Time
		Value float64
	}

	// Convert the map to a slice of key-value pairs
	var keyValueSlice []KeyValue
	for k, v := range DaySpent {
		keyValueSlice = append(keyValueSlice, KeyValue{k, v})
	}

	// Sort the slice by keys
	sort.Slice(keyValueSlice, func(i, j int) bool {
		return keyValueSlice[i].Key.Before(keyValueSlice[j].Key)
	})

	// Print the sorted map
	util.PrintCyanString("DAY SUMMARY\n")
	for _, kv := range keyValueSlice {
		util.PrintPurpleString("(" + kv.Key.Format("02-01-2006") + ") ")
		util.PrintGrayString(kv.Key.Weekday().String() + ": ")
		util.PrintRedString(fmt.Sprintf("%.2f", kv.Value) + "\n")
	}
	util.PressAnyKey()
}

func Backup(db []database.History) {
	currentTime := time.Now()
	previousMonth := currentTime.AddDate(0, -1, 0).Format("January2006")

	byteArray, err := json.MarshalIndent(db, "", " ")
	util.HandleError(err)

	util.WriteDataToFile("./history/history_json/"+previousMonth+".json", byteArray)

	util.RemoveFile("./history.json")
	util.WriteDataToFile("./history.json", []byte("[]"))

	database.SaveToDatabase(BACKUP_BALANCE, "oldbalance")
	util.PressAnyKey()
}

func QuitCheck(s string) {
	if s == "q" || s == "Q" {
		util.PrintRedString("\n<< Command Canceled! >>\n")
		util.PressAnyKey()
		CMD()
	}
}

func GetComment() string {
	comment := util.UserInputString("Comment: ")
	QuitCheck(comment)
	return comment
}

func GetSum() float64 {
start:
	sum := util.UserInputString("Spend Sum: ")
	QuitCheck(sum)

	float, err := strconv.ParseFloat(sum, 64)

	if util.HandleError(err) {
		util.PrintPurpleString("<< Enter a number! >>\n\n")
		goto start
	}

	return float
}
