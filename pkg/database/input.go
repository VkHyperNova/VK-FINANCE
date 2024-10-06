package database

import (
	"strconv"
	"strings"
	"time"

	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

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

func (h *History) SplitInput(userInput string) bool {

	input := strings.TrimSpace(userInput)

	// Split the input into parts based on whitespace
	parts := strings.Fields(input)

	// Check if the input contains exactly two parts
	if len(parts) == 2 {

		// Assume the first part is the command
		item := strings.ToLower(parts[0])

		// Try to convert the second part to an integer
		sum, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			panic(err)
		}

		if util.ArrayContainsString(EXPENCES, item) {
			sum = -sum
		}

		h.Append(item, sum)

		return true
	}

	return false
}
