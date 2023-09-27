package database

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/util"
	"os"
)

func DoesDirectoryExist(dir_name string) bool {

	// Get directory information
	if _, err := os.Stat(dir_name); os.IsNotExist(err) {
		return false
	}
	return true
}

func WriteDataToFile(filename string, dataBytes []byte) {
	// os.WriteFile writes data to a file named by filename
	// 0644 is the file mode
	var err = os.WriteFile(filename, dataBytes, 0644)
	// handleError checks if an error occurred
	util.HandleError(err)
}

func ReadFile(filename string) []byte {
	// ReadFile reads the file named by filename and returns the contents.
	file, err := os.ReadFile(filename)

	// HandleError checks if an error occurred and panics if it did.
	util.HandleError(err)

	// Return the contents of the file.
	return file
}
