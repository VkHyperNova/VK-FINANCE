package cmd

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func Calculate() {
	calculateDaylySpendingAmount()
	calculateWeeklySpendingAmount()
	calculateMoneyLeft()
}

func calculateDaylySpendingAmount() {
	// Calculate maximum savings budget per day
	util.MaxSavingsBudgetDay = (util.INCOME - util.PERFECT_SAVE) / 31
	// Calculate maximum spendable amount per day
	util.MaxSpendableAmountDay = util.EXPENSES / 31
}

func calculateWeeklySpendingAmount() {
	// Max savings budget calculation
	util.MaxSavingsBudgetWeek = ((util.INCOME - util.PERFECT_SAVE) / 31) * 7
	// Max spendable amount calculation
	util.MaxSpendableAmountWeek = (util.EXPENSES / 31) * 7
}

func calculateMoneyLeft() {
	// Calculate the money left after saving for a perfect save
	util.MONEY = util.BALANCE - util.PERFECT_SAVE
}




