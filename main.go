package main

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/cmd"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

//go:generate goversioninfo

func main() {
	util.ClearScreen()
	util.ValidateRequiredFiles()
	cmd.CMD()
}
