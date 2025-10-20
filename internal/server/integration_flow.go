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

type IntegrationFlowsGetInput struct{}

type IntegrationFlowsSearchInput struct {
	ID      string `json:"id" jsonschema:"Integration flow ID"`
	Version string `json:"version" jsonschema:"Version"`
	Name    string `json:"name" jsonschema:"Name"`
}

type IntegrationFlowsOutput struct {
	Results []IntegrationFlow `json:"results"`
}

func getIntegrationFlowsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_integration_flows",
		Description: "Get all integration flows",
	}
}

func searchIntegrationFlowsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "search_integration_flows",
		Description: "Search integration flows by ID, version or name",
	}
}

func getIntegrationFlowsHandler(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	_ IntegrationFlowsGetInput,
) (
	*mcp.CallToolResult,
	IntegrationFlowsOutput,
	error,
) {
	logger.GetInstance().Debug("executing tool get_integration_flows")

	integrationFlows, err := client.IntegrationFlows(ctx)
	if err != nil {
		return nil, IntegrationFlowsOutput{}, err
	}

	logger.GetInstance().Debug(fmt.Sprintf("integration flows found: %d", len(integrationFlows)))

	results := convert(integrationFlows, convertIntegrationFlow)

	return nil, IntegrationFlowsOutput{Results: results}, nil
}

func searchIntegrationFlowsHandler(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input IntegrationFlowsSearchInput,
) (
	*mcp.CallToolResult,
	IntegrationFlowsOutput,
	error,
) {
	logger.GetInstance().Debug("executing tool search_integration_flows")

	integrationFlows, err := client.IntegrationFlows(ctx)
	if err != nil {
		return nil, IntegrationFlowsOutput{}, err
	}

	logger.GetInstance().Debug(fmt.Sprintf("integration flows found: %d", len(integrationFlows)))

	matches := slices.Collect(func(yield func(client.IntegrationFlow) bool) {
		for _, integrationFlow := range integrationFlows {
			if input.ID != "" && (integrationFlow.ID == "" ||
				!strings.Contains(strings.ToLower(integrationFlow.ID), strings.ToLower(input.ID))) {
				continue
			}

			if input.Version != "" && (integrationFlow.Version == "" ||
				!strings.Contains(strings.ToLower(integrationFlow.Version), strings.ToLower(input.Version))) {
				continue
			}

			if input.Name != "" && (integrationFlow.Name == "" ||
				!strings.Contains(strings.ToLower(integrationFlow.Name), strings.ToLower(input.Name))) {
				continue
			}

			if !yield(integrationFlow) {
				return
			}
		}
	})

	logger.GetInstance().Debug(fmt.Sprintf("matches found: %d", len(matches)))

	results := convert(matches, convertIntegrationFlow)

	return nil, IntegrationFlowsOutput{Results: results}, nil
}
