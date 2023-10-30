package spcgi

import (
	"crypto/tls"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Conf struct {
	User     string
	Password string
	Host     string
	Port     int
	Insecure bool
}

type Client struct {
	conf      *Conf
	resty     *resty.Client
	sessionID string
}

type Request struct {
	Module    string         `json:"module,omitempty"`
	Command   []string       `json:"command,omitempty"`
	SessionID string         `json:"sessionid,omitempty"`
	Arguments map[string]any `json:"arguments,omitempty"`
}

type Response struct {
	SessionID string `json:"sessionid,omitempty"`
	Mode      string `json:"mode,omitempty"`
	Result    struct {
		Module  string `json:"module,omitempty"`
		Code    int    `json:"code,omitempty"`
		Status  string `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
		Content []any  `json:"content,omitempty"`
	} `json:"result,omitempty"`
	Version string `json:"version,omitempty"`
}

func NewCLient(conf Conf) *Client {
	var c Conf

	if len(conf.User) == 0 {
		c.User = "admin"
	} else {
		c.User = conf.User
	}

	if len(conf.Password) == 0 {
		c.Password = "insecure"
	} else {
		c.Password = conf.Password
	}

	if len(conf.Host) == 0 {
		c.Host = "192.168.175.1"
	} else {
		c.Host = conf.Host
	}

	if conf.Port < 1 {
		c.Port = 11115
	} else {
		c.Port = conf.Port
	}

	c.Insecure = conf.Insecure

	client := Client{
		conf: &c,
	}

	client.setResty()

	return &client
}

func (c *Client) SetUser(i string) *Client {
	c.conf.User = i
	return c
}

func (c *Client) SetPassword(i string) *Client {
	c.conf.Password = i
	return c
}

func (c *Client) SetHost(i string) *Client {
	c.conf.Host = i
	return c
}

func (c *Client) SetPort(i int) *Client {
	c.conf.Port = i
	return c
}

func (c *Client) SetInsecure(i bool) *Client {
	c.conf.Insecure = i
	return c
}

func (c *Client) setResty() {
	c.resty = resty.New().
		SetBaseURL("https://" + c.conf.Host + ":" + fmt.Sprint(c.conf.Port) + "/spcgi.cgi").
		SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: c.conf.Insecure,
		})
}
