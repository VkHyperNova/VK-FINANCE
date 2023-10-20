package dir

import (
	"github.com/VkHyperNova/VK-FINANCE/pkg/print"
	"os"
)

func ValidateRequiredFiles() {

	if !DoesDirectoryExist("./finance.json") {
		WriteDataToFile("./finance.json", []byte("[]"))
	}
	

	if !DoesDirectoryExist("./history.json") {
		WriteDataToFile("./history.json", []byte("[]"))
	}
}

func DoesDirectoryExist(dir_name string) bool {
	if _, err := os.Stat(dir_name); os.IsNotExist(err) {
		return false
	}
	return true
}

func WriteDataToFile(filename string, dataBytes []byte) {
	var err = os.WriteFile(filename, dataBytes, 0644)
	print.HandleError(err)
}

func ReadFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	print.HandleError(err)
	return file
}
