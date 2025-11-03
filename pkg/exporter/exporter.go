package exporter

import (
	"fmt"

	"github.com/mdryaan/netmapper/internal/models"
)

type Exporter interface {
	Export(report *models.AnalysisReport) ([]byte, error)
}

func New(format string) (Exporter, error) {
	switch format {
	case "json":
		return &JSONExporter{}, nil
	case "csv":
		return &CSVExporter{}, nil
	case "text", "txt":
		return &TextExporter{}, nil
	default:
		return nil, fmt.Errorf("unsupported export format: %s", format)
	}
}
