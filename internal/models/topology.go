package models

type Topology struct {
	Nodes       []Node       `yaml:"nodes" json:"nodes"`
	Connections []Connection `yaml:"connections" json:"connections"`
}

func (t *Topology) NodeMap() map[string]*Node {
	m := make(map[string]*Node, len(t.Nodes))
	for i := range t.Nodes {
		m[t.Nodes[i].Name] = &t.Nodes[i]
	}
	return m
}
