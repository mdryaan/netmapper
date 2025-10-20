package visualizer

import (
	"fmt"
	"strings"
)

var nodeSymbols = map[string]string{
	"router":   "[R] ",
	"switch":   "[SW]",
	"server":   "[SV]",
	"gateway":  "[GW]",
	"firewall": "[FW]",
}

func (v *Visualizer) RenderASCII() string {
	var sb strings.Builder

	sb.WriteString("Network Topology Graph\n")
	sb.WriteString(strings.Repeat("─", 60) + "\n\n")

	for _, id := range v.graph.NodeOrder {
		node := v.graph.Nodes[id]
		sym := nodeSymbols[node.Type]
		if sym == "" {
			sym = "[??]"
		}

		edges := v.graph.Adjacency[id]
		if len(edges) == 0 {
			sb.WriteString(fmt.Sprintf("  %s  %s  (isolated)\n\n", sym, node.Name))
			continue
		}

		sb.WriteString(fmt.Sprintf("  %s  %s\n", sym, node.Name))
		for i, edge := range edges {
			connector := "  ├──"
			if i == len(edges)-1 {
				connector = "  └──"
			}
			nbSym := nodeSymbols[v.graph.Nodes[edge.To].Type]
			if nbSym == "" {
				nbSym = "[??]"
			}
			sb.WriteString(fmt.Sprintf("     %s (%.0fms) ── %s  %s\n", connector, edge.Weight, nbSym, edge.To))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
