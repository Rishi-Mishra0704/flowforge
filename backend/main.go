package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Rishi-Mishra0704/flowforge/backend/config"
	"github.com/Rishi-Mishra0704/flowforge/backend/models"
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

	// Reset the file pointer to the beginning of the file
	if _, err := file.Seek(0, 0); err != nil {
		log.Fatalf("Failed to reset file pointer: %v", err)
	}

	// Decode the JSON data into the Flowchart struct
	var flowchart models.Flowchart
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&flowchart); err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	// Log the unmarshaled Flowchart struct
	log.Printf("Flowchart loaded successfully: %+v", flowchart)

	// Optionally, iterate over nodes and edges to demonstrate access
	for _, node := range flowchart.Nodes {
		log.Printf("Node: ID=%d, Label=%s, Type=%s", node.ID, node.Label, node.Type)
	}

	for _, edge := range flowchart.Edges {
		log.Printf("Edge: Source=%d, Target=%d, Condition=%s", edge.Source, edge.Target, edge.Condition)
	}

	log.Println("Result saved to result.json")
}
