package spcgi

import "encoding/json"

// Erstellt einen neuen DHCP Pool
//
// pool-name, pool-start and pool-end
func (c *Client) DhcpPoolNew(n string, s string, e string) error {
	_, err := c.Request(Request{
		Module: "dhcp",
		Command: []string{
			"pool",
			"new",
		},
		Arguments: map[string]any{
			"name":  n,
			"start": s,
			"end":   e,
		},
	})

	return err
}

// Ändert die Eigenschaften eines bestehenden DHCP Pool
func (c *Client) DhcpPoolSet(d DhcpPool) error {
	arguments := map[string]any{
		"id": d.ID,
	}

	if len(d.Start) > 0 {
		arguments["start"] = d.Start
	}

	if len(d.End) > 0 {
		arguments["end"] = d.End
	}

	_, err := c.Request(Request{
		Module: "dhcp",
		Command: []string{
			"pool",
			"set",
		},
		Arguments: arguments,
	})

	return err

}

// Auflistung der angelegten DHCP Pools
func (c *Client) DhcpPoolGet() ([]DhcpPool, error) {
	var data []DhcpPool

	d, err := c.Request(Request{
		Module: "dhcp",
		Command: []string{
			"pool",
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

// Löschen eines bestehenden DHCP Pool
//
// name
func (c *Client) DhcpPoolDelete(n string) error {
	_, err := c.Request(Request{
		Module: "dhcp",
		Command: []string{
			"pool",
			"delete",
		},
		Arguments: map[string]any{
			"name": n,
		},
	})

	return err
}

// Erstellt eine neues statisches DHCP Lease
//
// host, mac, ip
func (c *Client) DhcpLeaseNew(h string, m string, i string) error {
	_, err := c.Request(Request{
		Module: "dhcp",
		Command: []string{
			"lease",
			"new",
		},
		Arguments: map[string]any{
			"host":     h,
			"ethernet": m,
			"ip":       i,
		},
	})

	return err
}

// Ändert ein bestehendes statisches DHCP Lease
//
// id, host, mac, ip
func (c *Client) DhcpLeaseSet(id int, h string, m string, i string) error {
	_, err := c.Request(Request{
		Module: "dhcp",
		Command: []string{
			"lease",
			"set",
		},
		Arguments: map[string]any{
			"id":       id,
			"host":     h,
			"ethernet": m,
			"ip":       i,
		},
	})

	return err
}

// Auflistung der vorhandenen statischen DHCP Leases
func (c *Client) DhcpLeaseGet() ([]DhcpLeaseStatic, error) {
	var data []DhcpLeaseStatic

	d, err := c.Request(Request{
		Module: "dhcp",
		Command: []string{
			"lease",
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

// Löschen eines bestehenden statischen DHCP Lease
//
// host
func (c *Client) DhcpLeaseDelete(h string) error {
	_, err := c.Request(Request{
		Module: "dhcp",
		Command: []string{
			"lease",
			"delete",
		},
		Arguments: map[string]any{
			"host": h,
		},
	})

	return err
}

// Auflistung der zugeteilten DHCP Leases
func (c *Client) DhcpLeaseList() ([]DhcpLease, error) {
	var data []DhcpLease

	d, err := c.Request(Request{
		Module: "dhcp",
		Command: []string{
			"lease",
			"list",
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
