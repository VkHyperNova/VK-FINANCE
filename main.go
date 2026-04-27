package main

import (
	"log"

	"github.com/VkHyperNova/VK-FINANCE/pkg/cmd"
	"github.com/VkHyperNova/VK-FINANCE/pkg/config"
	"github.com/VkHyperNova/VK-FINANCE/pkg/db"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func main() {

	mounted, err := util.InitStorage()
	if err != nil {
		log.Fatalf("init storage: %v", err)
	}
	
	if mounted {
		defer util.UnmountDrive()
	}

	f := &db.Finance{}

	if err := f.LoadFromFile(config.LocalFile); err != nil {
		log.Fatalf("load from file: %v", err)
	}

	cmd.Start(f)
}
