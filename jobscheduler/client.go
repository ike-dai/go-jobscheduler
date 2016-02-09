package jobscheduler

import (
	"net/http"
)

type Client struct {
	Url string
	*http.Client
}

func NewClient(url string) *Client {
	return &Client{url, http.DefaultClient}
}
