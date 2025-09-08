package parser

import (
	"fmt"

	"github.com/mdryaan/netmapper/internal/models"
	"github.com/mdryaan/netmapper/internal/utils"
)

type Parser interface {
	Parse(data []byte) (*models.Topology, error)
}

func Load(path string) (*models.Topology, error) {
	if !utils.FileExists(path) {
		return nil, fmt.Errorf("config file not found: %s", path)
	}

	data, err := utils.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	ext := utils.Extension(path)
	var p Parser

	switch ext {
	case "yaml", "yml":
		p = &YAMLParser{}
	case "json":
		p = &JSONParser{}
	default:
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}

	return p.Parse(data)
}
