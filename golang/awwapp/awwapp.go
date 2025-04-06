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

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("awwapp")

	// Proportional font (default)
	proportionalLabel := widget.NewLabel("awwapp version: " + asu.Version)

	txtString := ""
	lines, err := asu.GetWorkspaceWindowsSimple()
	if err != nil {
		txtString = fmt.Sprintf("Error getting workspaces:", err)
	} else {
		for _, line := range lines {
			txtString += line + "\n"
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
