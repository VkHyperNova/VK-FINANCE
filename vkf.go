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
	displayAndHandleCommandLineCommands()
}

func displayAndHandleCommandLineCommands() {
	// Clear the screen
	clearScreen()

	// Validate the database
	validateDatabase()

	// Fetch the finance data
	fetchFinanceData()

	// Print all the data
	displayAllVariables()

	// Display a separator
	displaySeparatorDoubleDash()

	// Print the program options
	printCyan("Program Options: \n\n")

	// Display the command names
	displayCommandName("add")
	displayCommandName("bills")
	displayCommandName("gas")
	displayCommandName("food")
	displayCommandName("other")
	displayCommandName("grow")
	displayCommandName("reset")
	displayCommandName("q")

	// Get the user input
	var user_input string
	printGray("\n\n=> ")
	fmt.Scanln(&user_input)

	// Handle the user input
	for {
		switch user_input {
		case "add":
			calculateCashFlowTransactions("Add")
		case "bills":
			calculateCashFlowTransactions("Bills")
		case "gas":
			calculateCashFlowTransactions("Gas")
		case "food":
			calculateCashFlowTransactions("Food")
		case "other":
			calculateCashFlowTransactions("Other")
		case "grow":
			growNetWorth()
		case "reset":
			resetVariables()
			SaveData("Reset", 0)
			displayAndHandleCommandLineCommands()
		case "q":
			clearScreen()
			os.Exit(0)
		default:
			displayAndHandleCommandLineCommands()
		}
	}
}

func calculateIncomeAndBalance(amount float64) {
	// Add amount to income
	INCOME = INCOME + amount
	// Add amount to balance
	BALANCE = BALANCE + amount
}

func calculateExpensesAndBalance(amount float64) {
	// subtract amount from balance
	BALANCE = BALANCE - amount
	// subtract amount from expenses
	EXPENSES = EXPENSES - amount
}

// growNetWorth function increases the net worth by the balance amount
func growNetWorth() {
	// increase net worth by balance amount
	NET_WORTH = NET_WORTH + BALANCE
	// set saved amount to balance amount
	SAVED_AMOUNT := BALANCE
	// reset balance to 0
	BALANCE = 0
	// reset other variables
	resetVariables()
	// save data to file
	SaveData("Grow", SAVED_AMOUNT)
	// display and handle command line commands
	displayAndHandleCommandLineCommands()
}

// Reset all variables to 0
func resetVariables() {
	INCOME = 0
	EXPENSES = 0
	BILLS = 0
	GAS = 0
	FOOD = 0
	OTHER = 0
}

// calculateCashFlowTransactions calculates cash flow transactions
func calculateCashFlowTransactions(name string) {
	// getUserInput retrieves user input for a specific transaction
	sum_of_money := getUserInput(name + ": ")

	// switch statement handles different transaction types
	switch name {
	case "Bills":
		// subtract transaction amount from BILLS
		BILLS -= sum_of_money
		// calculate expenses and balance
		calculateExpensesAndBalance(sum_of_money)
	case "Gas":
		// subtract transaction amount from GAS
		GAS -= sum_of_money
		// calculate expenses and balance
		calculateExpensesAndBalance(sum_of_money)
	case "Food":
		// subtract transaction amount from FOOD
		FOOD -= sum_of_money
		// calculate expenses and balance
		calculateExpensesAndBalance(sum_of_money)
	case "Other":
		// subtract transaction amount from OTHER
		OTHER -= sum_of_money
		// calculate expenses and balance
		calculateExpensesAndBalance(sum_of_money)
	case "Add":
		// calculate income and balance
		calculateIncomeAndBalance(sum_of_money)
	}

	// SaveData saves transaction data to a file
	SaveData(name, sum_of_money)
	// displayAndHandleCommandLineCommands displays the updated data and handles user commands
	displayAndHandleCommandLineCommands()
}

/* MAIN FINANCE VARIABLES */
var financeJsonObject finance
var NET_WORTH float64
var BALANCE float64
var EXPENSES float64
var BILLS float64
var GAS float64
var FOOD float64
var OTHER float64
var INCOME float64
var PERFECT_SAVE float64

