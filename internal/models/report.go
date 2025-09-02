package models

type ValidationError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PathResult struct {
	From  string
	To    string
	Path  []string
	Cost  float64
	Hops  int
	Found bool
}

type UnreachablePair struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type TopologyStats struct {
	TotalNodes       int            `json:"total_nodes"`
	TotalConnections int            `json:"total_connections"`
	IsolatedCount    int            `json:"isolated_count"`
	AvgDegree        float64        `json:"avg_degree"`
	MaxDegree        int            `json:"max_degree"`
	MinDegree        int            `json:"min_degree"`
	NodeTypeCounts   map[string]int `json:"node_type_counts"`
}

type AnalysisReport struct {
	Topology         *Topology
	Stats            *TopologyStats
	ValidationErrors []ValidationError
	IsolatedNodes    []string
	UnreachablePairs []UnreachablePair
}
