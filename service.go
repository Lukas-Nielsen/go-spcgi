package spcgi

import (
	"encoding/json"
)

// Gibt eine Liste mit den Diensten zurück
func (c *Client) ServiceGet() ([]Service, error) {
	var data []Service

	d, err := c.Request(Request{
		Module: "service",
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

// Gibt eine Liste mit den Dienstgruppen zurück
func (c *Client) ServiceGroupGet() ([]ServiceGroup, error) {
	var data []ServiceGroup

	d, err := c.Request(Request{
		Module: "service",
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
