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

type MessageMappingsGetInput struct{}

type MessageMappingsSearchInput struct {
	ID      string `json:"id" jsonschema:"Message mapping ID"`
	Version string `json:"version" jsonschema:"Version"`
	Name    string `json:"name" jsonschema:"Name"`
}

type MessageMappingsOutput struct {
	Results []MessageMapping `json:"results"`
}

func getMessageMappingsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_message_mappings",
		Description: "Get all message mappings",
	}
}

func searchMessageMappingsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "search_message_mappings",
		Description: "Search message mappings by ID, version or name",
	}
}

func getMessageMappingsHandler(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	_ MessageMappingsGetInput,
) (
	*mcp.CallToolResult,
	MessageMappingsOutput,
	error,
) {
	logger.GetInstance().Debug("executing tool get_message_mappings")

	messageMappings, err := client.MessageMappings(ctx)
	if err != nil {
		return nil, MessageMappingsOutput{}, err
	}

	logger.GetInstance().Debug(fmt.Sprintf("message mappings found: %d", len(messageMappings)))

	results := convert(messageMappings, convertMessageMapping)

	return nil, MessageMappingsOutput{Results: results}, nil
}

func searchMessageMappingsHandler(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input MessageMappingsSearchInput,
) (
	*mcp.CallToolResult,
	MessageMappingsOutput,
	error,
) {
	logger.GetInstance().Debug("executing tool search_message_mappings")

	messageMappings, err := client.MessageMappings(ctx)
	if err != nil {
		return nil, MessageMappingsOutput{}, err
	}

	logger.GetInstance().Debug(fmt.Sprintf("message mappings found: %d", len(messageMappings)))

	matches := slices.Collect(func(yield func(client.MessageMapping) bool) {
		for _, messageMapping := range messageMappings {
			if input.ID != "" && (messageMapping.ID == "" ||
				!strings.Contains(strings.ToLower(messageMapping.ID), strings.ToLower(input.ID))) {
				continue
			}

			if input.Version != "" && (messageMapping.Version == "" ||
				!strings.Contains(strings.ToLower(messageMapping.Version), strings.ToLower(input.Version))) {
				continue
			}

			if input.Name != "" && (messageMapping.Name == "" ||
				!strings.Contains(strings.ToLower(messageMapping.Name), strings.ToLower(input.Name))) {
				continue
			}

			if !yield(messageMapping) {
				return
			}
		}
	})

	logger.GetInstance().Debug(fmt.Sprintf("matches found: %d", len(matches)))

	results := convert(matches, convertMessageMapping)

	return nil, MessageMappingsOutput{Results: results}, nil
}
