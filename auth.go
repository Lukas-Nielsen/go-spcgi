package spcgi

func (c *Client) Login() error {
	var resp Response
	if _, err := c.resty.R().
		SetHeader("User-Agent", useragent).
		SetBody(Request{
			Module: "auth",
			Command: []string{
				"login",
			},
			Arguments: map[string]any{
				"user": c.conf.User,
				"pass": c.conf.Password,
			},
		}).
		SetResult(&resp).
		Post(""); err != nil {
		return err
	}

	c.sessionID = resp.SessionID

	return nil
}

func (c *Client) Logout() error {
	if _, err := c.Request(Request{
		Module: "auth",
		Command: []string{
			"logout",
		},
	}); err != nil {
		return err
	}

	return nil
}
