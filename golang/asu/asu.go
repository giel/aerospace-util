package asu

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

var AeroSpacePath string
var Version = "not set" // will be set by build.sh

func GetWorkspaces() ([]string, error) {
	cmd := exec.Command(AeroSpacePath, "list-workspaces", "--all")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return parseLines(output), nil
}

func GetWindows(workspace string) ([]string, error) {
  // aerospace list-windows --all --format "%{monitor-id}%{tab}%{app-name}%{tab}%{window-title}"
  // aerospace list-monitors
	cmd := exec.Command(AeroSpacePath, "list-windows", "--workspace", workspace)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return parseLines(output), nil
}

func GetWorkspaceWindowsSimple() ([]string, error) {
	var lines []string
	// Check if the executable exists
	var err error
	AeroSpacePath, err = findExecutable("aerospace")
	if err != nil {
		return []string{"aeropspace not found"}, err
	}
	// Step 1: Get all workspaces
	workspaces, err := GetWorkspaces()
	if err != nil {
		return lines, err
	}

	// Step 2: For each workspace, list its windows
	for _, workspace := range workspaces {
		windows, err := GetWindows(workspace)
		if err != nil {
			continue
		}
    for _, window := range windows {
		if len(windows) != 0 {
			txtString := fmt.Sprintf("%s  %s", workspace, window)
			lines = append(lines, txtString)
		}
    }
	}
	return lines, nil
}

func findExecutable(name string) (string, error) {
	// todo: fix this to avoid hardcoded path.
	// In awwapp it cannot find the executable, and in aww it CAN find the executable.
	return "/opt/homebrew/bin/aerospace", nil
	return "/opt/homebrew/bin/aerospace", nil
}

func parseLines(data []byte) []string {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var lines []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}
