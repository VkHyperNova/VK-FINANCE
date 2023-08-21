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
	NET_WORTH float64    `json:"net_worth"`
	BALANCE   float64    `json:"balance"`
	EXPENSES  float64    `json:"expences"`
	MONTH     time.Month `json:"month"`
}

/* DATABASE */
var NET_WORTH float64 = 0
var BALANCE float64 = 0
var EXPENSES float64 = 0

/* CURRENT TIME */
var CURRENT time.Time
var CURRENT_MONTH time.Month
var DAYS_UNTIL_SUNDAY time.Weekday

/* CALCULATIONS */
var DAYSLEFT int = 0
var DAYMAX float64
var WEEKMAX float64
var SAVINGS float64
var SAVER_DAYMAX float64
var SAVER_WEEKMAX float64

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
	CLEAR_SCREEN()
	CL()
}

/* MAIN CMD */
func CL() {
	CHECK_DB()
	GET_DATA()
	CALCULATE()
	PRINT()

	fmt.Println(Cyan + "\n<< COMMANDS: add | exp | grow | q >>" + Reset)
	fmt.Print("=> ")

	reader := bufio.NewReader(os.Stdin)

	for {

		command := CONVERT_CRLF_TO_LF(reader)

		switch command {
		case "add":
			ADD()
		case "exp":
			EXP()
		case "grow":
			GROW()
		case "q":
			QUIT("clear")
		default:
			CLEAR_SCREEN()
			CL()
		}
	}
}

func ADD() {
	BAL := PROMPT("Money: ")
	BALANCE = BALANCE + BAL
	SAVE_DB()
	CLEAR_SCREEN()
	CL()
}

func EXP() {
	EXP := PROMPT("How much did you spend? ")
	BALANCE = BALANCE - EXP
	EXPENSES = EXPENSES - EXP
	SAVE_DB()
	CLEAR_SCREEN()
	CL()
}

func GROW() {
	NET_WORTH = NET_WORTH + BALANCE
	BALANCE = 0
	EXPENSES = 0
	SAVE_DB()
	CLEAR_SCREEN()
	CL()
}

/* CALCULATORS */
func CALCULATE() {
	CALCULATE_DAYSLEFT()
	CALCULATE_DAYMAX()
	CALCULATE_WEEKMAX()
	CALCULATE_SAVER_DAYMAX()
	CALCULATE_SAVER_WEEKMAX()
	CALCULATE_SAVINGS()
}

func CALCULATE_DAYSLEFT() {
	Year, Month, _ := CURRENT.Date()
	Location := CURRENT.Location()
	FIRST_DAY_OF_MONTH := time.Date(Year, Month, 1, 0, 0, 0, 0, Location)
	LAST_DAY_OF_MONTH := FIRST_DAY_OF_MONTH.AddDate(0, 1, -1)

	DAYSLEFT = CHECK_WEEKEND((LAST_DAY_OF_MONTH.Day() - CURRENT.Day()) + 5)
}

func CALCULATE_DAYMAX() {
	DAYMAX = BALANCE / float64(DAYSLEFT)
}

func CALCULATE_WEEKMAX() {
	DAYS_UNTIL_SUNDAY = time.Sunday - CURRENT.Weekday()

	if CURRENT.Weekday() == time.Sunday {
		DAYS_UNTIL_SUNDAY += 7
	} else {
		DAYS_UNTIL_SUNDAY += 8
	}

	WEEKMAX = DAYMAX * float64(DAYS_UNTIL_SUNDAY)
}

func CALCULATE_SAVINGS() {
	SAVINGS = BALANCE * 0.25 /* 25% SAVING */
}

func CALCULATE_SAVER_DAYMAX() {
	SAVER_DAYMAX = DAYMAX - (DAYMAX * 0.25)
}

func CALCULATE_SAVER_WEEKMAX() {
	SAVER_WEEKMAX = WEEKMAX - ((DAYMAX * float64(DAYS_UNTIL_SUNDAY)) * 0.25)
}

/* PRINTS */
func PRINT() {
	PRINT_PROGRAM_INFO()
	PRINT_MONTH()
	PRINT_NET_WORTH()
	PRINT_BALANCE()
	PRINT_EXPENCES()
	PRINT_DAYMAX()
	PRINT_WEEKMAX()
	PRINT_SAVER_DAYMAX()
	PRINT_SAVER_WEEKMAX()
	PRINT_SAVINGS()
}

func PRINT_PROGRAM_INFO() {
	fmt.Print("\n" + Cyan + "<<___________ VK FINANCE v1 ___________>>\n")
}

func PRINT_MONTH() {
	fmt.Print(Yellow+"\n", CURRENT_MONTH, Reset+"\n")
}

func PRINT_NET_WORTH() {
	fmt.Print("\n" + Cyan + "NET WORTH: " + Reset + Green + CONVERT_TO_TWO_DECIMAL_POINTS_STRING(NET_WORTH) + " EUR" + Reset + "\n\n")
}

func PRINT_BALANCE() {
	fmt.Println(Cyan + "BALANCE: " + Reset + Yellow + CONVERT_TO_TWO_DECIMAL_POINTS_STRING(BALANCE) + " EUR" + Reset)
}

