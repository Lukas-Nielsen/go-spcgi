package spcgi

type Rule struct {
	Applayer    string   `json:"applayer,omitempty"`
	Comment     string   `json:"comment,omitempty"`
	Dst         string   `json:"dst,omitempty"`
	Flags       []string `json:"flags,omitempty"`
	Group       string   `json:"group,omitempty"`
	ID          int      `json:"id,omitempty"`
	Log         any      `json:"log,omitempty"`
	NatNode     string   `json:"nat_node,omitempty"`
	NatService  string   `json:"nat_service,omitempty"`
	Pos         int      `json:"pos,omitempty"`
	Qos         any      `json:"qos,omitempty"`
	Route       any      `json:"route,omitempty"`
	Service     string   `json:"service,omitempty"`
	Src         string   `json:"src,omitempty"`
	Timeprofile any      `json:"timeprofile,omitempty"`
}
