package spcgi

type Appmgmt struct {
	Application string `json:"application,omitempty"`
	State       string `json:"state,omitempty"`
	Flags       []any  `json:"flags,omitempty"`
}
