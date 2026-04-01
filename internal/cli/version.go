package cli

import (
	"fmt"

	"github.com/jopnelli/billtop/internal/version"
	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print billtop version information",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "billtop %s\n", version.Version)
			fmt.Fprintf(cmd.OutOrStdout(), "  commit: %s\n", version.Commit)
			fmt.Fprintf(cmd.OutOrStdout(), "  built:  %s\n", version.Date)
		},
	}
}
