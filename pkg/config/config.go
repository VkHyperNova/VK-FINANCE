package config

// import "time"

/* Paths */

var DefaultPath = "./vk-finance.json"
var BackupPath = "/media/veikko/VK DATA/DATABASES/FINANCE/vk-finance.json"

// var HistoryPath = "/media/veikko/VK DATA/DATABASES/FINANCE/" + time.Now().AddDate(0, -1, 0).Format("January2006") + ".json"

/* Testing paths */

// var HistoryPath = "./history/"+ time.Now().AddDate(0, -1, 0).Format("January2006") + ".json"
// var BackupPath = "./history/history.json"

/* Constants */

var AllItems = []string{"pension", "programming", "wolt", "bolt", "bonus", "arved", "food", "catfood", "saun", "bensiin", "w", "other", "dept", "correction"}
var IncomeItems = []string{"pension", "programming", "wolt", "bolt", "bonus"}
var LastAddedItemName = ""
var LastAddedItemSum = 0.0

/* Colors */
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	Bold   = "\033[1m"
	Italic = "\033[3m"
)
