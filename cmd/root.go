package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "netmapper",
	Short: "Network topology visualizer and analyzer",
	Long: `NetMapper is a production-grade CLI tool for visualizing, analyzing,
and exporting network topology from YAML or JSON configuration files.

Use --config to specify a topology file, then run any subcommand to inspect,
validate, visualize, or export the network.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "path to topology config file (YAML or JSON)")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "enable verbose output")
}
