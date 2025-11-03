package exporter

import (
	"encoding/json"

	"github.com/mdryaan/netmapper/internal/models"
)

type JSONExporter struct{}

type jsonReport struct {
	Nodes            []models.Node            `json:"nodes"`
	Connections      []models.Connection      `json:"connections"`
	Stats            *models.TopologyStats    `json:"stats"`
	ValidationErrors []models.ValidationError `json:"validation_errors"`
	IsolatedNodes    []string                 `json:"isolated_nodes"`
	UnreachablePairs []models.UnreachablePair `json:"unreachable_pairs"`
}

func (e *JSONExporter) Export(report *models.AnalysisReport) ([]byte, error) {
	r := jsonReport{
		Nodes:            report.Topology.Nodes,
		Connections:      report.Topology.Connections,
		Stats:            report.Stats,
		ValidationErrors: report.ValidationErrors,
		IsolatedNodes:    report.IsolatedNodes,
		UnreachablePairs: report.UnreachablePairs,
	}
	return json.MarshalIndent(r, "", "  ")
}
