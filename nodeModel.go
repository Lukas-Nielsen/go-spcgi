package spcgi

type Node struct {
	Address string `json:"address,omitempty"`
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Zone    string `json:"zone,omitempty"`
}

type NodeGroup struct {
	ID    int      `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Nodes []string `json:"nodes,omitempty"`
}
