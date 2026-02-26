package main

import (
	"fmt"
	"log"
	"os"

	"github.com/VkHyperNova/VK-FINANCE/pkg/cmd"
	"github.com/VkHyperNova/VK-FINANCE/pkg/database"
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
)

func main() {

	if err := util.CreateFilesAndFolders(); err != nil {
		fmt.Println("Error creating files/folders:", err)
		os.Exit(1)
	}

	db := database.Finance{}

	err := db.ReadFromFile()
	if err != nil {
		log.Fatalf("Fatal error: failed to load fastings database: %v", err)
	}
	
	cmd.CommandLine(&db)
}
