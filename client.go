package shelly

import (
	"github.com/jodydadescott/shelly-go-sdk/plus"
)

type config interface {
	GetHostname() string
	GetPassword() string
	IsDebugEnabled() bool
}

type Client struct {
	config config
	_plus  *PlusClient
}

func New(config config) *Client {
	return &Client{
		config: config,
	}
}

func (t *Client) PlusClient() (*plus.Client, error) {
	if t._plus != nil {
		return t._plus, nil
	}
	plus, err := plus.New(t.config)
	if err != nil {
		return nil, err
	}
	t._plus = plus
	return t._plus, nil
}

func (t *Client) Close() {
	if t._plus != nil {
		t._plus.Close()
	}
}
