package analyzer

import (
	"fmt"

	"github.com/mdryaan/netmapper/internal/models"
	"github.com/mdryaan/netmapper/pkg/graph"
)

func Validate(topology *models.Topology, g *graph.Graph) []models.ValidationError {
	var errs []models.ValidationError
	nodeMap := topology.NodeMap()

	seen := make(map[string]bool)
	for _, n := range topology.Nodes {
		if n.Name == "" {
			errs = append(errs, models.ValidationError{Code: "EMPTY_NODE_NAME", Message: "node has empty name"})
		}
		if seen[n.Name] {
			errs = append(errs, models.ValidationError{Code: "DUPLICATE_NODE", Message: fmt.Sprintf("duplicate node name: '%s'", n.Name)})
		}
		seen[n.Name] = true
		if !n.IsValid() {
			errs = append(errs, models.ValidationError{Code: "INVALID_NODE_TYPE", Message: fmt.Sprintf("node '%s' has invalid type '%s'", n.Name, n.Type)})
		}
	}

	seenConns := make(map[string]bool)
	for _, c := range topology.Connections {
		if _, ok := nodeMap[c.From]; !ok {
			errs = append(errs, models.ValidationError{Code: "MISSING_NODE", Message: fmt.Sprintf("connection references unknown node: '%s'", c.From)})
		}
		if _, ok := nodeMap[c.To]; !ok {
			errs = append(errs, models.ValidationError{Code: "MISSING_NODE", Message: fmt.Sprintf("connection references unknown node: '%s'", c.To)})
		}
		if c.From == c.To {
			errs = append(errs, models.ValidationError{Code: "SELF_LOOP", Message: fmt.Sprintf("self-loop detected on node: '%s'", c.From)})
		}
		key := c.From + "↔" + c.To
		rkey := c.To + "↔" + c.From
		if seenConns[key] || seenConns[rkey] {
			errs = append(errs, models.ValidationError{Code: "DUPLICATE_CONNECTION", Message: fmt.Sprintf("duplicate connection: '%s' <-> '%s'", c.From, c.To)})
		}
		seenConns[key] = true
		if c.Latency < 0 {
			errs = append(errs, models.ValidationError{Code: "NEGATIVE_LATENCY", Message: fmt.Sprintf("negative latency on '%s' -> '%s': %.2f", c.From, c.To, c.Latency)})
		}
	}

	return errs
}
