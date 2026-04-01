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
			w := cmd.OutOrStdout()
			_, _ = fmt.Fprintf(w, "billtop %s\n", version.Version)
			_, _ = fmt.Fprintf(w, "  commit: %s\n", version.Commit)
			_, _ = fmt.Fprintf(w, "  built:  %s\n", version.Date)
		},
	}
}
