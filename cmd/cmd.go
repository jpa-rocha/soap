// Package cmd sets up the main cmd and gathers all subcommands
package cmd

import (
	"context"
	"fmt"

	config "soap/internal/config"
	logger "soap/internal/logger"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	ctx := context.Background()
	cmd := cobra.Command{
		Use:   "soap",
		Short: "soap test",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			logger.Init()
			config.InitConfig(ctx)
		},
	}
	cmd.SetContext(ctx)
	ctx, cancel := context.WithTimeout(cmd.Context(), config.HTTPTimeout)

	defer cancel()

	return &cmd
}

func Execute() error {
	ctx := RootCmd().Context()
	if err := RootCmd().ExecuteContext(ctx); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
