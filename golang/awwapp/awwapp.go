package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"golang/asu"
	"time"
)

var AppName = "awwapp"

func buildContent(window fyne.Window) fyne.CanvasObject {
	buttonRefresh := widget.NewButton("refresh (Command + R)", func() { refreshContent(window) })

	// Proportional font (default)
	const layout = "2006-01-02 15:04:05"
	now := time.Now().Format(layout)
	timeLabel := widget.NewLabel("window state of: " + now)
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
		buttonRefresh,
		timeLabel,
		proportionalLabel,
		monospaceLabel,
	)
	return content
}

func refreshContent(window fyne.Window) {
	content := buildContent(window)
	window.SetContent(content)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("awwapp")
	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyR, Modifier: fyne.KeyModifierSuper}
	myWindow.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		refreshContent(myWindow)
	})

	refreshContent(myWindow)

	// myWindow.Resize(fyne.NewSize(400, 1000))
	myWindow.ShowAndRun()
}
