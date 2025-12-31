package main

import (
	"context"
	"log"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Input defines the input parameters for the hello_world tool
type Input struct {
	Name string `json:"name" jsonschema:"Name of the person to greet"`
}

// Output defines the output of the hello_world tool
type Output struct {
	Greeting string `json:"greeting" jsonschema:"The greeting message"`
}

// SayHello is the handler function for the hello_world tool
func SayHello(ctx context.Context, req *mcp.CallToolRequest, input Input) (
	*mcp.CallToolResult,
	Output,
	error,
) {
	greeting := "Hello, " + input.Name + "!"
	return nil, Output{Greeting: greeting}, nil
}

func main() {
	// Set timezone to UTC for the entire program
	time.Local = time.UTC

	// Create a new MCP server with implementation details
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "Simple MCP Server",
			Version: "1.0.0",
		},
		nil, // No custom features
	)

	// Add the hello_world tool
	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "hello_world",
			Description: "Say hello to someone",
		},
		SayHello,
	)

	// Run the server over stdin/stdout
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
