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

func (h *History) Save(parts []string) bool {

	comment := strings.ToLower(parts[0])

	// Check for random items
	if !util.Contains(config.IncomeItems, comment) && !util.Contains(config.ExpensesItems, comment) {
		comment = "other"
	}

	// Try to convert the second part to a float
	value, err := strconv.ParseFloat(parts[1], 64)

	if err != nil {
		fmt.Println(err)
		return false
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

	values := h.Calculate()
	income := values[0]
	expenses := values[1]

	// Save old balance
	dept := strconv.FormatFloat(income+expenses, 'f', 2, 64)

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
	var itemParts []string
	itemParts = append(itemParts, "dept", dept)
	h.Save(itemParts)

	fmt.Println(colors.Bold+colors.Green, "\n\tBackup Done!\n", colors.Reset)

	util.PressAnyKey()
}

func (h *History) Calculate() []float64 {

	totalIncome := 0.0
	totalExpenses := 0.0

	for _, item := range h.History {

		if item.VALUE < 0 {
			totalExpenses += item.VALUE
		} else {
			totalIncome += item.VALUE
		}
	}

	values := []float64{totalIncome, totalExpenses}

	return values
}
