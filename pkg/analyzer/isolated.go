package analyzer

import "github.com/mdryaan/netmapper/pkg/graph"

func FindIsolated(g *graph.Graph) []string {
	var isolated []string
	for _, id := range g.NodeOrder {
		if g.Nodes[id].Degree == 0 {
			isolated = append(isolated, id)
		}
	}
	return isolated
}
