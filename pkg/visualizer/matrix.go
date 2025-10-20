package visualizer

import (
	"fmt"
	"strings"
)

func (v *Visualizer) RenderMatrix() string {
	var sb strings.Builder
	nodes := v.graph.NodeOrder
	n := len(nodes)

	maxLen := 10
	for _, id := range nodes {
		if len(id) > maxLen {
			maxLen = len(id)
		}
	}

	sb.WriteString("Connection Matrix\n")
	sb.WriteString(strings.Repeat("─", 60) + "\n\n")

	header := strings.Repeat(" ", maxLen+2)
	for _, id := range nodes {
		header += fmt.Sprintf("%-14s", shortName(id, 12))
	}
	sb.WriteString(header + "\n")

	matrix := make(map[string]map[string]float64)
	for _, id := range nodes {
		matrix[id] = make(map[string]float64)
	}
	for _, id := range nodes {
		for _, edge := range v.graph.Adjacency[id] {
			matrix[id][edge.To] = edge.Weight
		}
	}

	for i := 0; i < n; i++ {
		sb.WriteString(fmt.Sprintf("%-*s  ", maxLen, shortName(nodes[i], maxLen)))
		for j := 0; j < n; j++ {
			if i == j {
				sb.WriteString(fmt.Sprintf("%-14s", "  ─"))
			} else if w, ok := matrix[nodes[i]][nodes[j]]; ok {
				sb.WriteString(fmt.Sprintf("%-14s", fmt.Sprintf("  %.0fms", w)))
			} else {
				sb.WriteString(fmt.Sprintf("%-14s", "  ✗"))
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func shortName(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-1] + "…"
}