// FetchFinanceData reads the finance data from a file and stores it in variables
func fetchFinanceData() {
    // Read the finance data from a file
    byteArray := readFile("./finance.json")

    // Convert the byte array to a FinanceJsonObject
    financeJsonObject = byteToFinanceJsonObject(byteArray)

    // Store the values from the FinanceJsonObject in variables
    NET_WORTH = financeJsonObject.NET_WORTH
    BALANCE = financeJsonObject.BALANCE
    EXPENSES = financeJsonObject.EXPENSES
    BILLS = financeJsonObject.BILLS
    GAS = financeJsonObject.GAS
    FOOD = financeJsonObject.FOOD
    OTHER = financeJsonObject.OTHER
    INCOME = financeJsonObject.INCOME

    // Calculate the perfect save amount
    PERFECT_SAVE = INCOME * 0.25
}

func displayProgramStart() {
    // Display separator line
    displaySeparatorSingleDash()

    // Display program title and version
    printGray("============== VK FINANCE v1 ===============\n")

    // Display separator line
    displaySeparatorSingleDash()

    // Display current month
    printGray(financeJsonObject.MONTH.String() + "\n")
}

// Function to display current month history
func displayCurrentMonthHistory() {
	// Get current date and time
	now := time.Now()
	// Get current month
	CurrentMonth := now.Month()

	// Read history.json file and convert it to byte array
	byteArray := readFile("./history.json")
	// Convert byte array to historyJsonArray
	historyJsonArray := byteToHistoryJsonArray(byteArray)

	// Clear the screen
	clearScreen()

	// Print cyan color text
	printCyan("History: \n")

	// Loop through historyJsonArray
	for _, value := range historyJsonArray {
		// Define date layout format
		layout := "02-01-2006"

		// Parse date string to time.Time object
		t, err := time.Parse(layout, value.DATE)
		// Handle error if any
		handleError(err)

		// Check if the month of the current date is equal to the current month
		if t.Month() == CurrentMonth {
			// Print the value
			fmt.Println(value)
		}
	}
}

func displayCommandName(name string) {
	printCyan("[")
	printYellow(name)
	printCyan("] ")
}

// Display separator with single dash
func displaySeparatorSingleDash() {
	printGray("============================================\n")
}

// Display separator with double dashes
func displaySeparatorDoubleDash() {
	printGray("--------------------------------------------\n")
}

func displayNetWorth() {
	printCyan("NET WORTH: ")
	printGreen(Float64ToStringWithTwoDecimalPoints(NET_WORTH) + " EUR\n\n")
}

func displayIncome() {
	printCyan("INCOME: ")
	printGreen("+" + Float64ToStringWithTwoDecimalPoints(INCOME) + " EUR\n")
}

func displayAllExpences() {
	printCyan("EXPENCES: ")
	printRed(Float64ToStringWithTwoDecimalPoints(EXPENSES) + " EUR\n\n")

	printCyan("-> Bills: ")
	printRed(Float64ToStringWithTwoDecimalPoints(BILLS) + " EUR\n")
	printCyan("-> Gas: ")
	printRed(Float64ToStringWithTwoDecimalPoints(GAS) + " EUR\n")
	printCyan("-> Food: ")
	printRed(Float64ToStringWithTwoDecimalPoints(FOOD) + " EUR\n")
	printCyan("-> Other: ")
	printRed(Float64ToStringWithTwoDecimalPoints(OTHER) + " EUR\n\n")
}

func displayBalance() {
	printCyan("BALANCE: ")
	printYellow(Float64ToStringWithTwoDecimalPoints(BALANCE) + " EUR\n")
}

// Calculate estimated daily spending amount
func calculateEstimatedDaylySpendingAmount() {
	// Calculate maximum savings budget per day
	MaxSavingsBudgetDay := (INCOME - PERFECT_SAVE) / 31
	// Calculate maximum spendable amount per day
	MaxSpendableAmountDay := EXPENSES / 31
	// Print estimated daily savings budget
	printCyan("ESTIMATED DAY: ")
	printGreen(Float64ToStringWithTwoDecimalPoints(MaxSavingsBudgetDay) + " EUR")
	printCyan(" | ")
	// Print estimated daily spendable amount
	printRed(Float64ToStringWithTwoDecimalPoints(MaxSpendableAmountDay) + " EUR\n")
}

