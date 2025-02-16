package main

import (
	"log"
	"os"
	"template-go/backend/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize Gin router
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Next.js frontend
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// API routes
	apiGroup := r.Group("/api")
	{
		// OpenAI routes
		openai := apiGroup.Group("/openai")
		{
			openai.POST("/chat", api.HandleOpenAIChat)
			openai.POST("/transcribe", handleOpenAITranscribe)
		}

		// Anthropic routes
		anthropic := apiGroup.Group("/anthropic")
		{
			anthropic.POST("/chat", handleAnthropicChat)
		}

		// Replicate routes
		replicate := apiGroup.Group("/replicate")
		{
			replicate.POST("/generate-image", handleReplicateGenerateImage)
		}

		// Deepgram routes
		deepgram := apiGroup.Group("/deepgram")
		{
			deepgram.GET("", handleDeepgramKey)
		}
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

// Remaining handler function placeholders
func handleOpenAITranscribe(c *gin.Context) {
	// TODO: Implement OpenAI transcribe handler
	c.JSON(501, gin.H{"error": "Not implemented"})
}

func handleAnthropicChat(c *gin.Context) {
	// TODO: Implement Anthropic chat handler
	c.JSON(501, gin.H{"error": "Not implemented"})
}

func handleReplicateGenerateImage(c *gin.Context) {
	// TODO: Implement Replicate image generation handler
	c.JSON(501, gin.H{"error": "Not implemented"})
}

func handleDeepgramKey(c *gin.Context) {
	key := os.Getenv("DEEPGRAM_API_KEY")
	c.JSON(200, gin.H{"key": key})
}
