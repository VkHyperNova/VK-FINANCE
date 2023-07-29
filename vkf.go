package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)



const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
)

func main() {

	CL()
}

func CL() {

	PRINT_STATISTICS()

	fmt.Println(Cyan + "\n<< COMMANDS: add | q >>" + Reset)
	fmt.Print("=> ")

	reader := bufio.NewReader(os.Stdin)

	for {

		command := Convert_CRLF_To_LF(reader)

		switch command {
		case "add":
			Add()
		case "spent":
			Spent()
		case "q":
			Quit("clear")
		default:
			Clear_Screen()
			CL()
		}
	}
}

var MONEY float64 = 0
var SPENT float64 = 0

func Add() {

	MONEY = Question("Money: ")
	Clear_Screen()

	CL()
}

func Spent() {
	SPENT = Question("How much did you spent? ")
	Clear_Screen()

	CL()
}

func PRINT_STATISTICS() {

	fmt.Println()
	fmt.Println(Cyan + "<<___________ VK FINANCE v1 ___________>>" + Reset)
	fmt.Println()

	Days := CalculateDaysLeft()
	DayMaxSpending := TWO_DECIMAL_POINTS(MONEY/float64(Days))
	NET_WORTH := 1300
	fmt.Println("NET WORTH: ", NET_WORTH)
	fmt.Println("NET WORTH: ", NET_WORTH)
	fmt.Println("Avalable money: " + Yellow + TWO_DECIMAL_POINTS(MONEY) + Reset + " EUR")
	fmt.Println("Expences: ", SPENT)
	fmt.Println("Days left: ", Days)
	fmt.Println("You can spend daily: ", DayMaxSpending, " EUR")
}

func CalculateDaysLeft() int {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstDayOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)
	diff := lastDayOfMonth.Day() - now.Day()

	return diff + 6
}

func Question(question string) float64 {
start:
	var answer string
	fmt.Print("\n", question)
	fmt.Scanln(&answer)

	if answer == "" {
		answer = "0"
	}

	floatValue, err := strconv.ParseFloat(answer, 64)
	if err != nil {
		fmt.Println("<< Error:", err)
		goto start
	}

	return floatValue
}

func Clear_Screen() {

	if runtime.GOOS == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Convert_CRLF_To_LF(reader *bufio.Reader) string {

	// Read the answer
	input, _ := reader.ReadString('\n')

	// Convert CRLF to LF
	input = strings.Replace(input, "\r\n", "", -1) /* "\r\n" was before.  */

	return input
}

func TWO_DECIMAL_POINTS(number float64) string {
	return fmt.Sprintf("%.2f", number)
}

func Quit(clear string) {

	if clear == "clear" {
		Clear_Screen()
	}

	os.Exit(0)
}

func Error(err error, location string) {
	if err != nil {
		fmt.Println(" << Function name: ", location+" >>")
		fmt.Println(err.Error())

	}
}