// Calculate estimated weekly spending amount
func calculateEstimatedWeeklySpendingAmount() {
	// Max savings budget calculation
	MaxSavingsBudgetWeek := ((INCOME - PERFECT_SAVE) / 31) * 7
	// Max spendable amount calculation
	MaxSpendableAmountWeek := (EXPENSES / 31) * 7
	// Print estimated weekly spending amount
	printCyan("ESTIMATED WEEK: ")
	printGreen(Float64ToStringWithTwoDecimalPoints(MaxSavingsBudgetWeek) + " EUR")
	printCyan(" | ")
	printRed(Float64ToStringWithTwoDecimalPoints(MaxSpendableAmountWeek) + " EUR\n")
}

// Display saving amount
func displaySavingAmount() {
    // Print cyan text
    printCyan("SAVING (25%): ")
    // Convert float to string with two decimal points
    savingAmount := Float64ToStringWithTwoDecimalPoints(PERFECT_SAVE)
    // Print green text
    printGreen(savingAmount + " EUR\n\n")
}

// calculateMoneyLeft calculates the money left after saving for a perfect save
func calculateMoneyLeft() {
	// Calculate the money left after saving for a perfect save
	MONEY := BALANCE - PERFECT_SAVE
	
	// Print the text "MONEY: " in cyan color
	printCyan("MONEY: ")
	
	// Check if the money left is less than 0
	if MONEY < 0 {
		// Print the money left in red color with two decimal points
		printRed(Float64ToStringWithTwoDecimalPoints(MONEY) + " EUR\n\n")
	} else {
		// Print the money left in green color with two decimal points
		printGreen(Float64ToStringWithTwoDecimalPoints(MONEY) + " EUR\n\n")
	}
}

func displayAllVariables() {
    // Display program start
    displayProgramStart()
    
    // Display current month history
    displayCurrentMonthHistory()
    
    // Display net worth
    displayNetWorth()
    
    // Display income
    displayIncome()
    
    // Display all expenses
    displayAllExpences()
    
    // Calculate estimated daily spending amount
    calculateEstimatedDaylySpendingAmount()
    
    // Calculate estimated weekly spending amount
    calculateEstimatedWeeklySpendingAmount()
    
    // Display saving amount
    displaySavingAmount()
    
    // Display balance
    displayBalance()
    
    // Calculate money left
    calculateMoneyLeft()
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

func printRed(a string) {
	fmt.Print(Red + a + Reset)
}

func printGreen(a string) {
	fmt.Print(Green + a + Reset)
}

func printYellow(a string) {
	fmt.Print(Yellow + a + Reset)
}

func printBlue(a string) {
	fmt.Print(Blue + a + Reset)
}

func printPurple(a string) {
	fmt.Print(Purple + a + Reset)
}

func printCyan(a string) {
	fmt.Print(Cyan + a + Reset)
}

func printGray(a string) {
	fmt.Print(Gray + a + Reset)
}

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

// GetFinanceData returns a finance struct with financial data
func getFinanceData() finance {

	// Get the current time
	now := time.Now()

	// Return a finance struct with financial data rounded to 2 decimal places
	return finance{
		NET_WORTH: math.Round(NET_WORTH*100) / 100,
		INCOME:    math.Round(INCOME*100) / 100,
		BALANCE:   math.Round(BALANCE*100) / 100,
		EXPENSES: math.Round(EXPENSES*100) / 100,
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

// getHistoryJsonArrayObject creates a history object with the current date, time, action, and value.
func getHistoryJsonArrayObject(action string, value float64) history {

	// Get the current time.
	now := time.Now()

	// Format the current time as a string.
	formattedTime := now.Format("15:04:05")

	// Format the current date as a string.
	formattedDate := now.Format("02-01-2006")

	// Return the history object with the current date, time, action, and value.
	return history{
		DATE:   formattedDate,
		TIME:   formattedTime,
		ACTION: action,
		VALUE: value,
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
    data := getFinanceData()

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
	printCyan("\n" + question)
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
		printRed(err.Error() + "\n")
	}
}
