package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/Rishi-Mishra0704/flowforge/backend/config"
	"github.com/Rishi-Mishra0704/flowforge/backend/models"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func AskAI(content string, config config.Config) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GeminiApi))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	example := `{
		"nodes": [],
		"edges": []
	  }`

	prompt := fmt.Sprintf(`
	  Create a minimal and structured flowchart in JSON format based on the given codebase. 
	  The flowchart should contain the following types of shapes:
	  - "Start": Represent the beginning of the flow.
	  - "Process": Represent basic operations or steps in the code.
	  - "Decision": Represent conditional statements or branches in the code.
	  - "End": Represent the conclusion of the flow.
	  
	  Ensure the flowchart uses only these shapes and focuses on the logic of the code. 
	  The flowchart should not just be shapes added arbitrarily to lines of code but should reflect the actual flow of execution, including decisions and processes, as well as how they are connected. 
	  
	  Return the flowchart as a plain JSON object without any additional formatting markers such as backticks or language tags. 
	  
	  The structure of the output should follow this example:
	  %s
	  
	  Use the following definitions for guidance:
	  Node: %v
	  Edge: %v
	  nodes have an "id", "label", and "type" field that represent the node's ID, label, and type.
	  edges have a "source", "target", and "condition" field that represent the edge's source node ID, target node ID, and condition.
	  Fill all the data based on the given Nodes, Edges, and codebase.

	  it will be a typical flowchart with nodes and edges.
	  Codebase:
	  %s
	  `, example, models.Node{}, models.Edge{}, content)

	// Generate the response
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	formattedResp := getResponse(resp)

	// Type assertion for genai.Text
	text, ok := formattedResp.(genai.Text)
	if !ok {
		log.Fatal("Failed to convert response to Text")
	}

	return string(text)
}

func getResponse(resp *genai.GenerateContentResponse) genai.Part {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {

				return part
			}
		}
	}
	return nil
}
