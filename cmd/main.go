package main

import (
	"fmt"
	"tracker/internal/activities"
	"tracker/internal/menu"
)

func main() {
	// Load Initial Activity Tracker
	tracker := activities.NewActivityTracker()

	// Print Previous Results

	// Load Start Menu
	menu := menu.NewMenu(tracker)
	menu.Run()
	fmt.Println("Exiting...")
}
