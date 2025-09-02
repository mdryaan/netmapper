package models

type NodeType string
type NodeStatus string

const (
	NodeTypeRouter   NodeType = "router"
	NodeTypeSwitch   NodeType = "switch"
	NodeTypeServer   NodeType = "server"
	NodeTypeGateway  NodeType = "gateway"
	NodeTypeFirewall NodeType = "firewall"
)

const (
	NodeStatusActive   NodeStatus = "active"
	NodeStatusInactive NodeStatus = "inactive"
	NodeStatusDegraded NodeStatus = "degraded"
)

type Node struct {
	Name   string     `yaml:"name" json:"name"`
	Type   NodeType   `yaml:"type" json:"type"`
	Status NodeStatus `yaml:"status" json:"status"`
}

func (n *Node) IsValid() bool {
	switch n.Type {
	case NodeTypeRouter, NodeTypeSwitch, NodeTypeServer, NodeTypeGateway, NodeTypeFirewall:
		return true
	}
	return false
}
