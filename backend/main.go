package main

import (
	"log"
	"os"

	"github.com/Rishi-Mishra0704/flowforge/backend/config"
	"github.com/Rishi-Mishra0704/flowforge/backend/utils"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	// Read multiple files from a specified directory
	dirPath := "./test" // Replace with the path of your directory containing multiple files
	content, err := utils.ReadMultipleFiles(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	// Log the combined content of all files
	// log.Println("Combined Codebase Content:\n", content)

	// Call AskAI with the combined content
	part := utils.AskAI(content, config)

	// Write the JSON data to a file
	file, err := os.OpenFile("result.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Failed to open result.json: %v", err)
	}
	defer file.Close()

	_, err = file.Write([]byte(part))
	if err != nil {
		log.Fatalf("Failed to write to result.json: %v", err)
	}

	log.Println("Result saved to result.json")
}
