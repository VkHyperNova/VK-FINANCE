package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

//go:generate goversioninfo

func main() {
	clearScreen()
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
	PRINT_COMMAND("q")

	var command string
	PRINT_GRAY("\n\n=> ")
	fmt.Scanln(&command)

	for {
		switch command {
		case "add":
			CALCULATE("Add")
		case "bills":
			CALCULATE("Bills")
		case "gas":
			CALCULATE("Gas")
		case "food":
			CALCULATE("Food")
		case "other":
			CALCULATE("Other")
		case "grow":
			GROW()
		case "reset":
			RESET()
			SAVE("Reset", 0)
		case "q":
			clearScreen()
			os.Exit(0)
		default:
			clearScreen()
			CL()
		}
	}
}




func ADD(amount float64) {
	INCOME = INCOME + amount
	BALANCE = BALANCE + amount
}

func EXP(amount float64) {
	BALANCE = BALANCE - amount
	EXPENSES = EXPENSES - amount
}

func GROW()	{
	NET_WORTH = NET_WORTH + BALANCE
	SAVED_AMOUNT := BALANCE
	BALANCE = 0
	RESET()
	SAVE("Grow", SAVED_AMOUNT)
}

func RESET() {
	INCOME = 0
	EXPENSES = 0
	BILLS = 0
	GAS = 0
	FOOD = 0
	OTHER = 0
}

