package visualizer

import "github.com/mdryaan/netmapper/pkg/graph"

type Visualizer struct {
	graph *graph.Graph
}

func New(g *graph.Graph) *Visualizer {
	return &Visualizer{graph: g}
}
