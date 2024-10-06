package util

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/peterh/liner"
)

/* Other Functions */

func GetString(question string) string {
	PrintCyanString(question)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	CommentString := scanner.Text()
	CommentString = strings.TrimSpace(CommentString)

	return CommentString
}

func PressAnyKey() {
	PrintGrayString("\n\nPress Any Key To Continue...")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

func ClearScreen() {
	if runtime.GOOS == "linux" { // check if the operating system is Linux
		cmd := exec.Command("clear") // execute the clear command
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") // execute the cls command
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func GetDayFromString(dateString string) time.Time {
	date, err := time.Parse("02-01-2006", dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}

	return date
}

func Input(prompt string) string {
	line := liner.NewLiner()
	defer line.Close()

	// Prompt the user with the given prompt string and read userInput
	userInput, err := line.Prompt(prompt)
	if err != nil {
		panic(err)
	}
	return userInput
}

func ValidateRequiredFiles() {

	folderName := "history"

	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		_ = os.Mkdir(folderName, 0700)
		fmt.Println("history folder created!")
	}

	path := "./history.json"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.WriteFile(path, []byte([]byte(`{"history": []}`)), 0644)
		if err != nil {
			panic(err)
		}
		PrintGrayString("\n=> " + path)
	}
}

func WriteDataToFile(filename string, dataBytes []byte) {
	var err = os.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		panic(err)
	}

	PrintGreenString("\n(" + filename + " saved!)")
}

func ArrayContainsString(arr []string, name string) bool {
	for _, n := range arr {
		if n == name {
			return true
		}
	}
	return false
}

