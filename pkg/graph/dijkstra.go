package graph

import "container/heap"

type priorityItem struct {
	node string
	cost float64
	idx  int
}

type priorityQueue []*priorityItem

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].cost < pq[j].cost }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i]; pq[i].idx = i; pq[j].idx = j }
func (pq *priorityQueue) Push(x interface{}) { item := x.(*priorityItem); item.idx = len(*pq); *pq = append(*pq, item) }
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

func (g *Graph) ShortestPath(from, to string) ([]string, float64, bool) {
	if !g.HasNode(from) || !g.HasNode(to) {
		return nil, 0, false
	}
	if from == to {
		return []string{from}, 0, true
	}

	dist := make(map[string]float64)
	prev := make(map[string]string)

	for id := range g.Nodes {
		dist[id] = Infinity
	}
	dist[from] = 0

	pq := &priorityQueue{}
	heap.Push(pq, &priorityItem{node: from, cost: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*priorityItem)
		if current.cost > dist[current.node] {
			continue
		}
		if current.node == to {
			break
		}
		for _, edge := range g.Adjacency[current.node] {
			newCost := dist[current.node] + edge.Weight
			if newCost < dist[edge.To] {
				dist[edge.To] = newCost
				prev[edge.To] = current.node
				heap.Push(pq, &priorityItem{node: edge.To, cost: newCost})
			}
		}
	}

	if dist[to] == Infinity {
		return nil, 0, false
	}

	var path []string
	node := to
	for node != "" {
		path = append([]string{node}, path...)
		node = prev[node]
	}

	return path, dist[to], true
}
