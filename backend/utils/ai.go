package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/Rishi-Mishra0704/flowforge/backend/config"
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
	prompt := fmt.Sprintf(`
<<<<<<< HEAD
	Create a minimal and structured flowchart in JSON format based on the given codebase. 
	The flowchart should be constructed using the following models:
	
	1. **Node**:
	   - Fields:
		 - "id" (int): A unique identifier for each node.
		 - "label" (string): A description of the action or decision represented by the node.
		 - "type" (string): The type of node, which can be one of the following: "Start", "Process", "Decision", or "End".
	
	2. **Edge**:
	   - Fields:
		 - "source" (int): The ID of the starting node for this edge.
		 - "target" (int): The ID of the ending node for this edge.
		 - "condition" (string, optional): A label for the condition leading to the edge (used for Decision nodes).
	
	3. **Flowchart**:
	   - Fields:
		 - "nodes" ([]Node): A list of all nodes in the flowchart.
		 - "edges" ([]Edge): A list of all edges connecting the nodes.
	
	Instructions:
	- Analyze the given codebase to identify logical steps and control flow.
	- Create nodes for each logical step in the code, with appropriate types and labels.
	- Connect the nodes using edges to represent the flow of execution, including conditions for Decision nodes.
	- Ensure the flowchart accurately represents the logic of the code, focusing on clarity and structure.
	
	Return the flowchart as a JSON object in the following format:
	{
	  "nodes": [
		{ "id": 1, "label": "Start", "type": "Start" },
		{ "id": 2, "label": "Some Process", "type": "Process" },
		{ "id": 3, "label": "Some Decision", "type": "Decision" },
		{ "id": 4, "label": "End", "type": "End" }
	  ],
	  "edges": [
		{ "source": 1, "target": 2 },
		{ "source": 2, "target": 3 },
		{ "source": 3, "target": 4, "condition": "Yes" },
		{ "source": 3, "target": 5, "condition": "No" }
	  ]
	}
	
	Codebase:
	%s`, content)
=======
Create a minimal and structured flowchart in JSON format based on the given codebase. 
The flowchart should contain the following types of shapes:
- "Start": Represent the beginning of the flow.
- "Process": Represent basic operations or steps in the code.
- "Decision": Represent conditional statements or branches in the code.
- "End": Represent the conclusion of the flow.

Ensure the flowchart uses only these shapes and focuses on the logic of the code. 
The flowchart should not just be shapes added arbitrarily to lines of code, but should reflect the actual flow of execution, including decisions and processes, as well as how they are connected. 

Return the flowchart as a JSON object with nodes and edges, where:
- Each node should represent a logical step in the code with a unique identifier, a shape, and the action or decision being represented.
- Each edge should represent the flow from one node to another with labels for the condition (for Decision nodes).

Don't add file extension.

Codebase:
%s`, content)
>>>>>>> 690f5b6 (json file creation)

	// Call CountTokens to get the input token count (`total tokens`).
	tokResp, err := model.CountTokens(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("total_tokens:", tokResp.TotalTokens)

	// Generate the response
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	part := printResponse(resp)

	// Type assertion for genai.Text
	text, ok := part.(genai.Text)
	if !ok {
		log.Fatal("Failed to convert Part to Text")
	}

	return string(text)
}

func printResponse(resp *genai.GenerateContentResponse) genai.Part {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
				return part
			}
		}
	}
	return nil
}