func PRINT_EXPENCES() {
	fmt.Println(Cyan + "EXPENSES: " + Reset + Red + CONVERT_TO_TWO_DECIMAL_POINTS_STRING(EXPENSES) + " EUR" + Reset)
}

func PRINT_DAYMAX() {
	fmt.Print("\n"+Cyan+"Day Max: (", DAYSLEFT, " Days): "+Reset+Yellow+CONVERT_TO_TWO_DECIMAL_POINTS_STRING(DAYMAX)+" EUR"+Reset+"\n")
}

func PRINT_WEEKMAX() {
	fmt.Print(Cyan+"Week Max: (", int(DAYS_UNTIL_SUNDAY), " Days): "+Reset+Yellow+CONVERT_TO_TWO_DECIMAL_POINTS_STRING(WEEKMAX)+" EUR"+Reset)
}

func PRINT_SAVINGS() {
	fmt.Print("\n" + Cyan + "SAVING (25%): " + Reset + Green + CONVERT_TO_TWO_DECIMAL_POINTS_STRING(SAVINGS) + " EUR" + Reset + "\n")
}

func PRINT_SAVER_DAYMAX() {
	fmt.Print("\n\n" + Cyan + "Day Max (25%): " + Reset + Green + CONVERT_TO_TWO_DECIMAL_POINTS_STRING(SAVER_DAYMAX) + " EUR" + Reset + "\n")
}

func PRINT_SAVER_WEEKMAX() {
	fmt.Print(Cyan + "Week Max (25%): " + Reset + Green + CONVERT_TO_TWO_DECIMAL_POINTS_STRING(SAVER_WEEKMAX) + " EUR" + Reset + "\n")
}

/* DATABASE */
func CONCSTRUCT_FINANCE_JSON() finance {

	return finance{
		NET_WORTH: NET_WORTH,
		BALANCE:   BALANCE,
		EXPENSES:  EXPENSES,
		MONTH:     CURRENT_MONTH,
	}
}

func CONVERT_TO_FINANCE(body []byte) finance {

	data := finance{}

	err := json.Unmarshal(body, &data)
	ERROR(err, "CONVERT_TO_FINANCE")

	return data
}

func OPEN_DB() finance {
	data := READ_FILE("./finance.json")
	return CONVERT_TO_FINANCE(data)
}

func CREATE_DB() {
	NET_WORTH = PROMPT("NET_WORTH: ")
	CLEAR_SCREEN()
	SAVE_DB()
}

func SAVE_DB() {
	data := CONCSTRUCT_FINANCE_JSON()
	dataBytes := CONVERT_TO_BYTE(data)
	WRITE_FILE("./finance.json", dataBytes)
}

func GET_DATA() {
	DB := OPEN_DB()
	NET_WORTH = DB.NET_WORTH
	BALANCE = DB.BALANCE
	EXPENSES = DB.EXPENSES
	CURRENT = time.Now()
	CURRENT_MONTH = CURRENT.Month()
}

/* CHECKERS */
func CHECK_DB() {
	if !DIR_CHECK("./finance.json") {
		CREATE_DB()
	}
}

func CHECK_WEEKEND(DaysLeftBeforePayday int) int {
	AddDays := DaysLeftBeforePayday
	NextPayDayDate := time.Now().AddDate(0, 0, DaysLeftBeforePayday)

	if NextPayDayDate.Weekday() == time.Saturday {
		AddDays += 2
	}

	if NextPayDayDate.Weekday() == time.Sunday {
		AddDays += 1
	}

	return AddDays
}

/* DIR */
func MAKE_DIR(dir_name string) {
	_ = os.Mkdir(dir_name, 0700)
}

func READ_FILE(filename string) []byte {
	file, err := os.ReadFile(filename)
	ERROR(err, "ReadFile")
	return file
}

func WRITE_FILE(filename string, dataBytes []byte) {

	var err = os.WriteFile(filename, dataBytes, 0644)
	ERROR(err, "WRITE_FILE FUNCTION")
}

func DIR_CHECK(dir_name string) bool {

	if _, err := os.Stat(dir_name); os.IsNotExist(err) {
		return false
	}
	return true
}

/* CONVERTERS */
func CONVERT_TO_BYTE(data interface{}) []byte {
	dataBytes, err := json.MarshalIndent(data, "", "  ")
	ERROR(err, "Convert_To_Byte")

	return dataBytes
}

func CONVERT_TO_TWO_DECIMAL_POINTS_STRING(number float64) string {
	return fmt.Sprintf("%.2f", number)
}

func CONVERT_CRLF_TO_LF(reader *bufio.Reader) string {

	// Read the answer
	input, _ := reader.ReadString('\n')

	// Convert CRLF to LF
	input = strings.Replace(input, "\r\n", "", -1) /* "\r\n" was before.  */

	return input
}

/* Other */

func PROMPT(question string) float64 {
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

func CLEAR_SCREEN() {

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

func QUIT(clear string) {

	if clear == "clear" {
		CLEAR_SCREEN()
	}

	os.Exit(0)
}

func ERROR(err error, location string) {
	if err != nil {
		fmt.Println(" << Function name: ", location+" >>")
		fmt.Println(err.Error())

	}
}