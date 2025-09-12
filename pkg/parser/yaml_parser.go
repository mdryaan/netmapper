package parser

import (
	"github.com/mdryaan/netmapper/internal/models"
	"gopkg.in/yaml.v3"
)

type YAMLParser struct{}

func (p *YAMLParser) Parse(data []byte) (*models.Topology, error) {
	var topology models.Topology
	if err := yaml.Unmarshal(data, &topology); err != nil {
		return nil, err
	}
	return &topology, nil
}
