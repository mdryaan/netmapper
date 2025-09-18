package graph

import "math"

const Infinity = math.MaxFloat64

type Graph struct {
	Nodes     map[string]*GraphNode
	Adjacency map[string][]Edge
	NodeOrder []string
}

func New() *Graph {
	return &Graph{
		Nodes:     make(map[string]*GraphNode),
		Adjacency: make(map[string][]Edge),
		NodeOrder: []string{},
	}
}

func (g *Graph) AddNode(node *GraphNode) {
	if _, exists := g.Nodes[node.ID]; !exists {
		g.Nodes[node.ID] = node
		g.NodeOrder = append(g.NodeOrder, node.ID)
		g.Adjacency[node.ID] = []Edge{}
	}
}

func (g *Graph) AddEdge(from, to string, weight float64) {
	if weight == 0 {
		weight = 1
	}
	g.Adjacency[from] = append(g.Adjacency[from], Edge{From: from, To: to, Weight: weight})
	g.Adjacency[to] = append(g.Adjacency[to], Edge{From: to, To: from, Weight: weight})
	if n, ok := g.Nodes[from]; ok {
		n.Degree++
	}
	if n, ok := g.Nodes[to]; ok {
		n.Degree++
	}
}

func (g *Graph) HasNode(id string) bool {
	_, ok := g.Nodes[id]
	return ok
}

func (g *Graph) NodeCount() int {
	return len(g.Nodes)
}

func (g *Graph) EdgeCount() int {
	total := 0
	for _, edges := range g.Adjacency {
		total += len(edges)
	}
	return total / 2
}
