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

type ValueMappingsGetInput struct{}

type ValueMappingsSearchInput struct {
	ID      string `json:"id" jsonschema:"Value mapping ID"`
	Version string `json:"version" jsonschema:"Version"`
	Name    string `json:"name" jsonschema:"Name"`
}

type ValueMappingsOutput struct {
	Results []ValueMapping `json:"results"`
}

func getValueMappingsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_value_mappings",
		Description: "Get all value mappings",
	}
}

func searchValueMappingsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "search_value_mappings",
		Description: "Search value mappings by ID, version or name",
	}
}

func getValueMappingsHandler(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	_ ValueMappingsGetInput,
) (
	*mcp.CallToolResult,
	ValueMappingsOutput,
	error,
) {
	logger.GetInstance().Debug("executing tool get_value_mappings")

	valueMappings, err := client.ValueMappings(ctx)
	if err != nil {
		return nil, ValueMappingsOutput{}, err
	}

	logger.GetInstance().Debug(fmt.Sprintf("value mappings found: %d", len(valueMappings)))

	results := convert(valueMappings, convertValueMapping)

	return nil, ValueMappingsOutput{Results: results}, nil
}

func searchValueMappingsHandler(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input ValueMappingsSearchInput,
) (
	*mcp.CallToolResult,
	ValueMappingsOutput,
	error,
) {
	logger.GetInstance().Debug("executing tool search_value_mappings")

	valueMappings, err := client.ValueMappings(ctx)
	if err != nil {
		return nil, ValueMappingsOutput{}, err
	}

	logger.GetInstance().Debug(fmt.Sprintf("value mappings found: %d", len(valueMappings)))

	matches := slices.Collect(func(yield func(client.ValueMapping) bool) {
		for _, valueMapping := range valueMappings {
			if input.ID != "" && (valueMapping.ID == "" ||
				!strings.Contains(strings.ToLower(valueMapping.ID), strings.ToLower(input.ID))) {
				continue
			}

			if input.Version != "" && (valueMapping.Version == "" ||
				!strings.Contains(strings.ToLower(valueMapping.Version), strings.ToLower(input.Version))) {
				continue
			}

			if input.Name != "" && (valueMapping.Name == "" ||
				!strings.Contains(strings.ToLower(valueMapping.Name), strings.ToLower(input.Name))) {
				continue
			}

			if !yield(valueMapping) {
				return
			}
		}
	})

	logger.GetInstance().Debug(fmt.Sprintf("matches found: %d", len(matches)))

	results := convert(matches, convertValueMapping)

	return nil, ValueMappingsOutput{Results: results}, nil
}
