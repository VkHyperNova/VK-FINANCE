package util

import (
	"bufio"

	"fmt"
	"os"
	"os/exec"

	"runtime"
	"strings"

	"github.com/VkHyperNova/VK-FINANCE/pkg/config"
	"github.com/peterh/liner"
)

/* Other Functions */

func PressAnyKey() {
	fmt.Print()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

func ClearScreen() {
	switch runtime.GOOS {
	case "linux": // check if the operating system is Linux
		cmd := exec.Command("clear") // execute the clear command
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") // execute the cls command
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Input(prompt string) string {

	line := liner.NewLiner()
	defer line.Close()

	userInput, err := line.Prompt(prompt)
	if err != nil {
		panic(err)
	}
	return userInput
}

func CreateDatabaseFile() {

	if _, err := os.Stat(config.DefaultPath); os.IsNotExist(err) {
		err = os.WriteFile(config.DefaultPath, []byte([]byte(`{"vk-finance": []}`)), 0644)
		if err != nil {
			panic(err)
		}
		fmt.Print("\n=> " + config.DefaultPath)
	}
}

func WriteToFile(filename string, dataBytes []byte) {
	var err = os.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func Contains(arr []string, name string) bool {
	for _, n := range arr {
		if n == name {
			return true
		}
	}
	return false
}

func HardDriveMountCheck() bool {
	if runtime.GOOS != "linux" {
		fmt.Println("This program only works on Linux.")
		return false
	}

	mountPoint := "/media/veikko/VK\\040DATA" // match /proc/mounts format

	file, err := os.Open("/proc/mounts")
	if err != nil {
		fmt.Println("Cannot open /proc/mounts:", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 2 && fields[1] == mountPoint {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning /proc/mounts:", err)
		return false
	}

	fmt.Println(config.Red + "\nVK DATA is NOT mounted" + config.Reset)
	return false
}

func Colorize(line string, value float64, highlight bool) string {
	if highlight {
		return config.Bold + config.Yellow + line + config.Reset
	}
	switch {
	case value > 0:
		return config.Green + line + config.Reset
	case value < 0:
		return config.Red + line + config.Reset
	default:
		return line
	}
}
