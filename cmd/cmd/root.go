package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gbooks",
	Short: "search google books",
	Long:  `search google books and pick results in a tui`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringSliceP("ext", "e", []string{".yaml", ".ini"}, "extension for metadata")
	viper.BindPFlag("ext", rootCmd.PersistentFlags().Lookup("ext"))
	rootCmd.PersistentFlags().Bool("dont-save", false, "write meta to stdout")
	viper.BindPFlag("dont-save", rootCmd.PersistentFlags().Lookup("dont-save"))
}
