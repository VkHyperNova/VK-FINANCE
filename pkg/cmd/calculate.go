package cmd

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func Calculate() {
	DayBudget()
	WeekBudget()
	Budget()
}

func DayBudget() {
	util.DayBudget = (util.INCOME - util.PERFECT_SAVE) / 31
	util.DayBudgetSpent = util.EXPENSES / 31
}

func WeekBudget() {
	util.WeekBudget = ((util.INCOME - util.PERFECT_SAVE) / 31) * 7
	util.WeekBudgetSpent = (util.EXPENSES / 31) * 7
}

func Budget() {
	util.Budget = util.BALANCE - util.PERFECT_SAVE
}
