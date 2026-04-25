package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	verbose bool
	version = "dev"
)

var rootCmd = &cobra.Command{
	Use:   "webscan",
	Short: "WebScan is a passive web audit CLI",
	Long: `WebScan is a modular CLI tool for passive web auditing.
It is designed for authorized security assessments and internal usage.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.Version = version
}
