package cmd

import (
	"github.com/spf13/cobra"
)

func searchCmdRun(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("author") {
		// todo
	}
	if cmd.Flags().Changed("title") {
		// todo
	}
	println(cmd.Name())
}
