package util

import (
	"bufio"
	"path/filepath"

	"fmt"
	"os"
	"os/exec"

	"runtime"
	"strings"

	"github.com/VkHyperNova/VK-FINANCE/pkg/color"
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

func Contains(arr []string, name string) bool {
	for _, n := range arr {
		if n == name {
			return true
		}
	}
	return false
}

func Colorize(line string, value float64, highlight bool) string {
	if highlight {
		return color.Bold + color.Yellow + line + color.Reset
	}
	switch {
	case value > 0:
		return color.Green + line + color.Reset
	case value < 0:
		return color.Red + line + color.Reset
	default:
		return line
	}
}

func ensureFile(path string, content string) error {

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("error creating directory for %s: %w", path, err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("error creating file %s: %w", path, err)
		}
	}

	return nil
}

func CreateFilesAndFolders() error {
	

	if err := ensureFile(config.LocalFile, config.DefaultContent); err != nil {
		return err
	}

	if !HardDriveMountCheck() {
		input := Input("Do you want to continue? (y/n) ")
		if strings.ToLower(strings.TrimSpace(input)) != "y" {
			fmt.Println("Exiting program.")
			os.Exit(0)
		}
	} else {
		if err := ensureFile(config.BackupFile, config.DefaultContent); err != nil {
			return err
		}
	}

	return nil
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

	fmt.Println(color.Red + "\nVK DATA is NOT mounted" + color.Reset)
	return false
}
