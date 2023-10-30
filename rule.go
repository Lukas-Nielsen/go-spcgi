package spcgi

import (
	"encoding/json"
)

// Gibt eine Liste mit den Regeln zurück
func (c *Client) RuleGet() ([]Rule, error) {
	var data []Rule

	d, err := c.Request(Request{
		Module: "rule",
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
