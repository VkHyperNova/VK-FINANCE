package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func AddIncome(db []database.History) {

	item := GetComment()
	sum := GetSum()

	database.SaveDatabase(sum, item)

	sumOfItem := FindDBItem(db, item) + sum
	util.PrintCyanString("\n" + item + " = ")
	util.PrintGreenString(fmt.Sprintf("%.2f", sumOfItem) + " EUR")

	util.PressAnyKey()
}

func AddExpenses(db []database.History) {

	item := GetComment()
	sum := GetSum()

	database.SaveDatabase(-1*sum, item)

	sumOfItem := FindDBItem(db, item) - sum
	util.PrintCyanString("\n" + item + " = ")
	util.PrintRedString(fmt.Sprintf("%.2f", sumOfItem) + " EUR")

	util.PressAnyKey()
}

func FindDBItem(db []database.History, comment string) float64 {
	sumOfItem := 0.0

	for _, value := range db {
		if strings.EqualFold(value.COMMENT, comment) {
			sumOfItem += value.VALUE
		}
	}

	return sumOfItem
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

func Backup(db []database.History) {
	currentTime := time.Now()
	previousMonth := currentTime.AddDate(0, -1, 0).Format("January2006")

	byteArray, err := json.MarshalIndent(db, "", " ")
	util.HandleError(err)

	util.WriteDataToFile("./history/history_json/"+previousMonth+".json", byteArray)

	util.RemoveFile("./history.json")
	util.WriteDataToFile("./history.json", []byte("[]"))

	database.SaveDatabase(BACKUP_BALANCE, "oldbalance")
	util.PressAnyKey()
}
