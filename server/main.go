package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"

	"infragraph-mcp-server/server/resources"
	"infragraph-mcp-server/server/tools"
	"infragraph-mcp-server/server/utils"
)

func main() {
	// Create a new MCP server with both tool and resource capabilities
	s := server.NewMCPServer(
		"Infragraph MCP Server 🚀",
		"0.0.1",
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, true),
	)

	// Initialize single HTTP client for all requests
	httpClient := utils.InitHTTPClient()

	// Add tools
	s.AddTool(tools.ListConnections(httpClient))
	s.AddTool(tools.QueryInfragraph(httpClient))

	// Add resources for query templates
	s.AddResource(resources.QueryTemplateResource("resource://resources/virtual-machines", "resources/virtual_machines.json"))
	s.AddResource(resources.QueryTemplateResource("resource://resources/vm-with-images", "resources/vm_with_images.json"))

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error is: %v\n", err)
	}
}
