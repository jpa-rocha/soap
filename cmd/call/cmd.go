package call

import (
	"log/slog"

	"soap/internal/call"

	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "call",
		Short: "call soap",
		RunE:  call.Cmd,
		Args:  cobra.NoArgs,
	}

	slog.Debug("command call:", "value", cmd.Use)

	return &cmd
}
