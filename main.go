package main

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/cmd"
	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

//go:generate goversioninfo

func main() {
	util.ValidateRequiredFiles()
	history := database.History{}
	history.Read()
	cmd.CommandLine(&history)
}
