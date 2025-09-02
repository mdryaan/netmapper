package models

type Connection struct {
	From    string  `yaml:"from" json:"from"`
	To      string  `yaml:"to" json:"to"`
	Latency float64 `yaml:"latency" json:"latency"`
}
