package util

import (
	"log"
	"os"
)

/* Directory Functions */

func ValidateRequiredFiles() {
	if !DoesDirectoryExist("./history.json") {
		WriteDataToFile("./history.json", []byte("[]"))
	}

	if !DoesDirectoryExist("./history/history_json"){
		CreateDir("./history/history_json")
	}
}

func CreateDir(dirs string) {
	if err := os.MkdirAll(dirs, os.ModePerm); err != nil {
        log.Fatal(err)
    }

}

func RemoveFile(file string) {
	err := os.Remove(file)
	HandleError(err)
}

func DoesDirectoryExist(dir_name string) bool {
	if _, err := os.Stat(dir_name); os.IsNotExist(err) {
		return false
	}
	return true
}

func WriteDataToFile(filename string, dataBytes []byte) {
	var err = os.WriteFile(filename, dataBytes, 0644)
	HandleError(err)
}

func ReadFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	HandleError(err)
	return file
}
