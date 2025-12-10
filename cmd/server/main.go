package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/wonjinsin/simple-mcp/internal/config"
	"github.com/wonjinsin/simple-mcp/internal/database"
	langchain "github.com/wonjinsin/simple-mcp/internal/repository/langchain/ollama"
	"github.com/wonjinsin/simple-mcp/internal/usecase"
	"github.com/wonjinsin/simple-mcp/pkg/logger"
)

func main() {
	// Print ASCII art banner
	printBanner()

	// Set timezone to UTC for the entire program
	time.Local = time.UTC

	// Load configuration
	cfg := config.Load()

	// Initialize logger
	logger.Initialize(cfg.Env)

	// Initialize LLM
	ollamaLLM, err := database.NewOllamaLLM()
	if err != nil {
		log.Fatalf("failed to initialize LLM: %v", err)
	}

	// Initialize database connection
	basicChatRepo := langchain.NewBasicChatRepo(ollamaLLM)

	// Wiring (Composition Root)
	_ = usecase.NewBasicChatService(basicChatRepo)

	// Create a new MCP server
	s := server.NewMCPServer(
		"Demo 🚀",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	// Add tool
	tool := mcp.NewTool("hello_world",
		mcp.WithDescription("Say hello to someone"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the person to greet"),
		),
	)

	// Add tool handler
	s.AddTool(tool, helloHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

	log.Println("bye")
}

func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}

func printBanner() {
	// Read banner from file
	bannerPath := "internal/config/banner.asc"
	bannerBytes, err := os.ReadFile(bannerPath)
	if err != nil {
		log.Printf("warning: could not read banner file: %v", err)
		return
	}

	log.Println(string(bannerBytes))
}
