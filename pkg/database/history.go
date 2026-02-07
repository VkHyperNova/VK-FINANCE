package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/config"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

/* Database Functions */

type FinanceItem struct {
	DATE    string  `json:"date"`
	TIME    string  `json:"time"`
	COMMENT string  `json:"comment"`
	VALUE   float64 `json:"value"`
}

// Slice containing multiple HistoryItem instances.
type Finance struct {
	Finance []FinanceItem `json:"vk-finance"`
}

func (h *Finance) ReadFile() {

	file, err := os.Open(config.DefaultPath)
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

func (h *Finance) Insert(item string, value float64) bool {

	comment := strings.ToLower(item)

	// Check for random items
	if !util.Contains(config.IncomeItems, comment) && !util.Contains(config.ExpensesItems, comment) {
		fmt.Println(config.Red, "No such item!", config.Reset)
		util.PressAnyKey()
		return false
	}

	// Assign +/- if it's not dept
	if util.Contains(config.ExpensesItems, comment) && comment != "dept" {
		value = -value
	}

	// Add time
	now := time.Now()

	NewItem := FinanceItem{
		DATE:    now.Format("02-01-2006"),
		TIME:    now.Format("15:04:05"),
		COMMENT: comment,
		VALUE:   value,
	}

	// Append New Item
	h.Finance = append(h.Finance, NewItem)

	// Convert to json
	byteArray, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Save to main path
	util.WriteToFile(config.DefaultPath, byteArray)

	// D-drive Backup
	util.WriteToFile(config.BackupPath, byteArray)

	return true
}

func (h *Finance) Backup() bool {
	answer := util.Input("Did you take a picture?(y/n)")
	if answer == "n" {
		fmt.Println(config.Bold+config.Red, "\n\tBackup Canceled!\n", config.Reset)
		util.PressAnyKey()
		return false
	}


	// Calculate old dept
	_, _, oldBalance := h.Calculate()

	// Remove old file
	err := os.Remove(config.DefaultPath)
	if err != nil {
		fmt.Println(err)
	}

	// New vk-finance.json
	util.WriteToFile(config.DefaultPath, []byte(`{"vk-finance": []}`))

	// Open new Empty DB
	h.ReadFile()

	// Append old balance
	h.Insert("dept", oldBalance)

	fmt.Println(config.Bold+config.Green, "\n\tBackup Done!\n", config.Reset)

	util.PressAnyKey()

	return true
}

func (h *Finance) Calculate() (float64, float64, float64) {

	income := 0.0
	expenses := 0.0

	for _, item := range h.Finance {

		if item.VALUE < 0 {
			expenses += item.VALUE
		} else {
			income += item.VALUE
		}
	}

	return income, expenses, income + expenses
}

func (h *Finance) Undo() bool {

	// Remove the last item
	h.Finance = h.Finance[:len(h.Finance)-1]

	// Convert to json
	byteArray, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Save to main path
	util.WriteToFile(config.DefaultPath, byteArray)

	// D-drive Backup
	util.WriteToFile(config.BackupPath, byteArray)

	return true
}
