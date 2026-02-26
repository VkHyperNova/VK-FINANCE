package config

import (
	"path/filepath"
	"time"
)

/* Paths */
var DefaultContent = `{"finance": []}`
var	file = "finance.json"
var fileDate = time.Now().AddDate(0, -1, 0).Format("January2006") + ".json"
var BaseDB = "FINANCE"
var BaseLocal = "DATABASES"
var	BaseBackup = "/media/veikko/VK DATA/"

var LocalFile = filepath.Join(BaseLocal, BaseDB, file)
var BackupFile = filepath.Join(BaseBackup, BaseLocal, BaseDB, file)
var BackupFileWithDate = filepath.Join(BaseBackup, BaseLocal, BaseDB,fileDate)

/* Constants */

var AllItems = []string{"pension", "programming", "wolt", "bolt", "bonus", "arved", "food", "catfood", "saun", "bensiin", "w", "other", "dept", "correction"}
var IncomeItems = []string{"pension", "programming", "wolt", "bolt", "bonus"}
var LastAddedItemName = ""
var LastAddedItemSum = 0.0



