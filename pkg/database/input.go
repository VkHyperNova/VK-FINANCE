package database

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func (h *History) GetUserInput(cmd string) bool {

	comment := util.GetString("Comment: ")

	if comment == "q" {
		return false
	}

s2:
	sum := util.GetString("Sum: ")

	if sum == "q" {
		return false
	}

	float, err := strconv.ParseFloat(sum, 64)

	if err != nil {
		util.PrintPurpleString("<< Enter a number! >>\n\n")
		goto s2
	}

	if cmd == "s" || cmd == "spend" {
		float = -float
	}

	h.Append(comment, float)

	// Print summary
	
	sumOfItem := 0.0
	for _, value := range h.History {
		if strings.EqualFold(value.COMMENT, comment) {
			sumOfItem += value.VALUE
		}
	}

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
