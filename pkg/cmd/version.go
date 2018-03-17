package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display program version.",
	Run: func(cmd *cobra.Command, args []string) {
		var version = "master"
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
