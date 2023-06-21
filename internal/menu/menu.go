/*
	About: Handle Menu for saving activity time
*/

package menu

// Imports

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"tracker/internal/activities"
	"tracker/internal/language"
)

// Structs
type Menu struct {
	tracker *activities.ActivityTracker
	scanner *bufio.Scanner
}

func NewMenu(tracker *activities.ActivityTracker) *Menu {
	return &Menu{
		tracker: tracker,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (m *Menu) displayMenu() {
	fmt.Println("1. Add Activity")
	fmt.Println("2. View Activities")
	fmt.Println("3. View Total Time for Language")
	fmt.Println("4. Exit")
}

func (m *Menu) displayActivityTypes() {
	fmt.Println("1. Listening")
	fmt.Println("2. Watching")
	fmt.Println("3. Talking")
}

func (menu *Menu) handleAddActivity() {
	// Extract Languages
	languages := language.GetLanguages()

	fmt.Print("Enter the language: ")
	language := menu.readLine()

	// Add Language to JSON file if not there
	languageStatus := languageExists(language, languages)
	fmt.Println("Language Status", languageStatus)

	if !languageStatus {
		fmt.Println("Language Exists")
	} else {
		fmt.Println("Adding Language")
		language.AddLanguage()
	}

	fmt.Print("Enter the time spent (in minutes): ")
	timeSpentStr := menu.readLine()
	timeSpent, err := strconv.Atoi(timeSpentStr)
	if err != nil {
		fmt.Println("Invalid time duration.")
		return
	}

	// fmt.Print("Enter Activity Type")
	// activityTypeStr := menu.runActivityType()
	// activityType, err := strconv.Atoi(activityTypeStr)
	// if err != nil {
	// 	fmt.Println("Invalid activity type.")
	// 	return
	// }

	// Add Activity
	newActivity := activities.Activity{
		// Type:      activityType,
		Date:      time.Now(),
		Language:  language,
		TimeSpent: time.Duration(timeSpent) * time.Minute,
	}
	menu.tracker.AddActivity(newActivity)
	menu.tracker.SaveActivity(newActivity)

	fmt.Println("Activity added successfully.")
}

func (m *Menu) handleViewActivities() {
	activities := m.tracker.GetActivities()
	fmt.Println("Activities:")
	for _, activity := range activities {
		fmt.Printf("Date: %s, Language: %s, Time Spent: %s\n", activity.Date.Format("2006-01-02 15:04:05"), activity.Language, activity.TimeSpent)
	}
}

func (m *Menu) handleViewTotalTime() {
	fmt.Print("Enter the language: ")
	language := m.readLine()

	totalTime := m.tracker.GetTotalTime(language)
	fmt.Printf("Total time spent on %s: %s\n", language, totalTime)
}

func (m *Menu) readLine() string {
	m.scanner.Scan()
	return m.scanner.Text()
}

func (m *Menu) getActivityType() {
	for {
		m.displayActivityTypes()
		fmt.Print("Enter activity type: ")
		choice := m.readLine()

		switch choice {
		case "1":
			"Listening"
		case "2":
			"Watching"
		case "3":
			"Talking"
		default:
			return
		}
	}

}

func (m *Menu) Run() {
	for {
		m.displayMenu()
		fmt.Print("Enter your choice: ")
		choice := m.readLine()

		switch choice {
		case "1":
			m.handleAddActivity()
		case "2":
			m.handleViewActivities()
		case "3":
			m.handleViewTotalTime()
		case "4":
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
