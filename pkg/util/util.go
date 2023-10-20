package util

import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/VkHyperNova/VK-FINANCE/pkg/print"
)



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

func InterfaceToByte(input interface{}) []byte {
	byteArray, err := json.MarshalIndent(input, "", " ")
	print.HandleError(err)                                     
	return byteArray                                     
}






