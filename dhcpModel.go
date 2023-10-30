package spcgi

type DhcpPool struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type DhcpLeaseStatic struct {
	Host     string `json:"host,omitempty"`
	Ethernet string `json:"ethernet,omitempty"`
	IP       string `json:"ip,omitempty"`
}

type DhcpLease struct {
	IP       string `json:"ip,omitempty"`
	Ethernet string `json:"ethernet,omitempty"`
	Host     string `json:"host,omitempty"`
	Cltt     string `json:"cltt,omitempty"`
	End      string `json:"end,omitempty"`
	Binding  string `json:"binding,omitempty"`
}
