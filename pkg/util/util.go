package util

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)


func HandleError(err error) {
	// check if error is not nil
	if err != nil {
		// print error message in red color
		PrintRed(err.Error() + "\n")
	}
}
func ClearScreen() {

	// check if the operating system is Linux
	if runtime.GOOS == "linux" {
		// execute the clear command
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "windows" {
		// execute the cls command
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func GetUserInput(question string) float64 {
start:
	var answer string
	PrintCyan("\n" + question)
	fmt.Scanln(&answer)

	if answer == "" {
		answer = "0"
	}

	floatValue, err := strconv.ParseFloat(answer, 64)
	HandleError(err)

	if err != nil {
		goto start
	}

	return floatValue
}

func Float64ToStringWithTwoDecimalPoints(number float64) string {
	// Use fmt.Sprintf to format the float64 number with two decimal points
	return fmt.Sprintf("%.2f", number)
}

func InterfaceToByteArray(data interface{}) []byte {
	// MarshalIndent converts the interface{} to a JSON byte array with indentation.
	dataBytes, err := json.MarshalIndent(data, "", " ")
	// handleError checks for errors and panics if there is one.
	HandleError(err)

	return dataBytes
}




