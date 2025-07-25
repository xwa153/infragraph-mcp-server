package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"infragraph-mcp-server/server/utils"
)

// ListConnections lists all connections in an organization
func ListConnections(httpClient *http.Client) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("ListConnections",
			mcp.WithDescription(`This tool lists all connections in one organization using the Infragraph service.`),
			mcp.WithTitleAnnotation("List Connections"),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithString("org_id", mcp.Required(), mcp.Description("The organization ID to list connections for")),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			orgID := request.GetString("org_id", "")
			if orgID == "" {
				return nil, fmt.Errorf("org_id is required and must be a string")
			}

			uri := utils.BuildInfragraphURI(orgID, "connections")
			responseBody, err := utils.MakeInfragraphRequest(httpClient, "GET", uri, nil)
			if err != nil {
				return nil, fmt.Errorf("error fetching connections for org %s: %w", orgID, err)
			}

			return mcp.NewToolResultText(string(responseBody)), nil
		}
}

// QueryInfragraph queries the infragraph service with complex graph queries
func QueryInfragraph(httpClient *http.Client) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("QueryInfragraph",
			mcp.WithDescription(`This tool allows users to query the Infragraph service with complex graph queries. 
				The query parameter must be a JSON string representing the graph query structure.
				
				Example query to find virtual machines and their images:
				{
				  "node": {
					"node_label": "VIRTUAL_MACHINE"
				  },
				  "edge": {
					"edge_label": "RUNS",
					"incoming": false,
					"node_query": {
					  "node": {
						"node_label": "VIRTUAL_MACHINE_IMAGE"
					  }
					}
				  }
				}
				
				Example query to find all virtual machines:
				{
				  "node": {
					"node_label": "VIRTUAL_MACHINE"
				  }
				}
				
				See the resources/ directory for more query templates.`),
			mcp.WithTitleAnnotation("Query Infragraph"),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithString("org_id", mcp.Required(), mcp.Description("The organization ID to query")),
			mcp.WithString("query", mcp.Required(), mcp.Description(`The query string to execute. Must be a valid JSON object representing the graph query structure.

			Common node labels:
			- VIRTUAL_MACHINE: Virtual machine instances
			- VIRTUAL_MACHINE_IMAGE: VM images/AMIs  
			- CONTAINER: Container instances
			- DATABASE: Database instances
			- STORAGE_BUCKET: Storage buckets
			
			Common edge labels:
			- RUNS: VM runs on image
			- CONNECTS_TO: Network connections
			- CONTAINS: Containment relationships
			
			Example: {"node": {"node_label": "VIRTUAL_MACHINE"}}
			
			Check the resources/ directory for query templates.`)),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			orgID := request.GetString("org_id", "")
			queryStr := request.GetString("query", "")

			if orgID == "" || queryStr == "" {
				return nil, fmt.Errorf("org_id and query are required and must be strings")
			}

			// Parse the query string as JSON to validate it
			var queryJSON interface{}
			if err := json.Unmarshal([]byte(queryStr), &queryJSON); err != nil {
				return nil, fmt.Errorf("query must be valid JSON. Example: {\"node\": {\"node_label\": \"VIRTUAL_MACHINE\"}}. Error: %w", err)
			}

			// Create the request body with the query field as specified in the API
			requestBody := map[string]interface{}{
				"query": queryJSON,
			}

			jsonData, err := json.Marshal(requestBody)
			if err != nil {
				return nil, fmt.Errorf("error marshaling request body: %w", err)
			}

			// Make the POST request to infragraph service
			uri := utils.BuildInfragraphURI(orgID, "query")
			responseBody, err := utils.MakeInfragraphRequest(httpClient, "POST", uri, jsonData)
			if err != nil {
				return nil, fmt.Errorf("error querying infragraph for org %s: %w", orgID, err)
			}

			return mcp.NewToolResultText(string(responseBody)), nil
		}
}
