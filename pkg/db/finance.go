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

type Item struct {
	DATE    string  `json:"date"`
	TIME    string  `json:"time"`
	COMMENT string  `json:"comment"`
	VALUE   float64 `json:"value"`
}

// Slice containing multiple HistoryItem instances.
type Finance struct {
	Finance []Item `json:"finance"`
}

/* Main */

func (f *Finance) Export() error {

	input, err := util.PromptWithSuggestion("Export db to d drive? (y/n) ", "n")
	if err != nil {
		return err
	}

	if input == "y" || input == "yes" {

		if err := util.InitBackupStorage(); err != nil {
			return err
		}

		if err := f.LoadFromFile(config.LocalFile); err != nil {
			return fmt.Errorf("load from file: %w", err)
		}

		finance, err := json.MarshalIndent(f, "", "  ")
		if err != nil {
			return err
		}

		if err := os.WriteFile(config.BackupFile, finance, 0644); err != nil {
			return err
		}

		fmt.Printf("Database exported to %s\nPress Enter!", config.BackupFile)
		return nil
	}

	fmt.Println("Export canceled!")
	return nil
}

func (f *Finance) Import() error {

	input, err := util.PromptWithSuggestion("Import db from d drive? (y/n) ", "n")
	if err != nil {
		return err
	}

	if input == "y" || input == "yes" {

		if err := util.InitBackupStorage(); err != nil {
			return err
		}

		if err := f.LoadFromFile(config.BackupFile); err != nil {
			return fmt.Errorf("load from file: %w", err)
		}

		finance, err := json.MarshalIndent(f, "", "  ")
		if err != nil {
			return err
		}

		if err := os.WriteFile(config.LocalFile, finance, 0644); err != nil {
			return err
		}

		fmt.Printf("Database imported from %s\nPress Enter!", config.BackupFile)
		return nil
	}

	fmt.Println("Import canceled!")

	return nil
}

func (f *Finance) Add(item string, value float64) error {

	comment := strings.ToLower(item)

	if !util.Contains(config.AllItems, comment) {
		return errors.New("No such Item!")
	}

	if !util.Contains(config.IncomeItems, comment) && comment != "dept" {
		value = -value
	}

	now := time.Now()

	NewItem := Item{
		DATE:    now.Format("02-01-2006"),
		TIME:    now.Format("15:04:05"),
		COMMENT: comment,
		VALUE:   value,
	}

	f.Finance = append(f.Finance, NewItem)

	return f.save()
}

func (f *Finance) Restart() error {

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

	if err := util.InitBackupStorage(); err != nil {
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
    if len(f.Finance) == 0 {
        return fmt.Errorf("nothing to undo")
    }
    f.Finance = f.Finance[:len(f.Finance)-1]
    config.LastAddedItemName = ""
    config.LastAddedItemSum = 0.0
    return f.save()
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

func (f *Finance) save() error {
    copySlice := make([]Item, len(f.Finance))
    copy(copySlice, f.Finance)
    copyFinance := Finance{Finance: copySlice}

    finance, err := json.MarshalIndent(copyFinance, "", "  ")
    if err != nil {
        return err
    }

    if err := os.WriteFile(config.LocalFile, finance, 0644); err != nil {
        return err
    }
    fmt.Println(color.Green + "Local save!" + color.Reset)

    if err := util.InitBackupStorage(); err != nil {
        fmt.Println(color.Yellow + "Backup init failed: " + err.Error() + color.Reset)
        return nil // or return err, depending on your needs
    }
    if err := os.WriteFile(config.BackupFile, finance, 0644); err != nil {
        fmt.Println(color.Yellow + "Backup write failed: " + err.Error() + color.Reset)
        return nil // same decision here
    }
    fmt.Println(color.Green + "Backup save!" + color.Reset)
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
