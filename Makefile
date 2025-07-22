SHELL := /usr/bin/env bash -euo pipefail -c

BINARY_NAME ?= infragraph-mcp-server
GO=go
# Build flags
LDFLAGS=-ldflags="-s -w -X infragraph-mcp-server/version.GitCommit=$(shell git rev-parse HEAD) -X infragraph-mcp-server/version.BuildDate=$(shell git show --no-show-signature -s --format=%cd --date=format:"%Y-%m-%dT%H:%M:%SZ" HEAD)"

# Build the binary
ARCH     = $(shell A=$$(uname -m); [ $$A = x86_64 ] && A=amd64; echo $$A)
OS       = $(shell uname | tr [[:upper:]] [[:lower:]])

.PHONY: all build
# Default target
all: build

build:
	CGO_ENABLED=0 GOARCH=$(ARCH) GOOS=$(OS) $(GO) build $(LDFLAGS) -o ./cmd/$(BINARY_NAME) ./main.go
