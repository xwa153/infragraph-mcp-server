package resources

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"infragraph-mcp-server/server/utils"
)

// QueryTemplateResource creates a resource for query templates
func QueryTemplateResource(uri, filePath string) (resource mcp.Resource, handler server.ResourceHandlerFunc) {
	return mcp.NewResource(uri, uri,
			mcp.WithResourceDescription("Query template for Infragraph service"),
			mcp.WithMIMEType("application/json"),
		),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Read the template file
			content, err := utils.ReadQueryTemplate(filePath)
			if err != nil {
				return nil, fmt.Errorf("error reading query template from %s: %w", filePath, err)
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      uri,
					MIMEType: "application/json",
					Text:     content,
				},
			}, nil
		}
}
