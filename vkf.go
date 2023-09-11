package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

//go:generate goversioninfo

func main() {
	CLEAR_SCREEN()
	CL()
}

/* MAIN CMD */
func CL() {
	CHECK_DB()
	SETUP()
	PRINT_ALL()

	PRINT_SEPARATOR_TWO()
	PRINT_CYAN("Program Options: \n\n")
	PRINT_COMMAND("add")
	PRINT_COMMAND("bills")
	PRINT_COMMAND("gas")
	PRINT_COMMAND("food")
	PRINT_COMMAND("other")
	PRINT_COMMAND("grow")
	PRINT_COMMAND("reset")
	PRINT_COMMAND("history")
	PRINT_COMMAND("q")


	PRINT_GRAY("\n\n=> ")

	reader := bufio.NewReader(os.Stdin)

	for {

		command := CONVERT_CRLF_TO_LF(reader)

		switch command {
		case "1", "add":
			ADD()
		case "2", "bills":
			BILLS_EXP()
		case "3", "gas":
			GAS_EXP()
		case "4", "food":
			FOOD_EXP()
		case "5", "other":
			OTHER_EXP()
		case "6", "grow":
			GROW()
		case "7", "reset":
			RESET()
		case "8", "history":
			PRINT_HISTORY()
		case "q":
			QUIT("clear")
		default:
			CLEAR_SCREEN()
			CL()
		}
	}
}


func ADD() {
	ADD := PROMPT("Add Money: ")
	INCOME = INCOME + ADD
	BALANCE = BALANCE + ADD
	SAVE_DB()
	SAVE_HISTORY("ADD", ADD)
	CLEAR_SCREEN()
	CL()
}

func BILLS_EXP() {
	EXP := PROMPT("Bills expenses: ")
	BILLS = BILLS + EXP
	BALANCE = BALANCE - EXP
	EXPENSES = EXPENSES - EXP
	SAVE_DB()
	SAVE_HISTORY("BILLS EXPENSES", EXP)
	CLEAR_SCREEN()
	CL()
}

func GAS_EXP() {
	EXP := PROMPT("Gas expenses: ")
	GAS = GAS + EXP
	BALANCE = BALANCE - EXP
	EXPENSES = EXPENSES - EXP
	SAVE_DB()
	SAVE_HISTORY("GAS EXPENSES", EXP)
	CLEAR_SCREEN()
	CL()
}

func FOOD_EXP() {
	EXP := PROMPT("Food expenses: ")
	FOOD = FOOD + EXP
	BALANCE = BALANCE - EXP
	EXPENSES = EXPENSES - EXP
	SAVE_DB()
	SAVE_HISTORY("FOOD EXPENSES", EXP)
	CLEAR_SCREEN()
	CL()
}

func OTHER_EXP() {
	EXP := PROMPT("Other expenses: ")
	OTHER = OTHER + EXP
	BALANCE = BALANCE - EXP
	EXPENSES = EXPENSES - EXP
	SAVE_DB()
	SAVE_HISTORY("OTHER EXPENSES", EXP)
	CLEAR_SCREEN()
	CL()
}

func GROW() {
	NET_WORTH = NET_WORTH + BALANCE
	BALANCE = 0
	EXPENSES = 0
	SAVE_DB()
	SAVE_HISTORY("GROW", BALANCE)
	CLEAR_SCREEN()
	CL()
}

func RESET() {
	INCOME = 0
	EXPENSES = 0
	SAVE_DB()
	SAVE_HISTORY("RESET", EXPENSES)
	REMOVE_FILE("./history.json")
	CLEAR_SCREEN()
	CL()
}

func PRINT_HISTORY() {
	now := time.Now()
	formattedDate := now.Format("02-01-2006")

	file := READ_FILE("./history.json")
	hdata := CONVERT_TO_HISTORY(file)

	CLEAR_SCREEN()

	PRINT_CYAN("History: \n")

	for _, value := range hdata {
		if strings.Contains(value.DATE, formattedDate) {
			fmt.Println(value)
		}
	}

	CL()
}

/* MAIN FINANCE VARIABLES */
var DATABASE finance
var NET_WORTH float64
var BALANCE float64
var EXPENSES float64
var BILLS float64
var GAS float64
var FOOD float64
var OTHER float64
var INCOME float64
var PERFECT_SAVE float64

func SETUP() {
	data := READ_FILE("./finance.json")
	DATABASE = CONVERT_TO_FINANCE(data)
	NET_WORTH = DATABASE.NET_WORTH
	BALANCE = DATABASE.BALANCE
	EXPENSES = DATABASE.EXPENSES
	BILLS = DATABASE.BILLS
	GAS = DATABASE.GAS
	FOOD = DATABASE.FOOD
	OTHER = DATABASE.OTHER
	INCOME = DATABASE.INCOME
	PERFECT_SAVE = INCOME * 0.25
}

