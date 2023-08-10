package plus

import (
	"go.uber.org/zap"

	"github.com/jodydadescott/shelly-go-sdk/plus/bluetooth"
	"github.com/jodydadescott/shelly-go-sdk/plus/cloud"
	"github.com/jodydadescott/shelly-go-sdk/plus/ethernet"
	"github.com/jodydadescott/shelly-go-sdk/plus/input"
	"github.com/jodydadescott/shelly-go-sdk/plus/light"
	"github.com/jodydadescott/shelly-go-sdk/plus/mqtt"
	"github.com/jodydadescott/shelly-go-sdk/plus/msghandlers"
	"github.com/jodydadescott/shelly-go-sdk/plus/shelly"
	"github.com/jodydadescott/shelly-go-sdk/plus/switchx"
	"github.com/jodydadescott/shelly-go-sdk/plus/system"
	"github.com/jodydadescott/shelly-go-sdk/plus/types"
	"github.com/jodydadescott/shelly-go-sdk/plus/websocket"
	"github.com/jodydadescott/shelly-go-sdk/plus/wifi"
)

type Config interface {
	GetHostname() string
	GetPassword() string
	IsDebugEnabled() bool
}

type Client struct {
	_system    *system.Client
	_shelly    *shelly.Client
	_wifi      *wifi.Client
	_bluetooth *bluetooth.Client
	_mqtt      *mqtt.Client
	_cloud     *cloud.Client
	_switch    *switchx.Client
	_light     *light.Client
	_input     *input.Client
	_websocket *websocket.Client
	_ethernet  *ethernet.Client
	types.MessageHandlerFactory
}

func New(config Config) (*Client, error) {

	messageHandlerFactory, err := msghandlers.NewWS(&msghandlers.Config{
		Hostname:     config.GetHostname(),
		Password:     config.GetPassword(),
		Username:     types.ShellyUser,
		DebugEnabled: config.IsDebugEnabled(),
	})

	if err != nil {
		return nil, err
	}

	return &Client{
		MessageHandlerFactory: messageHandlerFactory,
	}, nil
}

func (t *Client) System() *system.Client {
	if t._system == nil {
		t._system = system.New(t)
	}
	return t._system
}

func (t *Client) Shelly() *shelly.Client {
	if t._shelly == nil {
		t._shelly = shelly.New(t)
	}
	return t._shelly
}

func (t *Client) Bluetooth() *bluetooth.Client {
	if t._bluetooth == nil {
		t._bluetooth = bluetooth.New(t)
	}
	return t._bluetooth
}

func (t *Client) Mqtt() *mqtt.Client {
	if t._mqtt == nil {
		t._mqtt = mqtt.New(t)
	}
	return t._mqtt
}

func (t *Client) Ethernet() *ethernet.Client {
	if t._ethernet == nil {
		t._ethernet = ethernet.New(t)
	}
	return t._ethernet
}

func (t *Client) Wifi() *wifi.Client {
	if t._wifi == nil {
		t._wifi = wifi.New(t)
	}
	return t._wifi
}

func (t *Client) Cloud() *cloud.Client {
	if t._cloud == nil {
		t._cloud = cloud.New(t)
	}
	return t._cloud
}

func (t *Client) Switch() *switchx.Client {
	if t._switch == nil {
		t._switch = switchx.New(t)
	}
	return t._switch
}

func (t *Client) Light() *light.Client {
	if t._light == nil {
		t._light = light.New(t)
	}
	return t._light
}

func (t *Client) Input() *input.Client {
	if t._input == nil {
		t._input = input.New(t)
	}
	return t._input
}

func (t *Client) Websocket() *websocket.Client {
	if t._websocket == nil {
		t._websocket = websocket.New(t)
	}
	return t._websocket
}

func (t *Client) Close() {

	zap.L().Debug("(*Client) Close()")

	if t._system != nil {
		t._system.Close()
	}

	if t._shelly != nil {
		t._shelly.Close()
	}

	if t._wifi != nil {
		t._wifi.Close()
	}

	if t._bluetooth != nil {
		t._bluetooth.Close()
	}

	if t._mqtt != nil {
		t._mqtt.Close()
	}

	if t._cloud != nil {
		t._cloud.Close()
	}

	if t._switch != nil {
		t._switch.Close()
	}

	if t._light != nil {
		t._light.Close()
	}

	if t._input != nil {
		t._input.Close()
	}

	if t._websocket != nil {
		t._websocket.Close()
	}

	if t._ethernet != nil {
		t._ethernet.Close()
	}

	t.MessageHandlerFactory.Close()
}
