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

type IntegrationPackagesGetInput struct{}

type IntegrationPackagesSearchInput struct {
	ID      string `json:"id" jsonschema:"Integration package ID"`
	Version string `json:"version" jsonschema:"Version"`
	Name    string `json:"name" jsonschema:"Name"`
	Vendor  string `json:"vendor" jsonschema:"Vendor"`
	Mode    string `json:"mode" jsonschema:"Mode"`
}

type IntegrationPackagesOutput struct {
	Results []IntegrationPackage `json:"results"`
}

func getIntegrationPackagesTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_integration_packages",
		Description: "Get all integration packages",
	}
}

func searchIntegrationPackagesTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "search_integration_packages",
		Description: "Search integration packages by ID, version, name, vendor or mode",
	}
}

func getIntegrationPackagesHandler(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	_ IntegrationPackagesGetInput,
) (
	*mcp.CallToolResult,
	IntegrationPackagesOutput,
	error,
) {
	logger.GetInstance().Debug("executing tool get_integration_packages")

	integrationPackages, err := client.IntegrationPackages(ctx)
	if err != nil {
		return nil, IntegrationPackagesOutput{}, err
	}

	logger.GetInstance().Debug(fmt.Sprintf("integration packages found: %d", len(integrationPackages)))

	results := convert(integrationPackages, convertIntegrationPackage)

	return nil, IntegrationPackagesOutput{Results: results}, nil
}

func searchIntegrationPackagesHandler(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input IntegrationPackagesSearchInput,
) (
	*mcp.CallToolResult,
	IntegrationPackagesOutput,
	error,
) {
	logger.GetInstance().Debug("executing tool search_integration_packages")

	integrationPackages, err := client.IntegrationPackages(ctx)
	if err != nil {
		return nil, IntegrationPackagesOutput{}, err
	}

	logger.GetInstance().Debug(fmt.Sprintf("integration packages found: %d", len(integrationPackages)))

	matches := slices.Collect(func(yield func(client.IntegrationPackage) bool) {
		for _, integrationPackage := range integrationPackages {
			if input.ID != "" && (integrationPackage.ID == "" ||
				!strings.Contains(strings.ToLower(integrationPackage.ID), strings.ToLower(input.ID))) {
				continue
			}

			if input.Version != "" && (integrationPackage.Version == "" ||
				!strings.Contains(strings.ToLower(integrationPackage.Version), strings.ToLower(input.Version))) {
				continue
			}

			if input.Name != "" && (integrationPackage.Name == "" ||
				!strings.Contains(strings.ToLower(integrationPackage.Name), strings.ToLower(input.Name))) {
				continue
			}

			if input.Vendor != "" && (integrationPackage.Vendor == "" ||
				!strings.Contains(strings.ToLower(integrationPackage.Vendor), strings.ToLower(input.Vendor))) {
				continue
			}

			if input.Mode != "" && (integrationPackage.Mode == "" ||
				!strings.Contains(strings.ToLower(integrationPackage.Mode), strings.ToLower(input.Mode))) {
				continue
			}

			if !yield(integrationPackage) {
				return
			}
		}
	})

	logger.GetInstance().Debug(fmt.Sprintf("matches found: %d", len(matches)))

	results := convert(matches, convertIntegrationPackage)

	return nil, IntegrationPackagesOutput{Results: results}, nil
}
