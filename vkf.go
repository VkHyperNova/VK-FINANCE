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
	clearScreen()
	validateDatabase()
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
			SaveData("Reset", 0)
			CL()
		case "q":
			clearScreen()
			os.Exit(0)
		default:
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

func GROW() {
	NET_WORTH = NET_WORTH + BALANCE
	SAVED_AMOUNT := BALANCE
	BALANCE = 0
	RESET()
	SaveData("Grow", SAVED_AMOUNT)
	CL()
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

	SaveData(name, amount)
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
	data := readFile("./finance.json")
	DATABASE = byteToFinanceJsonObject(data)
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
	hdata := byteToHistoryJsonArray(file)

	clearScreen()

	PRINT_CYAN("History: \n")

	for _, value := range hdata {
		layout := "02-01-2006"

		t, err := time.Parse(layout, value.DATE)
		handleError(err)

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

func getHistoryJsonArrayObject(LAST_ACTION string, VALUE float64) history {

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

func byteToHistoryJsonArray(byteArray []byte) []history {

	historyJsonArray := []history{}

	err := json.Unmarshal(byteArray, &historyJsonArray)
	handleError(err)

	return historyJsonArray
}



func byteToFinanceJsonObject(byteArray []byte) finance {

	financeJsonObject := finance{}

	err := json.Unmarshal(byteArray, &financeJsonObject)
	handleError(err)

	return financeJsonObject
}

// SaveData saves the given name and amount to the database and history
func SaveData(name string, amount float64) {
    // Save the name and amount to the database
    SaveToDatabase()

    // Save the name and amount to the history
    SaveHistory(name, amount)
}

// SaveToDatabase saves finance data to a JSON file
func SaveToDatabase() {
    // Construct finance data as a JSON object
    data := CONCSTRUCT_FINANCE_JSON()

    // Convert finance data to a byte array
    dataBytes := interfaceToByteArray(data)

    // Write finance data to a JSON file
    writeDataToFile("./finance.json", dataBytes)
}

// SaveHistory saves the last action and its value to the history file
func SaveHistory(Action string, Value float64) {

	// Read the history file content
	historyByteArray := readFile("./history.json")

	// Convert the file content to history data
	historyJsonArray := byteToHistoryJsonArray(historyByteArray)

	// Construct a new history JSON object
	historyJsonArrayObject := getHistoryJsonArrayObject(Action, Value)

	// Append the new data to the history data
	historyJsonArray = append(historyJsonArray, historyJsonArrayObject)

	// Convert the history data to a byte array
	historyByteArrayUpdated := interfaceToByteArray(historyJsonArray)

	// Write the data to the history file
	writeDataToFile("./history.json", historyByteArrayUpdated)
}


// Validate the existence of the database files
func validateDatabase() {
	
	if !doesDirectoryExist("./finance.json") {
		NET_WORTH = getUserInput("NET_WORTH: ")
		SaveToDatabase()
	}

	if !doesDirectoryExist("./history.json") {
		writeDataToFile("./history.json", []byte("[]"))
	}
}

func readFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	handleError(err)
	return file
}

func writeDataToFile(filename string, dataBytes []byte) {

	var err = os.WriteFile(filename, dataBytes, 0644)
	handleError(err)
}

func doesDirectoryExist(dir_name string) bool {

	if _, err := os.Stat(dir_name); os.IsNotExist(err) {
		return false
	}
	return true
}

func interfaceToByteArray(data interface{}) []byte {
	dataBytes, err := json.MarshalIndent(data, "", "  ")
	handleError(err)

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

func handleError(err error) {
	if err != nil {
		PRINT_RED(err.Error() + "\n")
	}
}
