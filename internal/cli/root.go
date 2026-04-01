// Package cli defines the billtop command-line interface.
package cli

import (
	"github.com/spf13/cobra"
)

// NewRootCmd creates the top-level billtop command and registers all subcommands.
func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "billtop",
		Short: "Open source GCP cost optimization",
		Long:  "billtop connects to your GCP billing export in BigQuery, shows where your money goes, and tells you how to spend less.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.AddCommand(newVersionCmd())

	return root
}