func CALCULATE(name string) {
	amount := getUserInput(name + ": ")

	switch name {
	case "Bills":
		BILLS = BILLS - amount
		EXP(amount)
	case "Gas":
		GAS = GAS - amount
		EXP(amount)
	case "Food":
		FOOD = FOOD - amount
		EXP(amount)
	case "Other":
		OTHER = OTHER - amount
		EXP(amount)
	case "Add":
		ADD(amount)
	}

	SAVE(name, amount)
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
	data := readFile("./finance.json")
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

func PRINT_HISTORY() {
	now := time.Now()
	CurrentMonth := now.Month()

	file := readFile("./history.json")
	hdata := CONVERT_TO_HISTORY(file)

	clearScreen()

	PRINT_CYAN("History: \n")

	for _, value := range hdata {
		layout := "02-01-2006"

		t, err := time.Parse(layout, value.DATE)
		handleError(err, "PRINT_HISTORY")

		if t.Month() == CurrentMonth {
			fmt.Println(value)
		}
	}
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
	PRINT_GREEN(Float64ToStringWithTwoDecimalPoints(NET_WORTH) + " EUR\n\n")
}

func PRINT_INCOME() {
	PRINT_CYAN("INCOME: ")
	PRINT_GREEN("+" + Float64ToStringWithTwoDecimalPoints(INCOME) + " EUR\n")
}

func PRINT_EXPENCES() {
	PRINT_CYAN("EXPENCES: ")
	PRINT_RED(Float64ToStringWithTwoDecimalPoints(EXPENSES) + " EUR\n\n")

	PRINT_CYAN("-> Bills: ")
	PRINT_RED(Float64ToStringWithTwoDecimalPoints(BILLS) + " EUR\n")
	PRINT_CYAN("-> Gas: ")
	PRINT_RED(Float64ToStringWithTwoDecimalPoints(GAS) + " EUR\n")
	PRINT_CYAN("-> Food: ")
	PRINT_RED(Float64ToStringWithTwoDecimalPoints(FOOD) + " EUR\n")
	PRINT_CYAN("-> Other: ")
	PRINT_RED(Float64ToStringWithTwoDecimalPoints(OTHER) + " EUR\n\n")
}

func PRINT_BALANCE() {
	PRINT_CYAN("BALANCE: ")
	PRINT_YELLOW(Float64ToStringWithTwoDecimalPoints(BALANCE) + " EUR\n")
}

func PRINT_DAY() {
	PERFECT_DAY := (INCOME - PERFECT_SAVE) / 31
	REAL_DAY := EXPENSES / 31
	PRINT_CYAN("ESTIMATED DAY: ")
	PRINT_GREEN(Float64ToStringWithTwoDecimalPoints(PERFECT_DAY) + " EUR")
	PRINT_CYAN(" | ")
	PRINT_RED(Float64ToStringWithTwoDecimalPoints(REAL_DAY) + " EUR\n")
}

func PRINT_WEEK() {
	PERFECT_WEEK := ((INCOME - PERFECT_SAVE) / 31) * 7
	REAL_WEEK := (EXPENSES / 31) * 7
	PRINT_CYAN("ESTIMATED WEEK: ")
	PRINT_GREEN(Float64ToStringWithTwoDecimalPoints(PERFECT_WEEK) + " EUR")
	PRINT_CYAN(" | ")
	PRINT_RED(Float64ToStringWithTwoDecimalPoints(REAL_WEEK) + " EUR\n")
}

func PRINT_SAVING() {
	PRINT_CYAN("SAVING (25%): ")
	PRINT_GREEN(Float64ToStringWithTwoDecimalPoints(PERFECT_SAVE) + " EUR\n\n")
}

func PRINT_MONEY() {
	MONEY := BALANCE - PERFECT_SAVE
	PRINT_CYAN("MONEY: ")
	if MONEY < 0 {
		PRINT_RED(Float64ToStringWithTwoDecimalPoints(MONEY) + " EUR\n\n")
	} else {
		PRINT_GREEN(Float64ToStringWithTwoDecimalPoints(MONEY) + " EUR\n\n")
	}
}

func PRINT_ALL() {
	PRINT_PROGRAM_INFO()
	PRINT_HISTORY()

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
	handleError(err, "CONVERT_TO_FINANCE")

	return data
}

func SAVE_HISTORY(LAST_ACTION string, VALUE float64) {

	file := readFile("./history.json")
	hdata := CONVERT_TO_HISTORY(file)

	newdata := CONCSTRUCT_HISTORY_JSON(LAST_ACTION, VALUE)
	hdata = append(hdata, newdata)
	dataBytes := interfaceToByteArray(hdata)
	writeDataToFile("./history.json", dataBytes)
}

func CONVERT_TO_FINANCE(body []byte) finance {

	data := finance{}

	err := json.Unmarshal(body, &data)
	handleError(err, "CONVERT_TO_FINANCE")

	return data
}

func CREATE_DB() {
	NET_WORTH = getUserInput("NET_WORTH: ")
	clearScreen()
	SAVE_DB()
}

func SAVE_DB() {
	data := CONCSTRUCT_FINANCE_JSON()
	dataBytes := interfaceToByteArray(data)
	writeDataToFile("./finance.json", dataBytes)
}

func SAVE(name string, amount float64) {
	SAVE_DB()
	SAVE_HISTORY(name, amount)
	clearScreen()
	CL()
}

func CREATE_HISTROY_DB() {
	writeDataToFile("./history.json", []byte("[]"))
}

func CHECK_DB() {
	if !doesDirectoryExist("./finance.json") {
		CREATE_DB()
	}

	if !doesDirectoryExist("./history.json") {
		CREATE_HISTROY_DB()
	}
}

func createDirectory(dir_name string) {
	_ = os.Mkdir(dir_name, 0700)
}

func readFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	handleError(err, "ReadFile")
	return file
}

func writeDataToFile(filename string, dataBytes []byte) {

	var err = os.WriteFile(filename, dataBytes, 0644)
	handleError(err, "WRITE_FILE FUNCTION")
}

func doesDirectoryExist(dir_name string) bool {

	if _, err := os.Stat(dir_name); os.IsNotExist(err) {
		return false
	}
	return true
}

func interfaceToByteArray(data interface{}) []byte {
	dataBytes, err := json.MarshalIndent(data, "", "  ")
	handleError(err, "Convert_To_Byte")

	return dataBytes
}

func Float64ToStringWithTwoDecimalPoints(number float64) string {
	return fmt.Sprintf("%.2f", number)
}

func getUserInput(question string) float64 {
start:
	var answer string
	PRINT_CYAN("\n" + question)
	fmt.Scanln(&answer)

	if answer == "" {
		answer = "0"
	}

	floatValue, err := strconv.ParseFloat(answer, 64)
	if err != nil {
		fmt.Println("Must be a number!")
		fmt.Println(err)
		goto start
	}

	return floatValue
}

func clearScreen() {

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

func handleError(err error, location string) {
	if err != nil {
		PRINT_RED(" << Function name: ")
		PRINT_RED(location + " >>\n")
		PRINT_RED(err.Error() + "\n")
	}
}
