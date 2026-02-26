package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/color"
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
	Finance []FinanceItem `json:"finance"`
}

func (h *Finance) ReadFromFile() error {

	file, err := os.Open(config.LocalFile)
	if err != nil {
		return err
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, h)
	if err != nil {
		return err
	}

	return nil
}

func (h *Finance) Insert(item string, value float64) error {

	comment := strings.ToLower(item)

	// Check for random items
	if !util.Contains(config.AllItems, comment) {
		return errors.New("No such Item!")
	}

	// Assign +/- if it's not dept
	if !util.Contains(config.IncomeItems, comment) && comment != "dept" {
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

	err := h.Save()
	if err != nil {
		return err
	}

	return nil
}

func (h *Finance) Backup() error {

	// Ask before backup start
	answer := util.Input("Did you take a picture?(y/n)")
	if answer == "n" {
		fmt.Println(color.Bold+color.Red, "\n\tBackup Canceled!\n", color.Reset)
		util.PressAnyKey()
		return nil
	}

	// Convert to json
	finance, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return err
	}

	// Save a copy
	err = os.WriteFile(config.BackupFileWithDate, finance, 0644)
	if err != nil {
		return err
	}

	// Calculate old dept
	_, _, oldBalance := h.Calculate()

	// Remove old file
	err = os.Remove(config.LocalFile)
	if err != nil {
		return err
	}

	// New vk-finance.json
	err = os.WriteFile(config.LocalFile, []byte(config.DefaultContent), 0644)
	if err != nil {
		return err
	}

	// Open new Empty DB
	h.ReadFromFile()

	// Append old balance
	h.Insert("dept", oldBalance)

	fmt.Println(color.Bold+color.Green, "\n\tBackup Done!\n", color.Reset)

	util.PressAnyKey()

	return nil
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

func (h *Finance) Undo() error {

	// Remove the last item
	h.Finance = h.Finance[:len(h.Finance)-1]

	err := h.Save()
	if err != nil {
		return err
	}

	return nil
}

func (h *Finance) Save() error {

	finance, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(config.LocalFile, finance, 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile(config.BackupFile, finance, 0644)
	if err != nil {
		return err
	}

	return nil
}
