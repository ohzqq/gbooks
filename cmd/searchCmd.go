package cmd

import (
	"github.com/spf13/cobra"
)

func searchCmdRun(cmd *cobra.Command, args []string) {
	println(cmd.Name())
}
