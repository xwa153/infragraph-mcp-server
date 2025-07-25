# Infragraph MCP Server

An MCP (Model Context Protocol) server that provides tools to query the Infragraph service for infrastructure graph data.

## Overview

This MCP server enables users to query infrastructure graphs through the Infragraph service, allowing you to explore relationships between virtual machines, containers, databases, and other infrastructure components.

## Features

- **QueryInfragraph**: Query the Infragraph service with complex graph queries
- **ListConnections**: List all connections in an organization
- Query templates and examples in the `server/resources/` directory

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd infragraph-mcp-server
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the server:
```bash
make build
```

## Usage

### Starting the Server

```bash
./server/bin/infragraph-mcp-server
```

### Available Tools

#### QueryInfragraph

Query the Infragraph service with complex graph queries.

**Parameters:**
- `org_id` (required): The organization ID to query
- `query` (required): JSON string representing the graph query structure

#### ListConnections

List all connections in an organization.

**Parameters:**
- `org_id` (required): The organization ID to list connections for

## Query Templates

The `server/resources/` directory contains pre-built query templates:

- `virtual_machines.json`: Find all virtual machines
- `vm_with_images.json`: Find VMs and their associated images

## API Endpoints

The server connects to the Infragraph service using the following endpoints:

### List Connections
- **URL**: `http://localhost:28081/infragraph/2025-05-07/organizations/{org_id}/connections`
- **Method**: GET
- **Description**: Retrieves all connections for the specified organization

### Query Infragraph
- **URL**: `http://localhost:28081/infragraph/2025-05-07/organizations/{org_id}/query`
- **Method**: POST
- **Request Body**: JSON object with a `query` field containing the graph query
- **Description**: Executes complex graph queries against the infrastructure data
