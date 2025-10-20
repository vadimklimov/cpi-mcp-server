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

type ScriptCollectionsGetInput struct{}

type ScriptCollectionsSearchInput struct {
	ID      string `json:"id" jsonschema:"Script collection ID"`
	Version string `json:"version" jsonschema:"Version"`
	Name    string `json:"name" jsonschema:"Name"`
}

type ScriptCollectionsOutput struct {
	Results []ScriptCollection `json:"results"`
}

func getScriptCollectionsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "get_script_collections",
		Description: "Get all script collections",
	}
}

func searchScriptCollectionsTool() *mcp.Tool {
	return &mcp.Tool{
		Name:        "search_script_collections",
		Description: "Search script collections by ID, version or name",
	}
}

func getScriptCollectionsHandler(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	_ ScriptCollectionsGetInput,
) (
	*mcp.CallToolResult,
	ScriptCollectionsOutput,
	error,
) {
	logger.GetInstance().Debug("executing tool get_script_collections")

	scriptCollections, err := client.ScriptCollections(ctx)
	if err != nil {
		return nil, ScriptCollectionsOutput{}, err
	}

	logger.GetInstance().Debug(fmt.Sprintf("script collections found: %d", len(scriptCollections)))

	results := convert(scriptCollections, convertScriptCollection)

	return nil, ScriptCollectionsOutput{Results: results}, nil
}

func searchScriptCollectionsHandler(
	ctx context.Context,
	_ *mcp.CallToolRequest,
	input ScriptCollectionsSearchInput,
) (
	*mcp.CallToolResult,
	ScriptCollectionsOutput,
	error,
) {
	logger.GetInstance().Debug("executing tool search_script_collections")

	scriptCollections, err := client.ScriptCollections(ctx)
	if err != nil {
		return nil, ScriptCollectionsOutput{}, err
	}

	logger.GetInstance().Debug(fmt.Sprintf("script collections found: %d", len(scriptCollections)))

	matches := slices.Collect(func(yield func(client.ScriptCollection) bool) {
		for _, scriptCollection := range scriptCollections {
			if input.ID != "" && (scriptCollection.ID == "" ||
				!strings.Contains(strings.ToLower(scriptCollection.ID), strings.ToLower(input.ID))) {
				continue
			}

			if input.Version != "" && (scriptCollection.Version == "" ||
				!strings.Contains(strings.ToLower(scriptCollection.Version), strings.ToLower(input.Version))) {
				continue
			}

			if input.Name != "" && (scriptCollection.Name == "" ||
				!strings.Contains(strings.ToLower(scriptCollection.Name), strings.ToLower(input.Name))) {
				continue
			}

			if !yield(scriptCollection) {
				return
			}
		}
	})

	logger.GetInstance().Debug(fmt.Sprintf("matches found: %d", len(matches)))

	results := convert(matches, convertScriptCollection)

	return nil, ScriptCollectionsOutput{Results: results}, nil
}
