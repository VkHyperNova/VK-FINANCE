package config

import "time"

/* Paths */
var Path = "./history.json"
// var HistoryPath = "/media/veikko/VK DATA/DATABASES/FINANCE/" + time.Now().AddDate(0, -1, 0).Format("January2006") + ".json"
// var BackupPath = "/media/veikko/VK DATA/DATABASES/FINANCE/history.json"

var HistoryPath = "./history/"+ time.Now().AddDate(0, -1, 0).Format("January2006") + ".json"
var BackupPath = "./history/history.json"



/* Constants */
var IncomeItems = []string{"pension", "programming", "wolt", "bolt", "muu"}
var ExpensesItems = []string{"arved", "food", "catfood", "saun", "bensiin", "vape", "w", "other", "dept", "correction"}

/* Variables */
var DEPT float64
