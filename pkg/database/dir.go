package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

var Path = "./history.json"
var BackupPath = "/media/veikko/VK DATA/DATABASES/FINANCE/history.json"
var HistoryPath = "/media/veikko/VK DATA/DATABASES/FINANCE/" + time.Now().AddDate(0, -1, 0).Format("January2006") + ".json"

func (h *History) ReadFromFile() {

	file, err := os.Open(Path)
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
	util.WriteDataToFile(Path, byteArray)

	// D-drive Backup
	util.WriteDataToFile(BackupPath, byteArray)

	return true
}

func (h *History) Backup() {

	byteArray, err := json.MarshalIndent(h.History, "", " ")
	if err != nil {
		panic(err)
	}

	// D-drive history by month Backup
	util.WriteDataToFile(HistoryPath, byteArray)

	// New history file
	err = os.Remove(Path)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\n" + Path + " Removed!")
	}

	// New history.json
	util.WriteDataToFile(Path, []byte(`{"history": []}`))

	// Open new Empty DB
	h.ReadFromFile()

	// Append old balance
	h.Append("old balance", OLDBALANCE)
	h.SaveToFile()


	util.PressAnyKey()
}
