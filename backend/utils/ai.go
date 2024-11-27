package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/Rishi-Mishra0704/flowforge/backend/config"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func AskAI(content string, config config.Config) genai.Part {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GeminiApi))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	prompt := fmt.Sprintf(`
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

Codebase:
%s`, content)

	// Call CountTokens to get the input token count (`total tokens`).
	tokResp, err := model.CountTokens(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("total_tokens:", tokResp.TotalTokens)
	// ( total_tokens: 10 )

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	part := printResponse(resp)
	return part
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
