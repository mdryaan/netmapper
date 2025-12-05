package cmd

import (
	"fmt"
	"strings"

	"github.com/mdryaan/netmapper/pkg/graph"
	"github.com/mdryaan/netmapper/pkg/output"
	"github.com/mdryaan/netmapper/pkg/parser"
	"github.com/spf13/cobra"
)

var (
	fromNode string
	toNode   string
)

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Find shortest path between two nodes using Dijkstra",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cfgFile == "" {
			output.Fatal("--config flag is required")
		}
		if fromNode == "" || toNode == "" {
			output.Fatal("--from and --to flags are required")
		}
		topology, err := parser.Load(cfgFile)
		if err != nil {
			output.Fatal(err.Error())
		}
		g := graph.Build(topology)

		path, cost, found := g.ShortestPath(fromNode, toNode)
		if !found {
			output.Warnf("No path found between '%s' and '%s'", fromNode, toNode)
			return nil
		}

		coloredNodes := make([]string, len(path))
		for i, n := range path {
			coloredNodes[i] = output.Info(n)
		}

		output.Successf("Path found: %s → %s", fromNode, toNode)
		fmt.Printf("  Route: %s\n", strings.Join(coloredNodes, output.Dim(" → ")))
		fmt.Printf("  Cost:  %s\n", output.Info(fmt.Sprintf("%.0fms total latency", cost)))
		fmt.Printf("  Hops:  %s\n", output.Info(fmt.Sprintf("%d", len(path)-1)))
		return nil
	},
}

func init() {
	pathCmd.Flags().StringVar(&fromNode, "from", "", "source node name (required)")
	pathCmd.Flags().StringVar(&toNode, "to", "", "destination node name (required)")
	rootCmd.AddCommand(pathCmd)
}
