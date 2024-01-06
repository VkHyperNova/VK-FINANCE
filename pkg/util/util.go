package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/VkHyperNova/VK-FINANCE/pkg/print"
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
		print.PrintCyan("\n" + question) 
		fmt.Scanln(&answer)
		

		if answer == "" {
			print.PrintRed("\nEnter a valid float!")
			goto start
		}

		floatValue, err := strconv.ParseFloat(answer, 64) 
		print.HandleError(err)                                  

		if err != nil {
			goto start 
		}

		return floatValue
}

func UserInputString(question string) string {
	// var answer string
	// print.PrintCyan("\n" + question) 
	// fmt.Scanln(&answer)

	// if answer == "" { 
	// 	answer = "No Comment" 
	// }

	// return answer

	// Write the string to a buffer
	print.PrintCyan(question)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	CommentString := scanner.Text()
	CommentString = strings.TrimSpace(CommentString)

	return CommentString

}







