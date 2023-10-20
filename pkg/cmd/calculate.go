package cmd

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/global"
)

func DayBudget() {
	global.DayBudget = (global.INCOME - global.SAVING) / 31
	global.DayBudgetSpent = global.EXPENSES / 31
}

func WeekBudget() {
	global.WeekBudget = ((global.INCOME - global.SAVING) / 31) * 7
	global.WeekBudgetSpent = (global.EXPENSES / 31) * 7
}

func Budget() {
	global.Budget = global.BALANCE - global.SAVING
}
