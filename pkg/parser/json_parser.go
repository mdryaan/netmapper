package parser

import (
	"encoding/json"

	"github.com/mdryaan/netmapper/internal/models"
)

type JSONParser struct{}

func (p *JSONParser) Parse(data []byte) (*models.Topology, error) {
	var topology models.Topology
	if err := json.Unmarshal(data, &topology); err != nil {
		return nil, err
	}
	return &topology, nil
}
