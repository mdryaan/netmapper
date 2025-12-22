package parser

import (
	"fmt"

	"github.com/mdryaan/netmapper/internal/models"
	"gopkg.in/yaml.v3"
)

type YAMLParser struct{}

func (p *YAMLParser) Parse(data []byte) (*models.Topology, error) {
	var topology models.Topology
	if err := yaml.Unmarshal(data, &topology); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}
	for i := range topology.Nodes {
		if topology.Nodes[i].Status == "" {
			topology.Nodes[i].Status = models.NodeStatusActive
		}
	}
	return &topology, nil
}
