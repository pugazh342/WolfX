package cmd

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("webscan %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
