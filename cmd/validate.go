package cmd

import (
	"fmt"

	"github.com/mdryaan/netmapper/pkg/analyzer"
	"github.com/mdryaan/netmapper/pkg/graph"
	"github.com/mdryaan/netmapper/pkg/output"
	"github.com/mdryaan/netmapper/pkg/parser"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate topology config for schema errors",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cfgFile == "" {
			output.Fatal("--config flag is required")
		}
		topology, err := parser.Load(cfgFile)
		if err != nil {
			output.Fatal(err.Error())
		}
		g := graph.Build(topology)
		errs := analyzer.Validate(topology, g)

		if len(errs) == 0 {
			output.Successf("Config is valid — %d nodes, %d connections", len(topology.Nodes), len(topology.Connections))
			return nil
		}

		output.Warnf("Found %d validation error(s):", len(errs))
		for _, ve := range errs {
			fmt.Printf("  %s  [%s]  %s\n", output.Red("✗"), output.Yellow(ve.Code), ve.Message)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
