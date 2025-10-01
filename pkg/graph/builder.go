package graph

import "github.com/mdryaan/netmapper/internal/models"

func Build(topology *models.Topology) *Graph {
	g := New()

	for _, n := range topology.Nodes {
		g.AddNode(&GraphNode{
			ID:     n.Name,
			Name:   n.Name,
			Type:   string(n.Type),
			Status: string(n.Status),
		})
	}

	for _, c := range topology.Connections {
		if g.HasNode(c.From) && g.HasNode(c.To) {
			g.AddEdge(c.From, c.To, c.Latency)
		}
	}

	return g
}
