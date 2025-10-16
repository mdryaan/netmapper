package analyzer

import (
	"github.com/mdryaan/netmapper/internal/models"
	"github.com/mdryaan/netmapper/internal/utils"
	"github.com/mdryaan/netmapper/pkg/graph"
)

func ComputeStats(topology *models.Topology, g *graph.Graph) *models.TopologyStats {
	stats := &models.TopologyStats{
		TotalNodes:       g.NodeCount(),
		TotalConnections: g.EdgeCount(),
		NodeTypeCounts:   make(map[string]int),
		MinDegree:        int(^uint(0) >> 1),
	}

	stats.IsolatedCount = len(FindIsolated(g))

	for _, n := range topology.Nodes {
		stats.NodeTypeCounts[string(n.Type)]++
	}

	totalDegree := 0
	for _, node := range g.Nodes {
		deg := node.Degree
		totalDegree += deg
		if deg > stats.MaxDegree {
			stats.MaxDegree = deg
		}
		if deg < stats.MinDegree {
			stats.MinDegree = deg
		}
	}

	if len(g.Nodes) == 0 {
		stats.MinDegree = 0
		return stats
	}

	if stats.MinDegree == int(^uint(0)>>1) {
		stats.MinDegree = 0
	}
	stats.AvgDegree = utils.RoundFloat(float64(totalDegree)/float64(len(g.Nodes)), 2)

	return stats
}
