package install

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func installSystemdService() error {
	serviceFilePath := fmt.Sprintf("/etc/systemd/system/%s", systemdServiceName)
	if err := os.WriteFile(serviceFilePath, []byte(systemdServiceContent), 0644); err != nil {
		return fmt.Errorf("failed to write systemd service file: %w", err)
	}
	if err := exec.Command("systemctl", "start", strings.Split(systemdServiceName, ".")[0]).Run(); err != nil {
		return fmt.Errorf("failed to start systemctl service: %w", err)
	}
	return nil
}

func installPlistService() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}
	serviceFilePath := fmt.Sprintf("%s/Library/LaunchDaemons/%s", home, plistServiceName)
	if err := os.WriteFile(serviceFilePath, []byte(plistServiceContent), 0644); err != nil {
		return fmt.Errorf("failed to write plist service file: %w", err)
	}
	if err := exec.Command("launchctl", "load", serviceFilePath).Run(); err != nil {
		return fmt.Errorf("failed to load plist service: %w", err)
	}
	return nil
}

func Install() error {
	switch runtime.GOOS {
	case "linux":
		return installSystemdService()
	case "darwin":
		return installPlistService()
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}
