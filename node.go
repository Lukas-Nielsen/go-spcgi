package spcgi

import (
	"encoding/json"
)

// Gibt eine Liste mit den Netwerkobjekten zurück
func (c *Client) NodeGet() ([]Node, error) {
	var data []Node

	d, err := c.Request(Request{
		Module: "node",
		Command: []string{
			"get",
		},
	})

	if err != nil {
		return nil, err
	}

	if json.Unmarshal(d, &data) != nil {
		return nil, err
	}
	return data, nil
}

// Gibt eine Liste mit den Netzwerkobjektgruppen zurück
func (c *Client) NodeGroupGet() ([]NodeGroup, error) {
	var data []NodeGroup

	d, err := c.Request(Request{
		Module: "node",
		Command: []string{
			"group",
			"get",
		},
	})

	if err != nil {
		return nil, err
	}

	if json.Unmarshal(d, &data) != nil {
		return nil, err
	}
	return data, nil
}
