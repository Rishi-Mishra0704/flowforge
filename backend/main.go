package main

import (
	"log"

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
	log.Println("Combined Codebase Content:\n", content)

	// Call AskAI with the combined content
	part := utils.AskAI(content, config)
	log.Println(part)
}
