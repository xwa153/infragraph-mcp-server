package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Demo 🚀",
		"0.0.1",
		server.WithToolCapabilities(true),
	)

	// Add tool
	registryClient := InitRegistryClient()
	s.AddTool(ListConnections(registryClient))

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

// ListConnections lists all connections in an organization.
func ListConnections(registryClient *http.Client) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("listConnections",
			mcp.WithDescription(`This tool lists all connections in one organization.`),
			mcp.WithTitleAnnotation("Fetch Connections"),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithString("orgID", mcp.Required(), mcp.Description("The organization ID to list connections for")),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

			connectionDetails, err := resolveConnectionDetails(request, registryClient)
			if err != nil {
				return nil, err
			}

			return mcp.NewToolResultText(connectionDetails), nil
		}
}

func resolveConnectionDetails(request mcp.CallToolRequest, registryClient *http.Client) (string, error) {
	orgID := request.GetString("orgID", "")
	if orgID == "" {
		return "", fmt.Errorf("orgID is required and must be a string")
	}

	var err error
	uri := fmt.Sprintf("http://localhost:28081/infragraph/2025-05-07/organizations/%s/connections", orgID)
	jsonData, err := sendRegistryCall(registryClient, "GET", uri)
	if err != nil {
		return "", fmt.Errorf("error fetching connections for org %s: %w", orgID, err)
	}

	return string(jsonData), nil
}

func sendRegistryCall(client *http.Client, method string, uri string) ([]byte, error) {
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", "404 Not Found")
	}

	defer resp.Body.Close()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func InitRegistryClient() *http.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient = cleanhttp.DefaultClient()
	retryClient.HTTPClient.Timeout = 10 * time.Second
	retryClient.RetryMax = 3
	return retryClient.StandardClient()
}
