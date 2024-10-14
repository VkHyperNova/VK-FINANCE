package database

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

func (h *History) Save(name string, sum float64) bool {

	if !util.ArrayContainsString(config.IncomeItems, name) && !util.ArrayContainsString(config.ExpensesItems, name) {
		fmt.Println("No such item: " + name)
		util.PressAnyKey()
		return false
	}

	now := time.Now()

	NewItem := HistoryItem{
		DATE:    now.Format("02-01-2006"),
		TIME:    now.Format("15:04:05"),
		COMMENT: name,
		VALUE:   sum,
	}

	h.History = append(h.History, NewItem)

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

func (h *History) Backup() {

	values := h.Calculate("")
	income := values[1]
	expenses := values[2]

	config.DEPT = income + expenses // income + (-expenses)

	byteArray, err := json.MarshalIndent(h.History, "", " ")
	if err != nil {
		panic(err)
	}

	// D-drive history by month Backup
	util.WriteDataToFile(config.HistoryPath, byteArray)

	// New history file
	err = os.Remove(config.Path)
	if err != nil {
		fmt.Println(err)
	} 

	// New history.json
	util.WriteDataToFile(config.Path, []byte(`{"history": []}`))

	// Open new Empty DB
	h.Read()

	// Append old balance
	h.Save("dept", config.DEPT)

	fmt.Println(colors.Bold + colors.Green, "\n\tBackup Done!\n", colors.Reset)

	util.PressAnyKey()
}

func (h *History) Calculate(name string) []float64 {

	sum := 0.0
	income := 0.0
	expenses := 0.0

	for _, item := range h.History {

		if item.COMMENT == name {
			sum += item.VALUE
		}

		if item.VALUE < 0 {
			expenses += item.VALUE
		} else {
			income += item.VALUE
		}
	}

	values := []float64{sum, income, expenses}

	return values
}
