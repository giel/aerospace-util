package asu

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

func GetWorkspaces() ([]string, error) {
	cmd := exec.Command("aerospace", "list-workspaces", "--all")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return parseLines(output), nil
}

func GetWindows(workspace string) ([]string, error) {
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
