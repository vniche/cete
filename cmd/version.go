package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vniche/cete/version"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Long:  "Print the version number",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("cete version: %s\n", version.Version)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
