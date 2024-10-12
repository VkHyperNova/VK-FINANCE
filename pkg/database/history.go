package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/colors"
	"github.com/VkHyperNova/VK-FINANCE/pkg/config"
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


func (h *History) Read() {

	file, err := os.Open(config.Path)
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
	util.WriteDataToFile(config.Path, byteArray)

	// D-drive Backup
	util.WriteDataToFile(config.BackupPath, byteArray)

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
		fmt.Println(colors.Red, err, colors.Reset)
		util.PressAnyKey()
		return false
	}

	// Adjust sum if item is in expenses category
	if util.ArrayContainsString(config.EXPENCES, item) {
		sum = -sum
	}

	// Append to history if item is in either category
	if util.ArrayContainsString(config.INCOME, item) || util.ArrayContainsString(config.EXPENCES, item) {
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
	util.WriteDataToFile(config.HistoryPath, byteArray)

	// New history file
	err = os.Remove(config.Path)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\n" + config.Path + " Removed!")
	}

	// New history.json
	util.WriteDataToFile(config.Path, []byte(`{"history": []}`))

	// Open new Empty DB
	h.Read()

	// Append old balance
	h.Append("old balance", config.OLDBALANCE)
	h.Save()

	util.PressAnyKey()
}
