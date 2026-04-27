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

/* Storage */

const (
	driveLabel      = "VK DATA"
	driveMountPoint = "/media/veikko/VK DATA"
	driveDevice     = "/dev/sda2"
)

func InitStorage() (bool, error) {

	if err := ensureFile(config.LocalFile, config.DefaultContent); err != nil {
		return false, err
	}

	mounted, err := IsDriveMounted()
	if err != nil {
		return false, fmt.Errorf("mount check failed: %w", err)
	}

	if !mounted {
		input, err := PromptWithSuggestion("Drive not mounted. Try to mount it? (y/n) ", "y")
		if err != nil {
			return false, err
		}
		if strings.ToLower(strings.TrimSpace(input)) != "y" {
			return false, nil
		}
		if err := unlockAndMount(); err != nil {
			return false, fmt.Errorf("failed to mount drive: %w", err)
		}
		if mounted, err = IsDriveMounted(); err != nil || !mounted {
			return false, fmt.Errorf("drive still not mounted after mount attempt")
		}
		// Program did the mounting
		return true, ensureFile(config.BackupFile, config.DefaultContent)
	}

	// Drive was already mounted manually
	if err := ensureFile(config.BackupFile, config.DefaultContent); err != nil {
		return false, err
	}

	return false, nil
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

func IsDriveMounted() (bool, error) {
	if runtime.GOOS != "linux" {
		return false, fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	file, err := os.Open("/proc/mounts")
	if err != nil {
		return false, fmt.Errorf("cannot open /proc/mounts: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), " ", 3)
		if len(parts) >= 2 && parts[0] == driveDevice {
			return true, nil
		}
	}
	return false, scanner.Err()
}

func unlockAndMount() error {
	device, err := findDeviceByLabel(driveLabel)
	if err != nil {
		return fmt.Errorf("could not find drive: %w", err)
	}
	fmt.Printf("Found drive at %s\n", device)

	if err := mountDevice(device, driveMountPoint); err != nil {
		return fmt.Errorf("mount failed: %w", err)
	}
	fmt.Printf("Drive mounted at %s\n", driveMountPoint)
	return nil
}

func findDeviceByLabel(label string) (string, error) {
	out, err := exec.Command("lsblk", "-o", "NAME,LABEL", "-r", "-n").Output()
	if err != nil {
		return "", fmt.Errorf("lsblk failed: %w", err)
	}

	for _, line := range strings.Split(string(out), "\n") {
		fields := strings.Fields(line)
		if len(fields) >= 2 && decodeLsblkLabel(fields[1]) == label {
			return "/dev/" + fields[0], nil
		}
	}
	return "", fmt.Errorf("label '%s' not found (is the drive plugged in?)", label)
}

func decodeLsblkLabel(s string) string {
	return strings.NewReplacer(
		`\x20`, " ",
		`\x09`, "\t",
		`\x0a`, "\n",
		`\x5c`, `\`,
	).Replace(s)
}

func mountDevice(device, mountPoint string) error {
	cmd := exec.Command("udisksctl", "mount", "-b", device)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func UnmountDrive() {
	fmt.Println("Unmounting drive...")
	cmd := exec.Command("udisksctl", "unmount", "-b", driveDevice)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Warning: failed to unmount drive:", err)
		return
	}
	fmt.Println("Drive unmounted successfully")
}

/* Other Functions */

func PromptWithSuggestion(name string, edit string) (string, error) {

	line := liner.NewLiner()
	defer line.Close()

	input, err := line.PromptWithSuggestion(name, edit, -1)
	if err != nil {
		return input, err
	}

	return input, nil
}

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
