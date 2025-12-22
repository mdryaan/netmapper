package parser

import (
	"encoding/json"
	"fmt"

	"github.com/mdryaan/netmapper/internal/models"
)

type JSONParser struct{}

func (p *JSONParser) Parse(data []byte) (*models.Topology, error) {
	var topology models.Topology
	if err := json.Unmarshal(data, &topology); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	for i := range topology.Nodes {
		if topology.Nodes[i].Status == "" {
			topology.Nodes[i].Status = models.NodeStatusActive
		}
	}
	return &topology, nil
}
