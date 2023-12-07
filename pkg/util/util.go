package util

import (
	"fmt"
	"strconv"
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
			answer = "0" 
		}

		floatValue, err := strconv.ParseFloat(answer, 64) 
		print.HandleError(err)                                  

		if err != nil {
			goto start 
		}

		return floatValue
}

func UserInputString(question string) string {
	var answer string
	print.PrintCyan("\n" + question) 
	fmt.Scanln(&answer)

	if answer == "" { 
		answer = "No Comment" 
	}

	return answer
}







