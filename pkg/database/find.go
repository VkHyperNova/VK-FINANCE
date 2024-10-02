package database

import "strings"

func (h *History) FindItemValue(comment string) float64 {

	sumOfItem := 0.0

	for _, value := range h.History {
		if strings.EqualFold(value.COMMENT, comment) {
			sumOfItem += value.VALUE
		}
	}
	return sumOfItem
}
