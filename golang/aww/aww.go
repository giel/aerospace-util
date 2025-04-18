package main

import (
	"flag"
	"fmt"
	"golang/asu"
	"os"
)

var AppName = "aww"

func main() {
	versionFlag := flag.Bool("version", false, "Display the version of the program")

	flag.Parse()

	// Handle the --version flag
	if *versionFlag {
		fmt.Printf("%s version %s\n", AppName, asu.Version)
		os.Exit(0)
	}

	lines, err := asu.GetWorkspaceWindowsSimple()
	if err != nil {
		fmt.Println("Error getting workspaces:", err)
		os.Exit(1)
	}
	for _, line := range lines {
		fmt.Println(line)
	}

}
