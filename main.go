package main

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/cmd"
)

//go:generate goversioninfo

func main() {
	cmd.DisplayAndHandleOptions()
}
