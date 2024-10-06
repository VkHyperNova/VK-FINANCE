package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)



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
		panic(err)
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
