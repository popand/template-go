package api

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

type ChatRequest struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func HandleOpenAIChat(c *gin.Context) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		c.JSON(500, gin.H{"error": "OpenAI API key not found"})
		return
	}

	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("Invalid request: %v", err)})
		return
	}

	client := openai.NewClient(apiKey)

	// Convert messages to OpenAI format
	messages := make([]openai.ChatCompletionMessage, len(req.Messages))
	for i, msg := range req.Messages {
		messages[i] = openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT4,
			Messages: messages,
			Stream:   true,
		},
	)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Error creating chat completion: %v", err)})
		return
	}
	defer stream.Close()

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	for {
		response, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			c.SSEvent("error", gin.H{"error": fmt.Sprintf("Error receiving response: %v", err)})
			return
		}

		c.SSEvent("message", response.Choices[0].Delta.Content)
		c.Writer.Flush()
	}
}
