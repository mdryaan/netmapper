package cmd

import (
	"fmt"
	"strings"

	"github.com/mdryaan/netmapper/pkg/analyzer"
	"github.com/mdryaan/netmapper/pkg/graph"
	"github.com/mdryaan/netmapper/pkg/output"
	"github.com/mdryaan/netmapper/pkg/parser"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Run full topology analysis report",
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

		output.Header("NetMapper Analysis Report")
		fmt.Println(strings.Repeat("─", 60))
		fmt.Println()

		output.Header("Summary Statistics")
		fmt.Printf("  %-26s %s\n", "Total Nodes:", output.Info(fmt.Sprintf("%d", report.Stats.TotalNodes)))
		fmt.Printf("  %-26s %s\n", "Total Connections:", output.Info(fmt.Sprintf("%d", report.Stats.TotalConnections)))
		fmt.Printf("  %-26s %s\n", "Isolated Nodes:", warnOrOk(report.Stats.IsolatedCount))
		fmt.Printf("  %-26s %s\n", "Avg Degree:", output.Info(fmt.Sprintf("%.2f", report.Stats.AvgDegree)))
		fmt.Printf("  %-26s %s\n", "Max Degree:", output.Info(fmt.Sprintf("%d", report.Stats.MaxDegree)))
		fmt.Printf("  %-26s %s\n", "Min Degree:", output.Info(fmt.Sprintf("%d", report.Stats.MinDegree)))
		fmt.Println()

		output.Header("Node Type Distribution")
		for nodeType, count := range report.Stats.NodeTypeCounts {
			fmt.Printf("  %-14s %d\n", nodeType+":", count)
		}
		fmt.Println()

		if len(report.IsolatedNodes) > 0 {
			output.Header("Isolated Nodes")
			for _, iso := range report.IsolatedNodes {
				fmt.Printf("  %s  %s\n", output.Warn("⚠"), iso)
			}
			fmt.Println()
		}

		if len(report.ValidationErrors) == 0 {
			output.Successf("No validation errors found")
		} else {
			output.Header("Validation Errors")
			for _, ve := range report.ValidationErrors {
				fmt.Printf("  %s  [%s]  %s\n", output.Red("✗"), output.Yellow(ve.Code), ve.Message)
			}
		}
		fmt.Println()

		if len(report.UnreachablePairs) > 0 {
			output.Header("Unreachable Pairs")
			shown := 0
			for _, pair := range report.UnreachablePairs {
				if shown >= 10 {
					fmt.Printf("  %s\n", output.Dim(fmt.Sprintf("... and %d more pairs", len(report.UnreachablePairs)-shown)))
					break
				}
				fmt.Printf("  %s  →  %s\n", output.Warn(pair.From), output.Warn(pair.To))
				shown++
			}
			fmt.Println()
		}

		return nil
	},
}

func warnOrOk(n int) string {
	if n == 0 {
		return output.Success(fmt.Sprintf("%d", n))
	}
	return output.Warn(fmt.Sprintf("%d", n))
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
