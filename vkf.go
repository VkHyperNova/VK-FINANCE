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

func main() {
	CLEAR_SCREEN()
	CL()
}

/* MAIN CMD */
func CL() {
	CHECK_DB()
	SETUP()
	PRINT_ALL()

	PRINT_CYAN("\n<< COMMANDS: ")
	PRINT_YELLOW("add")
	PRINT_CYAN(" | ")
	PRINT_RED("exp")
	PRINT_CYAN(" | ")
	PRINT_GREEN("grow")
	PRINT_CYAN(" | ")
	PRINT_RED("reset")
	PRINT_CYAN(" | ")
	PRINT_GRAY("q")
	PRINT_CYAN(" >>\n")
	PRINT_GRAY("=> ")

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
		case "reset":
			RESET_EXPENSES()
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

func RESET_EXPENSES() {
	EXPENSES = 0
	SAVE_DB()
	CLEAR_SCREEN()
	CL()
}

/* MAIN FINANCE VARIABLES */
var DATABASE finance
var NET_WORTH float64
var BALANCE float64
var EXPENSES float64
var INCOME float64
var PERFECT_SAVE float64

func SETUP() {
	data := READ_FILE("./finance.json")
	DATABASE = CONVERT_TO_FINANCE(data)
	NET_WORTH = DATABASE.NET_WORTH
	BALANCE = DATABASE.BALANCE
	EXPENSES = DATABASE.EXPENSES
	INCOME = BALANCE + (-1 * EXPENSES)
	PERFECT_SAVE = INCOME * 0.25

}

func PRINT_PROGRAM_INFO() {
	PRINT_CYAN("<<___________ VK FINANCE v1 ___________>>\n\n")
	PRINT_GRAY(DATABASE.MONTH.String() + "\n")
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
}

func PRINT_BALANCE() {
	PRINT_CYAN("BALANCE: ")
	PRINT_YELLOW(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(BALANCE) + " EUR\n\n")
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
		PRINT_RED(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(MONEY) + " EUR\n")
	} else {
		PRINT_GREEN(CONVERT_TO_TWO_DECIMAL_POINTS_STRING(MONEY) + " EUR\n")
	}
	
}

func PRINT_ALL() {
	PRINT_PROGRAM_INFO()

	PRINT_NET_WORTH()
	PRINT_INCOME()
	PRINT_EXPENCES()
	PRINT_BALANCE()
	

	PRINT_DAY()
	PRINT_WEEK()
	PRINT_SAVING()

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
	BALANCE   float64    `json:"balance"`
	EXPENSES  float64    `json:"expences"`
	MONTH     time.Month `json:"month"`
}

func CONCSTRUCT_FINANCE_JSON() finance {

	now := time.Now()

	return finance{
		NET_WORTH: NET_WORTH,
		BALANCE:   BALANCE,
		EXPENSES:  EXPENSES,
		MONTH:     now.Month(),
	}
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

func CHECK_DB() {
	if !DIR_CHECK("./finance.json") {
		CREATE_DB()
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
