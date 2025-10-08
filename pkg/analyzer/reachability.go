package analyzer

import (
	"github.com/mdryaan/netmapper/internal/models"
	"github.com/mdryaan/netmapper/pkg/graph"
)

func FindUnreachable(g *graph.Graph) []models.UnreachablePair {
	var pairs []models.UnreachablePair
	nodes := g.NodeOrder

	for i := range nodes {
		visited := make(map[string]bool)
		queue := []string{nodes[i]}
		visited[nodes[i]] = true

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, edge := range g.Adjacency[cur] {
				if !visited[edge.To] {
					visited[edge.To] = true
					queue = append(queue, edge.To)
				}
			}
		}

		for j := range nodes {
			if i != j && !visited[nodes[j]] {
				pairs = append(pairs, models.UnreachablePair{From: nodes[i], To: nodes[j]})
			}
		}
	}

	return pairs
}

func IsReachable(g *graph.Graph, from, to string) bool {
	visited := make(map[string]bool)
	queue := []string{from}
	visited[from] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur == to {
			return true
		}
		for _, edge := range g.Adjacency[cur] {
			if !visited[edge.To] {
				visited[edge.To] = true
				queue = append(queue, edge.To)
			}
		}
	}

	return false
}
