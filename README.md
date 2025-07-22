# Infragraph MCP Server

The Infragraph MCP Server is a [Model Context Protocol (MCP)](https://modelcontextprotocol.io/introduction)
server that provides integration with Infragraph APIs, enabling fetching information swiftly
through natural language in LLMs.

## Installation

### Usage with VS Code

Add the following JSON block to your User Settings (JSON) file in VS Code. You can do this by pressing `Ctrl + Shift + P` and typing `Preferences: Open User Settings (JSON)`. 

Compile the server by running `make build`.

More about using MCP server tools in VS Code's [agent mode documentation](https://code.visualstudio.com/docs/copilot/chat/mcp-servers).
Remember to update the path to the binary.

```json
{
  "servers": {
    "Infragraph": {
      "command": "path/to/infragraph-mcp-server/cmd/infragraph-mcp-server"
    }
  }
}
```

Optionally, you can add a similar example (i.e. without the mcp key) to a file called `.vscode/mcp.json` in your workspace. This will allow you to share the configuration with others.

```json
{
  "servers": {
    "Infragraph": {
      "command": "path/to/infragraph-mcp-server/cmd/infragraph-mcp-server"
    }
  }
}
```

## Tool Configuration

### Available Toolsets

The following sets of tools are available:

| Toolset | Tool | Description |
|---------|------|-------------|
| `connections` | `ListConnections` | List connections in specific org |

## Development

### Prerequisites
- Go (check [go.mod](./go.mod) file for specific version)

### Available Make Commands

| Command | Description |
|---------|-------------|
| `make build` | Build the binary |

## Contributing

1. Fork the repository
2. Create your feature branch
3. Make your changes
4. Run tests
5. Submit a pull request

## License

This project is licensed under the terms of the MPL-2.0 open source license. Please refer to [LICENSE](./LICENSE) file for the full terms.

## Support

For bug reports and feature requests, please open an issue on GitHub.

For general questions and discussions, open a GitHub Discussion.