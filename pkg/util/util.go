package util

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/VkHyperNova/VK-FINANCE/pkg/config"
	"github.com/peterh/liner"
)

/* Other Functions */

func PressAnyKey() {
	fmt.Print()
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

func Input() string {

	line := liner.NewLiner()
	defer line.Close()

	userInput, err := line.Prompt("=> ")
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
		fmt.Print("\n=> " + path)
	}
}

func WriteDataToFile(filename string, dataBytes []byte) {
	var err = os.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func ArrayContainsString(arr []string, name string) bool {
	for _, n := range arr {
		if n == name {
			return true
		}
	}
	return false
}

func Split(parts []string) (string, float64) {

	item := strings.ToLower(parts[0])

	// Try to convert the second part to a float
	sum, err := strconv.ParseFloat(parts[1], 64)

	if err != nil {
		fmt.Println(err)
	}

	// Adjust sum if item is in expenses category
	if ArrayContainsString(config.ExpensesItems, item) {
		sum = -sum
	}

	return item, sum
}