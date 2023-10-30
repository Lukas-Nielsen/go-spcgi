package spcgi

import "encoding/json"

func (c *Client) CertGet() ([]Cert, error) {
	var data []Cert

	d, err := c.Request(Request{
		Module: "cert",
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

func (c *Client) CertAcmeaccountGet() ([]CertAcmeaccount, error) {
	var data []CertAcmeaccount

	d, err := c.Request(Request{
		Module: "cert",
		Command: []string{
			"acme_account",
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

func (c *Client) CertAcmenamesGet() ([]CertAcmenames, error) {
	var data []CertAcmenames

	d, err := c.Request(Request{
		Module: "cert",
		Command: []string{
			"acme_names",
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
