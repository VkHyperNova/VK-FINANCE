package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//go:generate goversioninfo

type finance struct {
	NET_WORTH float64 `json:"net_worth"`
	BALANCE   float64 `json:"balance"`
	EXPENSES  float64 `json:"expences"`
	DATE      string `json:"date"`
}

var BALANCE float64 = 0
var EXPENSES float64 = 0
var NET_WORTH float64 = 0
var DATE string 

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
	CHECK_DB()
	CL()
}

func CL() {

	PRINT_STATISTICS()
	
	fmt.Println(Cyan + "\n<< COMMANDS: add | exp | grow | q >>" + Reset)
	fmt.Print("=> ")

	reader := bufio.NewReader(os.Stdin)

	for {

		command := Convert_CRLF_To_LF(reader)

		switch command {
		case "add":
			Add()
		case "exp":
			EXP()
		case "grow":
			Grow()
		case "q":
			Quit("clear")
		default:
			Clear_Screen()
			CL()
		}
	}
}

/* Main Commands */

func Add() {
	BAL := Question("Money: ")
	BALANCE = BALANCE + BAL
	SAVE()
	Clear_Screen()
	CL()
}

func EXP() {
	EXP := Question("How much did you spend? ")
	EXPENSES = EXPENSES - EXP
	SAVE()
	Clear_Screen()
	CL()
}

/* UNDER CONSTRUCTION */
func Grow() {
	NET_WORTH = NET_WORTH + BALANCE
	BALANCE = 0
	SAVE()
	Clear_Screen()
	CL()
}

func PRINT_STATISTICS() {

	SETUP()

	fmt.Println()
	fmt.Println(Cyan + "<<___________ VK FINANCE v1 ___________>>" + Reset)
	fmt.Println(Cyan + "<<_____________________________________>>" + Reset)
	fmt.Println()
	fmt.Println(Cyan + "<< ", DATE, " >>" + Reset)
	fmt.Println()

	fmt.Print(Cyan + "NET WORTH: " + Reset)
	fmt.Println(Green, TWO_DECIMAL_POINTS(NET_WORTH), "EUR" + Reset)
	fmt.Println()

	fmt.Print(Cyan + "BALANCE: " + Reset)
	fmt.Println(Yellow, TWO_DECIMAL_POINTS(BALANCE), "EUR" + Reset)

	fmt.Print(Cyan + "EXPENSES: " + Reset)
	fmt.Println(Red, TWO_DECIMAL_POINTS(EXPENSES), "EUR" + Reset)

	DaylySpending()
	MONEY_SAVER()
}

func DaylySpending() {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstDayOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)
	DaysLeftBeforeEndOfMonth := lastDayOfMonth.Day() - now.Day()
	DaysLeftBeforePayday := DaysLeftBeforeEndOfMonth + 5
	DaysLeftBeforePayday = CheckWeekend(DaysLeftBeforePayday)
	DayMaxSpending := TWO_DECIMAL_POINTS(BALANCE / float64(DaysLeftBeforePayday))
	DaysLeftString := strconv.Itoa(DaysLeftBeforePayday)

	fmt.Println()
	fmt.Print(Cyan + "Day Max: (" + Reset)
	fmt.Print(Yellow + DaysLeftString + Reset)
	fmt.Print(Cyan + "): " + Reset)
	fmt.Println(Yellow + DayMaxSpending + " EUR" + Reset)

	WeekMaxSpending := TWO_DECIMAL_POINTS((BALANCE / float64(DaysLeftBeforePayday)) * 7)
	fmt.Print(Cyan + "Week Max: " + Reset)
	fmt.Println(Yellow + WeekMaxSpending + " EUR" + Reset)
	fmt.Println()
}

func MONEY_SAVER() {
	Savings := BALANCE * 0.25
	fmt.Print(Cyan + "SAVING (25%): " + Reset)
	fmt.Println(Green + TWO_DECIMAL_POINTS(Savings) + " EUR" + Reset)
}


/* Other */

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

func READ_FILE(filename string) []byte {
	file, err := os.ReadFile(filename)
	Error(err, "ReadFile")
	return file
}

func WRITE_FILE(filename string, dataBytes []byte) {

	var err = os.WriteFile(filename, dataBytes, 0644)
	Error(err, "WRITE_FILE FUNCTION")
}

func DIR_CHECK(dir_name string) bool {

	if _, err := os.Stat(dir_name); os.IsNotExist(err) {
		return false
	}
	return true
}

func MAKE_DIR(dir_name string) {
	_ = os.Mkdir(dir_name, 0700)
}

func CONVERT_TO_BYTE(data interface{}) []byte {
	dataBytes, err := json.MarshalIndent(data, "", "  ")
	Error(err, "Convert_To_Byte")

	return dataBytes
}

func SAVE() {
	data := CONCSTRUCT_FINANCE_JSON()
	dataBytes := CONVERT_TO_BYTE(data)
	WRITE_FILE("./finance.json", dataBytes)
}

func CONVERT_TO_FINANCE(body []byte) finance {

	data := finance{}

	err := json.Unmarshal(body, &data)
	Error(err, "CONVERT_TO_FINANCE")

	return data
}

func OPEN_DB() finance {
	data := READ_FILE("./finance.json")
	return CONVERT_TO_FINANCE(data)
}

func CHECK_DB() {
	if !DIR_CHECK("./finance.json") {
		NET_WORTH = Question("NET_WORTH? ")
		BALANCE = Question("BALANCE? ")
		SAVE()
		CL()
	}
}

func CheckWeekend(DaysLeftBeforePayday int) int {
	var AddDays = DaysLeftBeforePayday
	NextPayDayDate := time.Now().AddDate(0, 0, DaysLeftBeforePayday)
	GetWeekDay := NextPayDayDate.Weekday()

	if GetWeekDay == time.Saturday {
		AddDays += 2
	}

	if GetWeekDay == time.Sunday {
		AddDays += 1
	}

	return AddDays
}

func SETUP() {
	DB := OPEN_DB()
	NET_WORTH = DB.NET_WORTH
	BALANCE = DB.BALANCE
	EXPENSES = DB.EXPENSES
	DATE = DB.DATE
}

func CONCSTRUCT_FINANCE_JSON() finance {
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")

	return finance {
		NET_WORTH: NET_WORTH,
		BALANCE: BALANCE,
		EXPENSES: EXPENSES,
		DATE: timeString,
	}
}