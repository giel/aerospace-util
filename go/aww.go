package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Step 1: Get all workspaces
	workspaces, err := getWorkspaces()
	if err != nil {
		fmt.Println("Error getting workspaces:", err)
		return
	}

	// Step 2: For each workspace, list its windows
	for _, workspace := range workspaces {
		windows, err := getWindows(workspace)
		if err != nil {
			fmt.Printf("  Error getting windows: %v\n", err)
			continue
		}
		if len(windows) != 0 {
			// fmt.Printf("Workspace: %s\n", workspace)
			for _, window := range windows {
				fmt.Printf("%s  %s\n", workspace, window)
			}
		}
	}
}

func getWorkspaces() ([]string, error) {
	cmd := exec.Command("aerospace", "list-workspaces", "--all")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return parseLines(output), nil
}

func getWindows(workspace string) ([]string, error) {
	cmd := exec.Command("aerospace", "list-windows", "--workspace", workspace)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return parseLines(output), nil
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
