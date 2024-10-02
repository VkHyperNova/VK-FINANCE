package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func (h *History) ReadFromFile() {

	path := "./history.json"

	file, err := os.Open(path)
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

	// Save
	path := "./history.json"
	util.WriteDataToFile(path, byteArray)

	// D-drive Backup
	dPath := "/media/veikko/VK DATA/DATABASES/FINANCE/history.json"
	util.WriteDataToFile(dPath, byteArray)

	return true
}

func (h *History) Backup() {
	currentTime := time.Now()
	previousMonth := currentTime.AddDate(0, -1, 0).Format("January2006")

	byteArray, err := json.MarshalIndent(h.History, "", " ")
	if err != nil {
		panic(err)
	}

	// Main Backup
	path := "./history/" + previousMonth + ".json"
	util.WriteDataToFile(path, byteArray)

	// D-drive Backup
	dPath := "/media/veikko/VK DATA/DATABASES/FINANCE/" + previousMonth + ".json"
	util.WriteDataToFile(dPath, byteArray)

	// New history file
	err = os.Remove("./history.json")
	if err != nil {
		panic(err)
	}	
	util.WriteDataToFile("./history.json", []byte(`{"history": []}`))

	// Add old balance
	h.Append("old balance", OLDBALANCE)
	h.SaveToFile()

	util.PressAnyKey()
}
