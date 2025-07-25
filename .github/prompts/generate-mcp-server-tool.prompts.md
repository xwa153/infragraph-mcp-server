---
mode: agent
model: Claude Sonnet 4
description: 'Add a tool to the MCP server that allows users to query infragraph service.'
---
# Goal MCP Server Tools for Infragraph Service

Your goal is to build tools to the MCP server that allows users to
- list connections
- query the Infragraph service
- provide resources for query templates

## Requirements

### General Requirements

1. The MCP server should use a single HTTP client for all requests to the Infragraph service.
1. The MCP server should abstract the HTTP call into a function that can be reused by other tools in the future.

### List Connections Tool

1. The endpoint for the Infragraph service is `http://localhost:28081/infragraph/2025-05-07/organizations/{org_id}/connections`, the method is `GET`.
1. The tool should be named `ListConnections`.
1. The tool should require the following parameters:
   - `org_id`: The organization ID to query.

### Query Infragraph Tool

1. The endpoint for the Infragraph service is `http://localhost:28081/infragraph/2025-05-07/organizations/{org_id}/query`, the method is `POST`, and the request body is a JSON object with a `query` field that contains the query string.
1. The tool should be named `QueryInfragraph`.
1. The tool should require the following parameters:
   - `org_id`: The organization ID to query.
   - `query`: The query string to execute.
1. This is an example of a query request body:
```json
{
    "query": {
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
}
```
1. This is a example of a query response body:
```json
{
  "nodes": [
    {
      "id": "e9d46f2f-3588-012a-8051-42639d3b70ba",
      "label": "VIRTUAL_MACHINE",
      "name": "",
      "source": "AWS",
      "properties": {
        "CustomerID": {
          "type": "PROPERTY_TYPE_STRING",
          "value": "ce7c97b8-5b4d-423b-803d-21838d57c6ca"
        },
        "arn": {
          "type": "PROPERTY_TYPE_STRING",
          "value": "arn:aws:ec2:us-east-1:912613162329:instance/i-03aa520ca0d8c6241"
        },
        "monitoring": {
          "type": "PROPERTY_TYPE_BOOL",
          "value": false
        },
        "privateIp": {
          "type": "PROPERTY_TYPE_STRING",
          "value": "172.31.21.54"
        },
        "resourceType": {
          "type": "PROPERTY_TYPE_STRING",
          "value": "AWS_EC2_INSTANCE"
        },
        "source_id": {
          "type": "PROPERTY_TYPE_STRING",
          "value": "i-03aa520ca0d8c6241"
        },
        "tenancy": {
          "type": "PROPERTY_TYPE_STRING",
          "value": "default"
        }
      }
    },
    {
      "id": "73d4ccbf-fefa-0bc3-a4f0-b81cefc585a1",
      "label": "VIRTUAL_MACHINE_IMAGE",
      "name": "al2023-ami-2023.8.20250715.0-kernel-6.1-x86_64",
      "source": "AWS",
      "properties": {
        "CustomerID": {
          "type": "PROPERTY_TYPE_STRING",
          "value": "ce7c97b8-5b4d-423b-803d-21838d57c6ca"
        },
        "arn": {
          "type": "PROPERTY_TYPE_STRING",
          "value": ""
        },
        "resourceType": {
          "type": "PROPERTY_TYPE_STRING",
          "value": "AWS_AMI"
        },
        "source_id": {
          "type": "PROPERTY_TYPE_STRING",
          "value": "ami-0cbbe2c6a1bb2ad63"
        }
      }
    }
  ],
  "edges": [
    {
      "label": "RUNS",
      "sourceNodeId": "e9d46f2f-3588-012a-8051-42639d3b70ba",
      "destNodeId": "73d4ccbf-fefa-0bc3-a4f0-b81cefc585a1"
    }
  ],
  "meta": {
    "elapsedQueryTime": "6",
    "nodeCount": 2
  }
}
```
1. The tool should return the response from the Infragraph service as a JSON object.

### Folder Structure

1. `README.md`: A file that provides documentation for the MCP server tool, including how to use it, available parameters, and examples.
1. server source code should be placed in the `server/` directory.
1. `server/main.go`: The main entry point for the MCP server tool.
1. `server/utils/`: A utility file that contains helper functions for the MCP server tool.
1. `server/resources/`: A directory that contains query templates or other resources needed by the tool.
1. `server/tools/`: A directory that contains the tools for the MCP server.
1. `server/bin/`: A directory that contains the compiled binary for the MCP server tool.
