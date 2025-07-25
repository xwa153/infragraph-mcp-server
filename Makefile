SHELL := /usr/bin/env bash -euo pipefail -c

# Infragraph MCP Server Makefile

# Variables
BINARY_NAME=infragraph-mcp-server
BUILD_DIR=server/bin
SOURCE_DIR=./server/
GO_FILES=$(shell find $(SOURCE_DIR) -name "*.go" -type f)

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build: $(BUILD_DIR)/$(BINARY_NAME)

$(BUILD_DIR)/$(BINARY_NAME): $(GO_FILES)
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -v -o $(BUILD_DIR)/$(BINARY_NAME) ./server
	@echo "✅ Build completed successfully: $(BUILD_DIR)/$(BINARY_NAME)"

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)/$(BINARY_NAME)
	go clean

# Format Go code
.PHONY: fmt
fmt:
	@echo "Formatting Go code..."
	go fmt ./...

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test ./...

# Lint code
.PHONY: lint
lint:
	@echo "Linting code..."
	golangci-lint run

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Run the server
.PHONY: run
run: build
	@echo "Starting MCP server..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build    - Build the MCP server binary"
	@echo "  clean    - Remove build artifacts"
	@echo "  fmt      - Format Go code"
	@echo "  test     - Run tests"
	@echo "  lint     - Lint code (requires golangci-lint)"
	@echo "  deps     - Install and tidy dependencies"
	@echo "  run      - Build and run the server"
	@echo "  help     - Show this help message"
