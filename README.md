# Simple MCP Server

A simple Model Context Protocol (MCP) server implementation in Go.

This project provides two MCP server implementations:

- **Official SDK** (`cmd/go-sdk`) - Using [modelcontextprotocol/go-sdk](https://github.com/modelcontextprotocol/go-sdk)
- **Third-party SDK** (`cmd/mark3labs`) - Using [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go)

## What is MCP?

MCP (Model Context Protocol) is a standard protocol that enables AI assistants to use external tools and resources.
This server communicates via JSON-RPC messages over stdio, using process-based communication without host/port like HTTP servers.

## Setup

### 1. Install Dependencies & Build

```bash
make all
```

This command will:

- Install development tools
- Initialize Go modules
- Tidy dependencies and create vendor directory
- Build the server

### 2. Run Server

```bash
make start
```

Builds and runs the server.

### 3. Configure Cursor

Create or edit `~/.cursor/mcp_config.json`:

```json
{
  "mcpServers": {
    "simple-mcp": {
      "command": "/Users/WonjinSin/Documents/project/simple-mcp/bin/mark3labs",
      "args": []
    }
  }
}
```

**Development mode** (auto-reload after code changes):

```json
{
  "mcpServers": {
    "simple-mcp": {
      "command": "go",
      "args": ["run", "/Users/WonjinSin/Documents/project/simple-mcp/cmd/mark3labs/main.go"],
      "cwd": "/Users/WonjinSin/Documents/project/simple-mcp"
    }
  }
}
```

### 4. Restart Cursor

Completely quit and restart Cursor to activate the MCP server.

## Available Tools

- **hello_world**: A simple example tool that takes a name and returns a greeting

## Resources

- [MCP Specification](https://spec.modelcontextprotocol.io/)
- [go-sdk](https://github.com/modelcontextprotocol/go-sdk)
- [mcp-go](https://github.com/mark3labs/mcp-go)
