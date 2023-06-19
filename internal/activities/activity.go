/*
	About: Activity Definitions
*/

package activities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Structs
type Activity struct {
	Type      string
	Date      time.Time
	Language  string
	TimeSpent time.Duration
}

type ActivityTracker struct {
	activities []Activity
}

// Constants
var statsFile = "./data/stats.json"

// Helper Functions
func loadActivities() []Activity {
	var activities []Activity

	// Load Previous Activity Stats
	data, err := ioutil.ReadFile(statsFile)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	err = json.Unmarshal(data, &activities)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	return activities
}

// Activity Tracker Functions
func NewActivityTracker() *ActivityTracker {
	activities := loadActivities()

	return &ActivityTracker{
		activities: activities,
	}
}

// Add Activity to Tracker
func (t *ActivityTracker) AddActivity(activity Activity) {
	t.activities = append(t.activities, activity)
}

// Get the Saved Activity Tracker details
func (t *ActivityTracker) GetActivities() []Activity {
	return t.activities
}

// Get total Activity Tracker time
func (t *ActivityTracker) GetTotalTime(language string) time.Duration {
	var totalTime time.Duration
	for _, activity := range t.activities {
		if activity.Language == language {
			totalTime += activity.TimeSpent
		}
	}
	return totalTime
}

// Save Activity
func (t *ActivityTracker) SaveActivity(activity Activity) {
	var activities []Activity

	// Check if the file already exists
	if _, err := os.Stat(statsFile); os.IsNotExist(err) {
		// Create the file
		file, err := os.Create(statsFile)
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
		file.Close()

		fmt.Println("File created successfully.")
	} else {
		fmt.Println("File already exists.")
		// Read the JSON file
		data, err := ioutil.ReadFile(statsFile)
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}

		err = json.Unmarshal(data, &activities)
		if err != nil {
			log.Fatalf("Failed to unmarshal JSON: %v", err)
		}
	}

	// Add the New Activity
	activities = append(activities, activity)

	// Marshal the updated object back to JSON
	updatedData, err := json.MarshalIndent(activities, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Write the updated JSON to the file
	err = ioutil.WriteFile(statsFile, updatedData, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	fmt.Println("JSON object updated successfully.")
}
