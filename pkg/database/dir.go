package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

var MainPath = "./history.json"
var MainPathBackup = "/media/veikko/VK DATA/DATABASES/FINANCE/history.json"

var currentTime = time.Now()
var previousMonth = currentTime.AddDate(0, -1, 0).Format("January2006")
var HistoryPathBackup = "/media/veikko/VK DATA/DATABASES/FINANCE/" + previousMonth + ".json"


func (h *History) ReadFromFile() {


	file, err := os.Open(MainPath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, h)
	if err != nil {
		panic(err)
	}
}

func (h *History) SaveToFile() bool {

	byteArray, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Save to main path
	util.WriteDataToFile(MainPath, byteArray)

	// D-drive Backup
	util.WriteDataToFile(MainPathBackup, byteArray)

	return true
}

func (h *History) Backup() {
	

	byteArray, err := json.MarshalIndent(h.History, "", " ")
	if err != nil {
		panic(err)
	}

	// D-drive history by month Backup
	util.WriteDataToFile(HistoryPathBackup, byteArray)

	// New history file
	err = os.Remove(MainPath)
	if err != nil {
		panic(err)
	}	

	// New history.json
	util.WriteDataToFile(MainPath, []byte(`{"history": []}`))

	// Add old balance
	h.Append("old balance", OLDBALANCE)
	h.SaveToFile()

	util.PressAnyKey()
}
