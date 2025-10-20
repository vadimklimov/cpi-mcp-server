package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/vadimklimov/cpi-mcp-server/internal/appinfo"
	"github.com/vadimklimov/cpi-mcp-server/internal/config"
)

func NewServer() error {
	server := mcp.NewServer(&mcp.Implementation{Name: appinfo.ID(), Version: appinfo.Version()}, nil)

	mcp.AddTool(server, getIntegrationPackagesTool(), getIntegrationPackagesHandler)
	mcp.AddTool(server, searchIntegrationPackagesTool(), searchIntegrationPackagesHandler)
	mcp.AddTool(server, getIntegrationFlowsTool(), getIntegrationFlowsHandler)
	mcp.AddTool(server, searchIntegrationFlowsTool(), searchIntegrationFlowsHandler)
	mcp.AddTool(server, getValueMappingsTool(), getValueMappingsHandler)
	mcp.AddTool(server, searchValueMappingsTool(), searchValueMappingsHandler)
	mcp.AddTool(server, getMessageMappingsTool(), getMessageMappingsHandler)
	mcp.AddTool(server, searchMessageMappingsTool(), searchMessageMappingsHandler)
	mcp.AddTool(server, getScriptCollectionsTool(), getScriptCollectionsHandler)
	mcp.AddTool(server, searchScriptCollectionsTool(), searchScriptCollectionsHandler)
	mcp.AddTool(server, getIntegrationRuntimeArtifactsTool(), getIntegrationRuntimeArtifactsHandler)
	mcp.AddTool(server, searchIntegrationRuntimeArtifactsTool(), searchIntegrationRuntimeArtifactsHandler)

	switch config.ServerTransport() {
	case config.TransportStdio:
		return server.Run(context.Background(), &mcp.StdioTransport{})
	case config.TransportHTTP:
		addr := fmt.Sprintf(":%s", config.ServerPort())

		handler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
			return server
		}, nil)

		return http.ListenAndServe(addr, handler)
	}

	return nil
}