func PRINT_PROGRAM_INFO() {
	
	PRINT_SEPARATOR_ONE()
	
	PRINT_GRAY("============== VK FINANCE v1 ===============\n")
	
	PRINT_SEPARATOR_ONE()
	PRINT_GRAY(DATABASE.MONTH.String() + "\n")
}

func PRINT_COMMAND(name string) {
	PRINT_CYAN("[")
	PRINT_YELLOW(name)
	PRINT_CYAN("] ")
}

func PRINT_SEPARATOR_ONE() {
	PRINT_GRAY("============================================\n")
}

func PRINT_SEPARATOR_TWO() {
	PRINT_GRAY("--------------------------------------------\n")
}

func PRINT_NET_WORTH() {
	PRINT_CYAN("NET WORTH: ")
	PRINT_GREEN(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(NET_WORTH) + " EUR\n\n")
}

func PRINT_INCOME() {
	PRINT_CYAN("INCOME: ")
	PRINT_GREEN("+" + CONVERT_TO_TWO_DECIMAL_POINTS_STRING(INCOME) + " EUR\n")
}

func PRINT_EXPENCES() {
	PRINT_CYAN("EXPENCES: ")
	PRINT_RED(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(EXPENSES) + " EUR\n\n")

	PRINT_CYAN("-> Bills: ")
	PRINT_RED(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(BILLS) + " EUR\n")
	PRINT_CYAN("-> Gas: ")
	PRINT_RED(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(GAS) + " EUR\n")
	PRINT_CYAN("-> Food: ")
	PRINT_RED(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(FOOD) + " EUR\n")
	PRINT_CYAN("-> Other: ")
	PRINT_RED(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(OTHER) + " EUR\n\n")
}

func PRINT_BALANCE() {
	PRINT_CYAN("BALANCE: ")
	PRINT_YELLOW(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(BALANCE) + " EUR\n")
}

func PRINT_DAY() {
	PERFECT_DAY := (INCOME - PERFECT_SAVE) / 31
	REAL_DAY := EXPENSES / 31
	PRINT_CYAN("ESTIMATED DAY: ")
	PRINT_GREEN(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(PERFECT_DAY) + " EUR")
	PRINT_CYAN(" | ")
	PRINT_RED(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(REAL_DAY) + " EUR\n")
}

func PRINT_WEEK() {
	PERFECT_WEEK := ((INCOME - PERFECT_SAVE) / 31) * 7
	REAL_WEEK := (EXPENSES / 31) * 7
	PRINT_CYAN("ESTIMATED WEEK: ")
	PRINT_GREEN(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(PERFECT_WEEK) + " EUR")
	PRINT_CYAN(" | ")
	PRINT_RED(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(REAL_WEEK) + " EUR\n")
}

func PRINT_SAVING() {
	PRINT_CYAN("SAVING (25%): ")
	PRINT_GREEN(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(PERFECT_SAVE) + " EUR\n\n")
}

func PRINT_MONEY() {
	MONEY := BALANCE - PERFECT_SAVE
	PRINT_CYAN("MONEY: ")
	if MONEY < 0 {
		PRINT_RED(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(MONEY) + " EUR\n\n")
	} else {
		PRINT_GREEN(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(MONEY) + " EUR\n\n")
	}
}

func PRINT_ALL() {
	PRINT_PROGRAM_INFO()

	PRINT_NET_WORTH()
	PRINT_INCOME()
	PRINT_EXPENCES()

	PRINT_DAY()
	PRINT_WEEK()
	PRINT_SAVING()

	PRINT_BALANCE()
	PRINT_MONEY()
}

/* COLORS */
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

func PRINT_RED(a string) {
	fmt.Print(Red + a + Reset)
}

func PRINT_GREEN(a string) {
	fmt.Print(Green + a + Reset)
}

func PRINT_YELLOW(a string) {
	fmt.Print(Yellow + a + Reset)
}

func PRINT_BLUE(a string) {
	fmt.Print(Blue + a + Reset)
}

func PRINT_PURPLE(a string) {
	fmt.Print(Purple + a + Reset)
}

func PRINT_CYAN(a string) {
	fmt.Print(Cyan + a + Reset)
}

func PRINT_GRAY(a string) {
	fmt.Print(Gray + a + Reset)
}

