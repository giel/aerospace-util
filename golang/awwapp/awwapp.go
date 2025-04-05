package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang/asu"
)

var AppName = "awwapp"
var Version = "not set" // will be set by build.sh

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("awwapp")

	// Proportional font (default)
	proportionalLabel := widget.NewLabel("awwapp version: " + Version)

	// Step 1: Get all workspaces
	workspaces, err := asu.GetWorkspaces()
	if err != nil {
		fmt.Println("Error getting workspaces:", err)
		return
	}

	// Step 2: For each workspace, list its windows
	txtString := ""
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
				txtString += fmt.Sprintf("%s  %s\n", workspace, window)
			}
		}
	}
	// Monospaced font
	monospaceLabel := widget.NewLabel(txtString)
	monospaceLabel.TextStyle = fyne.TextStyle{Monospace: true}

	// Stack everything vertically
	content := container.NewVBox(
		proportionalLabel,
		monospaceLabel,
	)

	myWindow.SetContent(content)
	// myWindow.Resize(fyne.NewSize(400, 1000))
	myWindow.ShowAndRun()
}
