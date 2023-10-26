package shelly

import (
	"fmt"
	"time"

	"github.com/jodydadescott/shelly-go-sdk/plus"
	"go.uber.org/zap"
)

type config interface {
	GetHostname() string
	GetPassword() string
	IsDebugEnabled() bool
	GetSendTimeout() time.Duration
}

type Client struct {
	config config
	_plus  *PlusClient
}

func New(config config) *Client {
	zap.L().Debug(fmt.Sprintf("SDK version %s", Version))
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
