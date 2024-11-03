package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

// Slice containing multiple HistoryItem instances.
type History struct {
	History []HistoryItem `json:"history"`
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

func (h *History) Save(item string, value float64) bool {

	comment := strings.ToLower(item)

	// Check for random items
	if !util.Contains(config.IncomeItems, comment) && !util.Contains(config.ExpensesItems, comment) {
		comment = "other"
	}

	// Assign +/- if it's not dept
	if util.Contains(config.ExpensesItems, comment) && comment != "dept" {
		value = -value
	}

	// Add time
	now := time.Now()
	NewItem := HistoryItem{
		DATE:    now.Format("02-01-2006"),
		TIME:    now.Format("15:04:05"),
		COMMENT: comment,
		VALUE:   value,
	}

	// Append New Item
	h.History = append(h.History, NewItem)

	// Convert to json
	byteArray, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Save to main path
	util.WriteToFile(config.Path, byteArray)

	// D-drive Backup
	util.WriteToFile(config.BackupPath, byteArray)

	// Print Income/Expense items
	h.PrintMessage()

	return true
}

func (h *History) Backup() {

	_, _, oldBalance := h.Calculate()

	byteArray, err := json.MarshalIndent(h, "", " ")
	if err != nil {
		panic(err)
	}

	// D-drive history by month Backup
	util.WriteToFile(config.HistoryPath, byteArray)

	// New history file
	err = os.Remove(config.Path)
	if err != nil {
		fmt.Println(err)
	}

	// New history.json
	util.WriteToFile(config.Path, []byte(`{"history": []}`))

	// Open new Empty DB
	h.Read()

	// Append old balance
	h.Save("dept", oldBalance)

	fmt.Println(colors.Bold+colors.Green, "\n\tBackup Done!\n", colors.Reset)

	util.PressAnyKey()
}

func (h *History) Calculate() (float64, float64, float64) {

	income := 0.0
	expenses := 0.0

	for _, item := range h.History {

		if item.VALUE < 0 {
			expenses += item.VALUE
		} else {
			income += item.VALUE
		}
	}

	return income, expenses, income + expenses
}
