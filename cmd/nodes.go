package cmd

import (
	"fmt"
	"strings"

	"github.com/mdryaan/netmapper/pkg/graph"
	"github.com/mdryaan/netmapper/pkg/output"
	"github.com/mdryaan/netmapper/pkg/parser"
	"github.com/spf13/cobra"
)

var nodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "List all nodes with type, status and connection details",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cfgFile == "" {
			output.Fatal("--config flag is required")
		}
		topology, err := parser.Load(cfgFile)
		if err != nil {
			output.Fatal(err.Error())
		}
		g := graph.Build(topology)

		table := output.NewTable([]string{"Name", "Type", "Status", "Degree", "Connected To"})
		for _, n := range topology.Nodes {
			node := g.Nodes[n.Name]
			degree := 0
			if node != nil {
				degree = node.Degree
			}

			var statusStr string
			switch n.Status {
			case "active":
				statusStr = output.Success(string(n.Status))
			case "inactive":
				statusStr = output.Error(string(n.Status))
			case "degraded":
				statusStr = output.Warn(string(n.Status))
			default:
				statusStr = string(n.Status)
			}

			var nbParts []string
			for _, edge := range g.Adjacency[n.Name] {
				nbParts = append(nbParts, fmt.Sprintf("%s(%.0fms)", edge.To, edge.Weight))
			}
			conns := strings.Join(nbParts, ", ")
			if conns == "" {
				conns = output.Warn("none (isolated)")
			}

			table.Append([]string{
				output.Info(n.Name),
				string(n.Type),
				statusStr,
				fmt.Sprintf("%d", degree),
				conns,
			})
		}
		table.Render()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(nodesCmd)
}
