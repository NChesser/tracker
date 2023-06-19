package main

import (
	"fmt"
)

func main() {
	// Load Initial Activity Tracker
	tracker := NewActivityTracker()

	// Print Previous Results

	// Load Start Menu
	menu := NewMenu(tracker)
	menu.Run()
	fmt.Println("Exiting...")
}
