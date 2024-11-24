package analyze

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/sashabaranov/go-openai"
)

var gptClient *openai.Client

// InitializeGPT initializes the GPT client.
func InitializeGPT(apiKey string) {
	gptClient = openai.NewClient(apiKey)
}

// ReviewFile sends a file's content to GPT for review.
func ReviewFile(filePath string) (string, error) {
	// Read the file content
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Create a context with a timeout for GPT requests
	ctx := context.Background()

	// Prepare the request for GPT
	request := openai.ChatCompletionRequest{
		Model: openai.GPT4, // Using GPT-4 for better results
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a code reviewer. Provide feedback on code quality, readability, and improvements.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("Review this code:\n\n%s", string(content)),
			},
		},
	}

	// Make the GPT API call
	response, err := gptClient.CreateChatCompletion(ctx, request)
	if err != nil {
		return "", fmt.Errorf("failed to review file %s: %w", filePath, err)
	}

	// Return the feedback from GPT
	return response.Choices[0].Message.Content, nil
}
