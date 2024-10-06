package database

import "time"

/* Database Functions */

type HistoryItem struct {
	DATE    string  `json:"date"`
	TIME    string  `json:"time"`
	COMMENT string  `json:"comment"`
	VALUE   float64 `json:"value"`
}

type History struct {
	History []HistoryItem `json:"history"` // Slice containing multiple Quote instances.
}

var Path = "./history.json"
var BackupPath = "/media/veikko/VK DATA/DATABASES/FINANCE/history.json"
var HistoryPath = "/media/veikko/VK DATA/DATABASES/FINANCE/" + time.Now().AddDate(0, -1, 0).Format("January2006") + ".json"

var INCOME = []string{"pension", "wolt", "bolt", "muu"}
var EXPENCES = []string{"arved", "food", "catfood", "saun", "bensiin", "vape", "w", "other", "oldbalance", "correction"}
var OLDBALANCE float64
