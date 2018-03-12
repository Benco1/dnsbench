package cmd

import (
	"github.com/spf13/cobra"
)

// localCmd represents the local system resolver benchmark command
var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Benchmark the local system resolver configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		resolveLocally := true
		benchmark(&resolveLocally, &concurrency, &count, &interval, &names, &qps)
	},
}

func init() {
	runCmd.AddCommand(localCmd)
}
