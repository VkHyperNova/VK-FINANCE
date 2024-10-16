package config

import "time"

/* Paths */

var Path = "./history.json"
var HistoryPath = "/media/veikko/VK DATA/DATABASES/FINANCE/" + time.Now().AddDate(0, -1, 0).Format("January2006") + ".json"
var BackupPath = "/media/veikko/VK DATA/DATABASES/FINANCE/history.json"

/* Testing paths */

// var HistoryPath = "./history/"+ time.Now().AddDate(0, -1, 0).Format("January2006") + ".json"
// var BackupPath = "./history/history.json"

/* Constants */

var IncomeItems = []string{"pension", "programming", "wolt", "bolt", "muu"}
var ExpensesItems = []string{"arved", "food", "catfood", "saun", "bensiin", "vape", "w", "other", "dept", "correction"}


/* Database Functions */

type HistoryItem struct {
	DATE    string  `json:"date"`
	TIME    string  `json:"time"`
	COMMENT string  `json:"comment"`
	VALUE   float64 `json:"value"`
}

// Slice containing multiple HistoryItem instances.

type History struct {
	History []HistoryItem `json:"history"`
}
