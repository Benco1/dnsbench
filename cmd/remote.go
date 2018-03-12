package cmd

import (
	"github.com/spf13/cobra"
)

var transport string

// remoteCmd represents the remote nameserver benchmark command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Benchmark a remote nameserver.",
	Run: func(cmd *cobra.Command, args []string) {
		resolveLocally := false
		benchmark(&resolveLocally, &concurrency, &count, &interval, &names, &qps)
	},
}

func init() {
	runCmd.AddCommand(remoteCmd)
	remoteCmd.Flags().StringVarP(&transport, "transport", "t", "udp", "query transport protocol")
}
