package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
)

// InitHTTPClient initializes a single HTTP client to be shared across all tools
func InitHTTPClient() *http.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient = cleanhttp.DefaultClient()
	retryClient.HTTPClient.Timeout = 10 * time.Second
	retryClient.RetryMax = 3
	return retryClient.StandardClient()
}

// MakeInfragraphRequest abstracts HTTP calls to the Infragraph service for reuse by other tools
func MakeInfragraphRequest(client *http.Client, method, uri string, body []byte) ([]byte, error) {
	var req *http.Request
	var err error

	if body != nil {
		req, err = http.NewRequest(method, uri, bytes.NewBuffer(body))
	} else {
		req, err = http.NewRequest(method, uri, nil)
	}

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set appropriate headers
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to infragraph: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			// Log the error if needed, but don't override the main error
		}
	}()

	if resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("infragraph service error (status %d): %s", resp.StatusCode, string(responseBody))
	}

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return responseBody, nil
}

// BuildInfragraphURI constructs the URI for infragraph API endpoints
func BuildInfragraphURI(orgID, endpoint string) string {
	return fmt.Sprintf("http://localhost:28081/infragraph/2025-05-07/organizations/%s/%s", orgID, endpoint)
}

// ReadQueryTemplate reads a query template file from the resources directory
func ReadQueryTemplate(filePath string) (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}
	exeDir := filepath.Dir(exePath)
	// Assume resources/ is in the project root, one level up from bin/
	absPath := filepath.Join(exeDir, "..", filePath)

	content, err := os.ReadFile(absPath)
	if err != nil {
		return "", fmt.Errorf("failed to read template file: %w", err)
	}
	return string(content), nil
}
