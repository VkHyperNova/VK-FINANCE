package cmd

import (
	"encoding/json"
	"strconv"
	"time"
	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func AddIncome() {

	comment := GetComment()
	sum := GetSum()

	database.SaveDatabase(sum, comment)
	util.PressAnyKey()
}

func AddExpenses() {

	comment := GetComment()
	sum := GetSum()
	
	database.SaveDatabase(-1*sum, comment)
	util.PressAnyKey()
}

func QuitCheck(s string) {
	if s == "q" || s == "Q" {
		util.PrintRed("\n<< Command Canceled! >>\n")
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
		util.PrintPurple("<< Enter a number! >>\n\n")
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

	database.SaveDatabase(database.RESTART_BALANCE, "oldbalance")
	util.PressAnyKey()
}