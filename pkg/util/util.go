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
	if err != nil {
		PrintRed(err.Error() + "\n")
	}
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

// In this code, the `UserInputFloat64` function prompts the user with a question and returns a float64 value.
func UserInputFloat64(question string) float64 {
	start:
		var answer string
		PrintCyan("\n" + question) // It uses a loop to repeatedly prompt the user until a valid float64 value is entered.
		fmt.Scanln(&answer)

		if answer == "" { // If user presses enter
			answer = "0" // Assign zero
		}

		floatValue, err := strconv.ParseFloat(answer, 64) // Convert string to float64
		HandleError(err)                                  // The `HandleError` function is called to handle any errors that occur during the conversion process.

		if err != nil {
			goto start // If an error occurs, the loop restarts and the user is prompted again.
		}

		return floatValue
}

func UserInputString(question string) string {
	var answer string
	PrintCyan("\n" + question) // It uses a loop to repeatedly prompt the user until a valid float64 value is entered.
	fmt.Scanln(&answer)

	if answer == "" { // If user presses enter
		answer = "No Comment" // Assign zero
	}

	return answer
}

func FloatToString(number float64) string {
	return fmt.Sprintf("%.2f", number) // Converts float64 to string with 2 decimal points.
}

func InterfaceToByte(input interface{}) []byte {
	byteArray, err := json.MarshalIndent(input, "", " ") // MarshalIndent Converts The interface{} To A JSON Byte Array With Indentation.
	HandleError(err)                                     // Handle Error If Any.
	return byteArray                                     // Return JSON As []byte Array.
}
