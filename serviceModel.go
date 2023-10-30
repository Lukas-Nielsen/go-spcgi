package spcgi

type Service struct {
	CtHelper  any      `json:"ct_helper,omitempty"`
	ID        int      `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Proto     string   `json:"proto,omitempty"`
	DstPorts  []string `json:"dst-ports,omitempty"`
	IcmpTypes []string `json:"icmp-types,omitempty"`
}

type ServiceGroup struct {
	Descr    string   `json:"descr,omitempty"`
	ID       int      `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Services []string `json:"services,omitempty"`
}
