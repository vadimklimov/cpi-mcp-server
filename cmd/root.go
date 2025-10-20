package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vadimklimov/cpi-mcp-server/internal/appinfo"
	"github.com/vadimklimov/cpi-mcp-server/internal/config"
	"github.com/vadimklimov/cpi-mcp-server/internal/server"
	"github.com/vadimklimov/cpi-mcp-server/internal/util/logger"
)

var cmd *cobra.Command

func NewCmd() *cobra.Command {
	cmd = &cobra.Command{
		Use:           appinfo.ID(),
		Version:       appinfo.Version(),
		Short:         appinfo.Name(),
		Long:          appinfo.FullName(),
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(_ *cobra.Command, _ []string) {
			if err := server.NewServer(); err != nil {
				log.Fatalf("MCP server failed to start: %v", err)
			}
		},
	}

	cobra.OnInitialize(
		func() {
			if err := config.Init(); err != nil {
				log.Fatalf("configuration initialization failed: %v", err)
			}
		},
		func() {
			if err := logger.Init(config.LogFile(), config.LogLevel()); err != nil {
				log.Fatalf("logger initialization failed: %v", err)
			}
		},
	)

	return cmd
}

func Execute() {
	cmd := NewCmd()
	if err := cmd.Execute(); err != nil {
		log.Fatalf("application failed to start: %v", err)
	}
}
