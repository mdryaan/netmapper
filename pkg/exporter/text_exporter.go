package exporter

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/mdryaan/netmapper/internal/models"
)

type TextExporter struct{}

func (e *TextExporter) Export(report *models.AnalysisReport) ([]byte, error) {
	var buf bytes.Buffer

	fmt.Fprintln(&buf, "NetMapper Analysis Report")
	fmt.Fprintln(&buf, strings.Repeat("=", 60))
	fmt.Fprintln(&buf)

	fmt.Fprintln(&buf, "SUMMARY STATISTICS")
	fmt.Fprintln(&buf, strings.Repeat("-", 30))
	fmt.Fprintf(&buf, "  Total Nodes:         %d\n", report.Stats.TotalNodes)
	fmt.Fprintf(&buf, "  Total Connections:   %d\n", report.Stats.TotalConnections)
	fmt.Fprintf(&buf, "  Isolated Nodes:      %d\n", report.Stats.IsolatedCount)
	fmt.Fprintf(&buf, "  Average Degree:      %.2f\n", report.Stats.AvgDegree)
	fmt.Fprintf(&buf, "  Max Degree:          %d\n", report.Stats.MaxDegree)
	fmt.Fprintln(&buf)

	fmt.Fprintln(&buf, "NODES")
	fmt.Fprintln(&buf, strings.Repeat("-", 30))
	for _, n := range report.Topology.Nodes {
		fmt.Fprintf(&buf, "  %-22s  %-12s  %s\n", n.Name, n.Type, n.Status)
	}
	fmt.Fprintln(&buf)

	fmt.Fprintln(&buf, "CONNECTIONS")
	fmt.Fprintln(&buf, strings.Repeat("-", 30))
	for _, c := range report.Topology.Connections {
		fmt.Fprintf(&buf, "  %-22s  →  %-22s  (%.0fms)\n", c.From, c.To, c.Latency)
	}
	fmt.Fprintln(&buf)

	if len(report.IsolatedNodes) > 0 {
		fmt.Fprintln(&buf, "ISOLATED NODES")
		fmt.Fprintln(&buf, strings.Repeat("-", 30))
		for _, iso := range report.IsolatedNodes {
			fmt.Fprintf(&buf, "  ! %s\n", iso)
		}
		fmt.Fprintln(&buf)
	}

	if len(report.ValidationErrors) > 0 {
		fmt.Fprintln(&buf, "VALIDATION ERRORS")
		fmt.Fprintln(&buf, strings.Repeat("-", 30))
		for _, ve := range report.ValidationErrors {
			fmt.Fprintf(&buf, "  [%s] %s\n", ve.Code, ve.Message)
		}
		fmt.Fprintln(&buf)
	}

	return buf.Bytes(), nil
}
