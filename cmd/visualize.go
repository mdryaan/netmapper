package cmd

import (
	"github.com/mdryaan/netmapper/pkg/graph"
	"github.com/mdryaan/netmapper/pkg/output"
	"github.com/mdryaan/netmapper/pkg/parser"
	"github.com/mdryaan/netmapper/pkg/visualizer"
	"github.com/spf13/cobra"
)

var visualizeCmd = &cobra.Command{
	Use:   "visualize",
	Short: "Render ASCII graph of the network topology",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cfgFile == "" {
			output.Fatal("--config flag is required")
		}
		topology, err := parser.Load(cfgFile)
		if err != nil {
			output.Fatal(err.Error())
		}
		g := graph.Build(topology)
		v := visualizer.New(g)
		output.Println(v.RenderASCII())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(visualizeCmd)
}
