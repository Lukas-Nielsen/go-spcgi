package spcgi

import "encoding/json"

func (c *Client) AlertingscenterAlertsGet() ([]AlertingcenterAlerts, error) {
	var data []AlertingcenterAlerts

	d, err := c.Request(Request{
		Module: "alertingcenter",
		Command: []string{
			"alerts",
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

func (c *Client) AlertingscenterAlertsProgramGet() ([]AlertingcenterAlertsProgram, error) {
	var data []AlertingcenterAlertsProgram

	d, err := c.Request(Request{
		Module: "alertingcenter",
		Command: []string{
			"alerts",
			"program",
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

func (c *Client) AlertingscenterSeverityGet() ([]AlertingcenterSeverity, error) {
	var data []AlertingcenterSeverity

	d, err := c.Request(Request{
		Module: "alertingcenter",
		Command: []string{
			"severity",
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

func (c *Client) AlertingscenterSyslogPatternGet() ([]AlertingcenterSyslogPattern, error) {
	var data []AlertingcenterSyslogPattern

	d, err := c.Request(Request{
		Module: "alertingcenter",
		Command: []string{
			"syslog",
			"pattern",
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

func (c *Client) AlertingscenterSyslogPatterngroupGet() ([]AlertingcenterSyslogPatterngroup, error) {
	var data []AlertingcenterSyslogPatterngroup

	d, err := c.Request(Request{
		Module: "alertingcenter",
		Command: []string{
			"syslog",
			"patterngroup",
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