/* DATABASE */
type finance struct {
	NET_WORTH float64    `json:"net_worth"`
	INCOME    float64    `json:"income"`
	BALANCE   float64    `json:"balance"`
	EXPENSES  float64    `json:"expences"`
	BILLS     float64    `json:"bills"`
	GAS       float64    `json:"gas"`
	FOOD      float64    `json:"food"`
	OTHER     float64    `json:"other"`
	MONTH     time.Month `json:"month"`
}

func CONCSTRUCT_FINANCE_JSON() finance {

	now := time.Now()

	return finance{
		NET_WORTH: math.Round(NET_WORTH*100) / 100,
		INCOME:    math.Round(INCOME*100) / 100,
		BALANCE:   math.Round(BALANCE*100) / 100,
		EXPENSES:  math.Round(EXPENSES*100) / 100,
		BILLS:     math.Round(BILLS*100) / 100,
		GAS:       math.Round(GAS*100) / 100,
		FOOD:      math.Round(FOOD*100) / 100,
		OTHER:     math.Round(OTHER*100) / 100,
		MONTH:     now.Month(),
	}
}

type history struct {
	DATE   string  `json:"date"`
	TIME   string  `json:"time"`
	ACTION string  `json:"action"`
	VALUE  float64 `json:"value"`
}

func CONCSTRUCT_HISTORY_JSON(LAST_ACTION string, VALUE float64) history {

	now := time.Now()
	formattedTime := now.Format("15:04:05")
	formattedDate := now.Format("02-01-2006")

	return history{
		DATE:   formattedDate,
		TIME:   formattedTime,
		ACTION: LAST_ACTION,
		VALUE:  VALUE,
	}
}

func CONVERT_TO_HISTORY(body []byte) []history {

	data := []history{}

	err := json.Unmarshal(body, &data)
	ERROR(err, "CONVERT_TO_FINANCE")

	return data
}

func SAVE_HISTORY(LAST_ACTION string, VALUE float64) {

	file := READ_FILE("./history.json")
	hdata := CONVERT_TO_HISTORY(file)

	newdata := CONCSTRUCT_HISTORY_JSON(LAST_ACTION, VALUE)
	hdata = append(hdata, newdata)
	dataBytes := CONVERT_TO_BYTE(hdata)
	WRITE_FILE("./history.json", dataBytes)
}

func CONVERT_TO_FINANCE(body []byte) finance {

	data := finance{}

	err := json.Unmarshal(body, &data)
	ERROR(err, "CONVERT_TO_FINANCE")

	return data
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

func CREATE_HISTROY_DB() {
	WRITE_FILE("./history.json", []byte("[]"))
}

func CHECK_DB() {
	if !DIR_CHECK("./finance.json") {
		CREATE_DB()
	}

	if !DIR_CHECK("./history.json") {
		CREATE_HISTROY_DB()
	}
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
	PRINT_CYAN("\n" + question)
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

func REMOVE_FILE(file string) {

	err := os.Remove(file)
	if err != nil {
		fmt.Println(err)
	}
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
		PRINT_RED(" << Function name: ")
		PRINT_RED(location + " >>\n")
		PRINT_RED(err.Error() + "\n")

	}
}

/**************************************************************************************************************************************************/

/* mby useful*/
// func CALCULATE_DAYSLEFT() {
// 	Year, Month, _ := CURRENT.Date()
// 	Location := CURRENT.Location()
// 	FIRST_DAY_OF_MONTH := time.Date(Year, Month, 1, 0, 0, 0, 0, Location)
// 	LAST_DAY_OF_MONTH := FIRST_DAY_OF_MONTH.AddDate(0, 1, -1)

// 	DAYSLEFT = CHECK_WEEKEND((LAST_DAY_OF_MONTH.Day() - CURRENT.Day()) + 5)
// }

// func CALCULATE_WEEKMAX() {
// 	DAYS_UNTIL_SUNDAY = time.Sunday - CURRENT.Weekday()

// 	if CURRENT.Weekday() == time.Sunday {
// 		DAYS_UNTIL_SUNDAY += 7
// 	} else {
// 		DAYS_UNTIL_SUNDAY += 8
// 	}

// 	WEEKMAX = DAYMAX * float64(DAYS_UNTIL_SUNDAY)
// }

// func CHECK_WEEKEND(DaysLeftBeforePayday int) int {
// 	AddDays := DaysLeftBeforePayday
// 	NextPayDayDate := time.Now().AddDate(0, 0, DaysLeftBeforePayday)

// 	if NextPayDayDate.Weekday() == time.Saturday {
// 		AddDays += 2
// 	}

// 	if NextPayDayDate.Weekday() == time.Sunday {
// 		AddDays += 1
// 	}

// 	return AddDays
// }
