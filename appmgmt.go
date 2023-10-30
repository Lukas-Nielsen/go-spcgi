package spcgi

import "encoding/json"

func (c *Client) AppmgmtUpdate() error {
	_, err := c.Request(Request{
		Module: "appmgmt",
		Command: []string{
			"update",
		},
	})

	return err
}

// pass application name
func (c *Client) AppmgmtUpdateApplication(a string) error {
	_, err := c.Request(Request{
		Module: "appmgmt",
		Command: []string{
			"update",
		},
		Arguments: map[string]any{
			"application": a,
		},
	})

	return err
}

func (c *Client) AppmgmtGet() ([]Appmgmt, error) {
	var data []Appmgmt

	d, err := c.Request(Request{
		Module: "appmgmt",
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

// pass application name and flags
func (c *Client) AppmgmtSet(a string, f ...string) error {
	_, err := c.Request(Request{
		Module: "appmgmt",
		Command: []string{
			"set",
		},
		Arguments: map[string]any{
			"application": a,
			"flags":       f,
		},
	})

	return err
}

// pass application name
func (c *Client) AppmgmtStart(a string) error {
	_, err := c.Request(Request{
		Module: "appmgmt",
		Command: []string{
			"start",
		},
		Arguments: map[string]any{
			"application": a,
		},
	})

	return err
}

// pass application name
func (c *Client) AppmgmtStop(a string) error {
	_, err := c.Request(Request{
		Module: "appmgmt",
		Command: []string{
			"stop",
		},
		Arguments: map[string]any{
			"application": a,
		},
	})

	return err
}

// pass application name
func (c *Client) AppmgmtRestart(a string) error {
	_, err := c.Request(Request{
		Module: "appmgmt",
		Command: []string{
			"restart",
		},
		Arguments: map[string]any{
			"application": a,
		},
	})

	return err
}

func (c *Client) AppmgmtStatus() ([]Appmgmt, error) {
	var data []Appmgmt

	d, err := c.Request(Request{
		Module: "appmgmt",
		Command: []string{
			"status",
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

func (c *Client) AppmgmtConfig() error {
	_, err := c.Request(Request{
		Module: "appmgmt",
		Command: []string{
			"config",
		},
	})

	return err
}
