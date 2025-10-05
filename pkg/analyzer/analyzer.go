package analyzer

import (
	"github.com/mdryaan/netmapper/internal/models"
	"github.com/mdryaan/netmapper/pkg/graph"
)

type Analyzer struct {
	topology *models.Topology
	graph    *graph.Graph
}

func New(topology *models.Topology, g *graph.Graph) *Analyzer {
	return &Analyzer{
		topology: topology,
		graph:    g,
	}
}

func (a *Analyzer) Run() *models.AnalysisReport {
	return &models.AnalysisReport{
		Topology:         a.topology,
		ValidationErrors: Validate(a.topology, a.graph),
		IsolatedNodes:    FindIsolated(a.graph),
		UnreachablePairs: FindUnreachable(a.graph),
		Stats:            ComputeStats(a.topology, a.graph),
	}
}
