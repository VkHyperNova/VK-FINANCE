package database

import (
	"fmt"
	"strconv"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func (h *History) UserInput(cmd string) bool {
	comment := util.UserInputString("Comment: ")

	if comment == "q" {
		return false
	}

start:
	sum := util.UserInputString("Sum: ")
	if sum == "q" {
		return false
	}

	float, err := strconv.ParseFloat(sum, 64)

	if err != nil {
		util.PrintPurpleString("<< Enter a number! >>\n\n")
		goto start
	}

	if cmd == "s" || cmd == "spend" {
		h.Append(comment, -float)
	} else {
		h.Append(comment, float)
	}

	sumOfItem := h.FindItemValue(comment)
	util.PrintRedString(fmt.Sprintf("%.2f", sumOfItem) + " EUR")

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
