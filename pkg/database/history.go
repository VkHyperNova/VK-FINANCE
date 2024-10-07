package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

/* Database Functions */

type HistoryItem struct {
	DATE    string  `json:"date"`
	TIME    string  `json:"time"`
	COMMENT string  `json:"comment"`
	VALUE   float64 `json:"value"`
}

type History struct {
	History []HistoryItem `json:"history"` // Slice containing multiple Quote instances.
}

/* Constants */
var Path = "./history.json"
var BackupPath = "/media/veikko/VK DATA/DATABASES/FINANCE/history.json"
var HistoryPath = "/media/veikko/VK DATA/DATABASES/FINANCE/" + time.Now().AddDate(0, -1, 0).Format("January2006") + ".json"

var INCOME = []string{"pension", "wolt", "bolt", "muu"}
var EXPENCES = []string{"arved", "food", "catfood", "saun", "bensiin", "vape", "w", "other", "old balance", "correction"}
var OLDBALANCE float64

/* Colors */
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
)

func (h *History) Read() {

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

func (h *History) Save() bool {

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

func (h *History) Append(comment string, sum float64) {
	now := time.Now()

	NewItem := HistoryItem{
		DATE:    now.Format("02-01-2006"),
		TIME:    now.Format("15:04:05"),
		COMMENT: comment,
		VALUE:   sum,
	}

	h.History = append(h.History, NewItem)
}

func (h *History) Split(userInput string) bool {

	input := strings.TrimSpace(userInput)
	parts := strings.Fields(input)

	// Check if the input contains exactly two parts
	if len(parts) != 2 {
		return false
	}

	item := strings.ToLower(parts[0])

	// Try to convert the second part to a float
	sum, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		fmt.Println(Red, err, Reset)
		util.PressAnyKey()
		return false
	}

	// Adjust sum if item is in expenses category
	if util.ArrayContainsString(EXPENCES, item) {
		sum = -sum
	}

	// Append to history if item is in either category
	if util.ArrayContainsString(INCOME, item) || util.ArrayContainsString(EXPENCES, item) {
		h.Append(item, sum)
		return true
	} else {
		fmt.Println("No such item!")
		util.PressAnyKey()
		return false
	}
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
	h.Read()

	// Append old balance
	h.Append("old balance", OLDBALANCE)
	h.Save()

	util.PressAnyKey()
}
