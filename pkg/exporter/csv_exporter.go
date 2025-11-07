package exporter

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"

	"github.com/mdryaan/netmapper/internal/models"
)

type CSVExporter struct{}

func (e *CSVExporter) Export(report *models.AnalysisReport) ([]byte, error) {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)

	w.Write([]string{"section", "key", "value"})
	w.Write([]string{"stats", "total_nodes", strconv.Itoa(report.Stats.TotalNodes)})
	w.Write([]string{"stats", "total_connections", strconv.Itoa(report.Stats.TotalConnections)})
	w.Write([]string{"stats", "isolated_nodes", strconv.Itoa(report.Stats.IsolatedCount)})
	w.Write([]string{"stats", "avg_degree", fmt.Sprintf("%.2f", report.Stats.AvgDegree)})

	for nodeType, count := range report.Stats.NodeTypeCounts {
		w.Write([]string{"stats", "type_" + nodeType, strconv.Itoa(count)})
	}

	for _, n := range report.Topology.Nodes {
		w.Write([]string{"node", n.Name, string(n.Type) + "/" + string(n.Status)})
	}

	for _, c := range report.Topology.Connections {
		w.Write([]string{"connection", c.From + "->" + c.To, fmt.Sprintf("%.2fms", c.Latency)})
	}

	for _, iso := range report.IsolatedNodes {
		w.Write([]string{"isolated", iso, "true"})
	}

	for _, ve := range report.ValidationErrors {
		w.Write([]string{"error", ve.Code, ve.Message})
	}

	w.Flush()
	return buf.Bytes(), w.Error()
}
