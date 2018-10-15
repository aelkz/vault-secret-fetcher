package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vault-agent-token-handler/handler"
)

var handlerCmd = &cobra.Command{
	Use:   "handler",
	Short: "handler retrieves credentials managed by the vault agent",
	Long:  `handler retrieves credentials managed by the vault agent`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.Start()
	},
}

func init() {
	RootCmd.AddCommand(handlerCmd)
}