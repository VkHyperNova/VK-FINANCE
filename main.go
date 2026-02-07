package main

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/cmd"
	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

//go:generate goversioninfo

func main() {
	if util.IsVKDataMounted() {
		util.CreateDatabaseFile()
		db := database.Finance{}
		db.ReadFile()
		cmd.CommandLine(&db)
	}
}
