/*
	About: Language based Activity tracking
*/

package language

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Structs
type Language struct {
	Name string `json:"name"`
}

// Variables
var filePath = "./data/languages.json"

// Functions
func GetLanguages() []Language {
	// Read the JSON file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Unmarshal JSON data into a struct
	var languages []Language
	err = json.Unmarshal(data, &languages)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Print the loaded data
	for i, language := range languages {
		fmt.Println("Language:", i+1)
		fmt.Println("Name:", language.Name)
		fmt.Println()
	}

	return languages
}

func AddLanguage(language string) {
	// Read the JSON file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Unmarshal JSON data into a struct
	var languages []Language
	err = json.Unmarshal(data, &languages)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Add the New Language
	newLanguage := Language{Name: language}
	languages = append(languages, newLanguage)

	// Marshal the updated object back to JSON
	updatedData, err := json.MarshalIndent(languages, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Write the updated JSON to the file
	err = ioutil.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	fmt.Println("JSON object updated successfully.")
}
