package cmd

import (
	"fmt"
	"os"

	"github.com/mdryaan/netmapper/pkg/analyzer"
	"github.com/mdryaan/netmapper/pkg/exporter"
	"github.com/mdryaan/netmapper/pkg/graph"
	"github.com/mdryaan/netmapper/pkg/output"
	"github.com/mdryaan/netmapper/pkg/parser"
	"github.com/spf13/cobra"
)

var (
	exportFormat string
	outputFile   string
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export topology report as JSON, CSV or plain text",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cfgFile == "" {
			output.Fatal("--config flag is required")
		}
		topology, err := parser.Load(cfgFile)
		if err != nil {
			output.Fatal(err.Error())
		}
		g := graph.Build(topology)
		a := analyzer.New(topology, g)
		report := a.Run()

		exp, err := exporter.New(exportFormat)
		if err != nil {
			output.Fatal(err.Error())
		}

		data, err := exp.Export(report)
		if err != nil {
			output.Fatal(err.Error())
		}

		if outputFile != "" {
			if err := os.WriteFile(outputFile, data, 0644); err != nil {
				output.Fatal(err.Error())
			}
			output.Successf("Report exported to %s (%d bytes)", outputFile, len(data))
			return nil
		}

		fmt.Print(string(data))
		return nil
	},
}

func init() {
	exportCmd.Flags().StringVar(&exportFormat, "format", "json", "export format: json, csv, text")
	exportCmd.Flags().StringVar(&outputFile, "output", "", "write to file instead of stdout")
	rootCmd.AddCommand(exportCmd)
}
