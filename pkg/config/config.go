package config

import "time"

/* Paths */
var Path = "./history.json"
var BackupPath = "/media/veikko/VK DATA/DATABASES/FINANCE/history.json"
var HistoryPath = "/media/veikko/VK DATA/DATABASES/FINANCE/" + time.Now().AddDate(0, -1, 0).Format("January2006") + ".json"

/* Constants */
var IncomeItems = []string{"pension", "programming", "wolt", "bolt", "muu"}
var ExpencesItems = []string{"arved", "food", "catfood", "saun", "bensiin", "vape", "w", "other", "old balance", "correction"}

/* Variables */
var OLDBALANCE float64
var INCOME float64
