package cmd

import (
	"github.com/ohzqq/gbooks"
	"github.com/spf13/cobra"
)

func searchCmdRun(cmd *cobra.Command, args []string) {
	req := gbooks.NewRequest()
	println(req.String())
}
