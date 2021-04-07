package gohelper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type client struct {
	url     string
	method  string
	query   string
	body    interface{}
	timeout time.Duration
}

func NewHttp(url string) *client {
	return &client{timeout: 5 * time.Second, url: url}
}

func (c *client) SetQuery(params map[string]string) {
	query := ""
	for k, v := range params {
		query = fmt.Sprintf("%s%s=%s&", query, k, url.QueryEscape(v))
	}
	c.query = query
}

func (c *client) SetTimeout(second int) {
	c.timeout = time.Second * time.Duration(second)
}

func (c *client) SetMethod(method string) *client {
	c.method = method
	return c
}

func (c *client) Run() (res []byte, err error) {
	client := &http.Client{}
	if c.query != "" {
		c.url = StrCombine(c.url, "?", c.query)
	}
	fmt.Println(c.url)
	req, err := http.NewRequest(c.method, c.url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
