package main

import (
	"encoding/json" // Importing the encoding/json package to handle JSON data
	"fmt" // Importing the fmt package for formatted I/O
	"log" // Importing the log package for logging errors
)

// User struct to match the JSON structure
type User struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	IsActive bool     `json:"is_active"`
	Roles    []string `json:"roles"`
}

func main() {
	// JSON string containing user information
	jsonString := `{"id": 123, "name": "Gemini AI", "is_active": true, "roles": ["admin", "editor"]}`

	// Variable to store the unmarshaled user data
	var user User

	// Unmarshal the JSON string into the User struct
	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Print user information in formatted way
	fmt.Println("User Information:")
	fmt.Printf("  ID: %d\n", user.ID)
	fmt.Printf("  Name: %s\n", user.Name)
	fmt.Printf("  Active: %v\n", user.IsActive)
	fmt.Printf("  Roles: %v\n", user.Roles)
}
