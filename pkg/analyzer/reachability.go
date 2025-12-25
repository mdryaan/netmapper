package analyzer

import (
	"github.com/mdryaan/netmapper/internal/models"
	"github.com/mdryaan/netmapper/pkg/graph"
)

func FindUnreachable(g *graph.Graph) []models.UnreachablePair {
	var pairs []models.UnreachablePair
	nodes := g.NodeOrder

	for i := range nodes {
		reachable := bfsReachable(g, nodes[i])
		for j := range nodes {
			if i != j && !reachable[nodes[j]] {
				pairs = append(pairs, models.UnreachablePair{From: nodes[i], To: nodes[j]})
			}
		}
	}

	return pairs
}

func bfsReachable(g *graph.Graph, start string) map[string]bool {
	visited := make(map[string]bool)
	queue := []string{start}
	visited[start] = true

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

	return visited
}

func IsReachable(g *graph.Graph, from, to string) bool {
	return bfsReachable(g, from)[to]
}
