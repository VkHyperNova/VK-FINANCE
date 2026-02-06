package main

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/cmd"
	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

//go:generate goversioninfo

func main() {
	if util.IsVKDataMounted() {
		util.ValidateFiles()
		db := database.History{}
		db.Read()
		cmd.CommandLine(&db)
	}
}
