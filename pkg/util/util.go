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
	if runtime.GOOS == "linux" { // check if the operating system is Linux
		cmd := exec.Command("clear") // execute the clear command
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") // execute the cls command
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Input() string {

	line := liner.NewLiner()
	defer line.Close()

	userInput, err := line.Prompt("=> ")
	if err != nil {
		panic(err)
	}
	return userInput
}

func ValidateFiles() {

	// folderName := "history"

	// if _, err := os.Stat(folderName); os.IsNotExist(err) {
	// 	_ = os.Mkdir(folderName, 0700)
	// 	fmt.Println("history folder created!")
	// }

	path := "./history.json"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.WriteFile(path, []byte([]byte(`{"history": []}`)), 0644)
		if err != nil {
			panic(err)
		}
		fmt.Print("\n=> " + path)
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

func isMounted(mountPoint string) (bool, error) {
    file, err := os.Open("/proc/mounts")
    if err != nil {
        return false, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        if len(fields) >= 2 && fields[1] == mountPoint {
            return true, nil
        }
    }

    return false, scanner.Err()
}

func IsVKDataMounted() bool {

	if runtime.GOOS != "linux" {
        fmt.Println("This program only works on Linux.")
        return false
    }

	mountPoint := "/media/veikko/VK\\040DATA" // change to your actual mount path

    mounted, err := isMounted(mountPoint)
    if err != nil {
        fmt.Println("Error:", err)
        return false
    }

    if mounted {
        fmt.Println(config.Green + "\nVK DATA is mounted" + config.Reset)
		return true
    } else {
        fmt.Println(config.Red + "\nVK DATA is NOT mounted" + config.Reset)
		return false
    }
}
