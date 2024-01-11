package util

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

/* Other Functions */

func Contains(arr []string, name string) bool {
	for _, n := range arr {
		if n == name {
			return true
		}
	}
	return false
}

func UserInputFloat64(question string) float64 {
	start:
		var answer string
		PrintCyan("\n" + question) 
		fmt.Scanln(&answer)
		

		if answer == "" {
			PrintRed("\nEnter a valid float!")
			goto start
		}

		floatValue, err := strconv.ParseFloat(answer, 64) 
		HandleError(err)                                  

		if err != nil {
			goto start 
		}

		return floatValue
}

func UserInputString(question string) string {
	PrintCyan(question)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	CommentString := scanner.Text()
	CommentString = strings.TrimSpace(CommentString)

	return CommentString
}

func PressAnyKey() {
	PrintGray("\nPress Any Key To Continue...")
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







