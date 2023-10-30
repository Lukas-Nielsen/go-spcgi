package spcgi

import (
	"encoding/json"
	"fmt"
)

const (
	useragent = "github.com/Lukas-Nielsen/go-spcgi"
)

func (c *Client) Request(r Request) ([]byte, error) {
	if len(c.sessionID) != 0 {
		r.SessionID = c.sessionID
	}

	var resp Response
	if _, err := c.resty.R().SetHeader("User-Agent", useragent).SetBody(r).SetResult(&resp).Post(""); err != nil {
		return nil, err
	}

	if resp.Result.Code != 200 {
		return nil, fmt.Errorf("%s", resp.Result.Message)
	}

	j, err := json.Marshal(resp.Result.Content)
	if err != nil {
		return nil, err
	}

	return j, nil
}
