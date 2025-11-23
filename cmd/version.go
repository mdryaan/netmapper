package cmd

import (
	"fmt"

	"github.com/mdryaan/netmapper/internal/config"
	"github.com/mdryaan/netmapper/pkg/output"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s  %s\n", output.Bold(config.AppName), output.Info("v"+config.Version))
		fmt.Println("Network Topology Visualizer & Analyzer")
		fmt.Println("Built with cobra · viper · go-yaml · tablewriter")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
