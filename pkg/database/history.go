package database

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

var INCOMESOURCES = []string{"pension", "sotsiaal", "wolt", "bolt", "muu"}
var MAINEXPENCES = []string{"arved", "food", "saun", "bensiin", "e-smoke", "weed", "other", "oldbalance", "correction"}

var OLDBALANCE float64


