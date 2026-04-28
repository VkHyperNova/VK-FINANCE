package db

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

/* Main */

func (f *Finance) ImportDB(source string) error {

	input, err := util.PromptWithSuggestion("Do you want to import db from d drive? (y/n) ", "n")
	if err != nil {
		return err
	}

	if input == "y" || input == "yes" {
		if err := f.LoadFromFile(source); err != nil {
			return fmt.Errorf("load from file: %w", err)
		}

		if err := f.save(config.LocalFile); err != nil {
			return err
		}
	}

	fmt.Printf("Database imported from %s\n", source)

	return nil
}

func (f *Finance) Add(item string, value float64) error {

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
	f.Finance = append(f.Finance, NewItem)

	// Local Save
	err := f.save(config.LocalFile)
	if err != nil {
		return err
	}

	// Backup save
	err = f.save(config.BackupFile)
	if err != nil {
		return err
	}

	return nil
}

func (f *Finance) Backup() error {

	// Ask before backup start
	answer := util.Input("Did you take a picture?(y/n)")
	if answer == "n" {
		fmt.Println(color.Bold+color.Red, "\n\tBackup Canceled!\n", color.Reset)
		return nil
	}

	// Convert to json
	finance, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}

	// Save a copy
	err = os.WriteFile(config.BackupFileWithDate, finance, 0644)
	if err != nil {
		return err
	}

	// Calculate old dept
	_, _, oldBalance := f.calculate()

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
	f.LoadFromFile(config.LocalFile)

	// Append old balance
	f.Add("dept", oldBalance)

	fmt.Println(color.Bold+color.Green, "\n\tBackup Done!\n", color.Reset)

	return nil
}

func (f *Finance) Undo() error {

	// Remove the last item
	f.Finance = f.Finance[:len(f.Finance)-1]

	config.LastAddedItemName = ""
	config.LastAddedItemSum = 0.0

	err := f.save(config.LocalFile)
	if err != nil {
		return err
	}

	err = f.save(config.BackupFile)
	if err != nil {
		return err
	}

	return nil
}

func (f *Finance) LoadFromFile(source string) error {

	file, err := os.Open(source)
	if err != nil {
		return err
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, f)
	if err != nil {
		return err
	}

	return nil
}

/* Other */

func (f *Finance) save(target string) error {

	finance, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(target, finance, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (f *Finance) calculate() (float64, float64, float64) {

	income := 0.0
	expenses := 0.0

	for _, item := range f.Finance {

		if item.VALUE < 0 {
			expenses += item.VALUE
		} else {
			income += item.VALUE
		}
	}

	return income, expenses, income + expenses
}
