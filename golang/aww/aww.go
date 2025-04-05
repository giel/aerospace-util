package main

import (
	"fmt"
	"golang/asu"
)

var Version = "not set" // will be set by build.sh

func main() {
	// Step 1: Get all workspaces
	workspaces, err := asu.GetWorkspaces()
	if err != nil {
		fmt.Println("Error getting workspaces:", err)
		return
	}

	// Step 2: For each workspace, list its windows
	for _, workspace := range workspaces {
		windows, err := asu.GetWindows(workspace)
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
