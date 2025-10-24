package visualizer

import (
	"fmt"
	"strings"
)

func (v *Visualizer) RenderTree(rootID string) string {
	var sb strings.Builder
	visited := make(map[string]bool)

	if rootID == "" && len(v.graph.NodeOrder) > 0 {
		rootID = v.graph.NodeOrder[0]
	}

	sb.WriteString(fmt.Sprintf("Topology Tree (root: %s)\n", rootID))
	sb.WriteString(strings.Repeat("─", 60) + "\n\n")

	v.renderTreeNode(&sb, rootID, "", true, visited)

	return sb.String()
}

func (v *Visualizer) renderTreeNode(sb *strings.Builder, id, prefix string, isRoot bool, visited map[string]bool) {
	if visited[id] {
		return
	}
	visited[id] = true

	node := v.graph.Nodes[id]
	nodeType := ""
	if node != nil {
		nodeType = node.Type
	}

	if isRoot {
		sb.WriteString(fmt.Sprintf("%s [%s]\n", id, nodeType))
	} else {
		sb.WriteString(fmt.Sprintf("%s\n", id))
	}

	edges := v.graph.Adjacency[id]
	for i, edge := range edges {
		if visited[edge.To] {
			continue
		}
		isLast := i == len(edges)-1
		connector := "├── "
		childPrefix := prefix + "│   "
		if isLast {
			connector = "└── "
			childPrefix = prefix + "    "
		}
		nb := v.graph.Nodes[edge.To]
		nbType := ""
		if nb != nil {
			nbType = nb.Type
		}
		sb.WriteString(fmt.Sprintf("%s%s(%.0fms) %s [%s] → ", prefix, connector, edge.Weight, edge.To, nbType))
		v.renderTreeNode(sb, edge.To, childPrefix, false, visited)
	}
}
