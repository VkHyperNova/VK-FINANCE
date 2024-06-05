package util

import "fmt"

/* Colors for Printing */

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

/* Red */
func PrintRedString(a string) {
	fmt.Print(Red + a + Reset)
}

func PrintRedFloat(a float64) {
	fmt.Print(Red, a, Reset)
}

/* Green */
func PrintGreenString(a string) {
	fmt.Print(Green + a + Reset)
}

func PrintGreenFloat(a float64) {
	fmt.Print(Green, a, Reset)
}

/* Yellow */
func PrintYellowString(a string) {
	fmt.Print(Yellow + a + Reset)
}

func PrintYellowFloat(a float64) {
	fmt.Print(Yellow, a, Reset)
}

/* Blue */
func PrintBlueString(a string) {
	fmt.Print(Blue + a + Reset)
}

func PrintBlueFloat(a float64) {
	fmt.Print(Blue, a, Reset)
}

/* Purple */
func PrintPurpleString(a string) {
	fmt.Print(Purple + a + Reset)
}

func PrintPurpleFloat(a float64) {
	fmt.Print(Purple, a, Reset)
}

/* Cyan */
func PrintCyanString(a string) {
	fmt.Print(Cyan + a + Reset)
}

func PrintCyanFloat(a float64) {
	fmt.Print(Cyan, a, Reset)
}

/* Gray */
func PrintGrayString(a string) {
	fmt.Print(Gray + a + Reset)
}

func PrintGrayFloat(a float64) {
	fmt.Print(Gray, a, Reset)
}
