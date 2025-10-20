package server

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/vadimklimov/cpi-mcp-server/internal/client"
	"github.com/vadimklimov/cpi-mcp-server/internal/util/logger"
)

type IntegrationRuntimeArtifactsGetInput struct{}

type IntegrationRuntimeArtifactsSearchInput struct {
	ID      string `json:"id" jsonschema:"Integration runtime artifact ID"`
	Version string `json:"version" jsonschema:"Version"`
	Name    string `json:"name" jsonschema:"Name"`
	Typ     string `json:"type" jsonschema:"Type"`
	Status  string `json:"status" jsonschema:"Status"`
}

type IntegrationRuntimeArtifactsOutput struct {
	Results []IntegrationRuntimeArtifact `json:"results"`
}

func getIntegrationRuntimeArtifactsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_integration_runtime_artifacts",
		Description: "Get all integration runtime artifacts",
	}
}

func searchIntegrationRuntimeArtifactsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "search_integration_runtime_artifacts",
		Description: "Search integration runtime artifacts by ID, version, name, type or status",
	}
}

func getIntegrationRuntimeArtifactsHandler(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	_ IntegrationRuntimeArtifactsGetInput,
) (
	*mcp.CallToolResult,
	IntegrationRuntimeArtifactsOutput,
	error,
) {
	logger.GetInstance().Debug("executing tool get_integration_runtime_artifacts")

	integrationRuntimeArtifacts, err := client.IntegrationRuntimeArtifacts(ctx)
	if err != nil {
		return nil, IntegrationRuntimeArtifactsOutput{}, err
	}

	logger.GetInstance().Debug(fmt.Sprintf("integration runtime artifacts found: %d", len(integrationRuntimeArtifacts)))

	results := convert(integrationRuntimeArtifacts, convertIntegrationRuntimeArtifact)

	return nil, IntegrationRuntimeArtifactsOutput{Results: results}, nil
}

func searchIntegrationRuntimeArtifactsHandler(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input IntegrationRuntimeArtifactsSearchInput,
) (
	*mcp.CallToolResult,
	IntegrationRuntimeArtifactsOutput,
	error,
) {
	logger.GetInstance().Debug("executing tool search_integration_runtime_artifacts")

	integrationRuntimeArtifacts, err := client.IntegrationRuntimeArtifacts(ctx)
	if err != nil {
		return nil, IntegrationRuntimeArtifactsOutput{}, err
	}

	logger.GetInstance().Debug(fmt.Sprintf("integration runtime artifacts found: %d", len(integrationRuntimeArtifacts)))

	matches := slices.Collect(func(yield func(client.IntegrationRuntimeArtifact) bool) {
		for _, integrationRuntimeArtifact := range integrationRuntimeArtifacts {
			if input.ID != "" && (integrationRuntimeArtifact.ID == "" ||
				!strings.Contains(strings.ToLower(integrationRuntimeArtifact.ID), strings.ToLower(input.ID))) {
				continue
			}

			if input.Version != "" && (integrationRuntimeArtifact.Version == "" ||
				!strings.Contains(strings.ToLower(integrationRuntimeArtifact.Version), strings.ToLower(input.Version))) {
				continue
			}

			if input.Name != "" && (integrationRuntimeArtifact.Name == "" ||
				!strings.Contains(strings.ToLower(integrationRuntimeArtifact.Name), strings.ToLower(input.Name))) {
				continue
			}

			if input.Typ != "" && (integrationRuntimeArtifact.Type == "" ||
				!strings.Contains(strings.ToLower(integrationRuntimeArtifact.Type), strings.ToLower(input.Typ))) {
				continue
			}

			if input.Status != "" && (integrationRuntimeArtifact.Status == "" ||
				!strings.Contains(strings.ToLower(integrationRuntimeArtifact.Status), strings.ToLower(input.Status))) {
				continue
			}

			if !yield(integrationRuntimeArtifact) {
				return
			}
		}
	})

	logger.GetInstance().Debug(fmt.Sprintf("matches found: %d", len(matches)))

	results := convert(matches, convertIntegrationRuntimeArtifact)

	return nil, IntegrationRuntimeArtifactsOutput{Results: results}, nil
}
