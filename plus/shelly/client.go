package shelly

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/jodydadescott/shelly-go-sdk/plus/bluetooth"
	"github.com/jodydadescott/shelly-go-sdk/plus/cloud"
	"github.com/jodydadescott/shelly-go-sdk/plus/ethernet"
	"github.com/jodydadescott/shelly-go-sdk/plus/input"
	"github.com/jodydadescott/shelly-go-sdk/plus/light"
	"github.com/jodydadescott/shelly-go-sdk/plus/mqtt"
	"github.com/jodydadescott/shelly-go-sdk/plus/switchx"
	"github.com/jodydadescott/shelly-go-sdk/plus/system"
	"github.com/jodydadescott/shelly-go-sdk/plus/websocket"
	"github.com/jodydadescott/shelly-go-sdk/plus/wifi"
)

type clientContract interface {
	System() *system.Client
	Bluetooth() *bluetooth.Client
	Mqtt() *mqtt.Client
	Wifi() *wifi.Client
	Cloud() *cloud.Client
	Switch() *switchx.Client
	Input() *input.Client
	Light() *light.Client
	Websocket() *websocket.Client
	Ethernet() *ethernet.Client
	NewHandle() MessageHandler
}

func New(clientContract clientContract) *Client {
	return &Client{
		clientContract: clientContract,
	}
}

type Client struct {
	clientContract
	_messageHandler MessageHandler
}

func (t *Client) getMessageHandler() MessageHandler {
	if t._messageHandler != nil {
		return t._messageHandler
	}

	t._messageHandler = t.NewHandle()
	return t._messageHandler
}

// GetStatus returns the status of all the components of the device.
func (t *Client) GetStatus(ctx context.Context) (*ShellyStatus, error) {

	method := Component + ".GetStatus"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: &method,
	})
	if err != nil {
		return nil, err
	}

	response := &GetStatusResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	if response.Result == nil {
		return nil, fmt.Errorf("Result is missing from response")
	}

	return response.Result.convert(), nil
}

// ListMethods lists all available RPC methods. It takes into account both ACL and authentication restrictions
// and only lists the methods allowed for the particular user/channel that's making the request.
func (t *Client) ListMethods(ctx context.Context) (*ShellyRPCMethods, error) {

	// Do NOT validate command here because it would be recursive

	method := Component + ".ListMethods"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: &method,
	})
	if err != nil {
		return nil, err
	}

	response := &ListMethodsResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	if response.Result == nil {
		return nil, fmt.Errorf("Result is missing from response")
	}

	return response.Result, nil
}

// GetConfig returns the configuration of all the components of the device.
func (t *Client) GetConfig(ctx context.Context, markup bool) (*ShellyConfig, error) {

	method := Component + ".GetConfig"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: &method,
	})
	if err != nil {
		return nil, err
	}

	response := &GetConfigResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	if response.Result == nil {
		return nil, fmt.Errorf("Result is missing from response")
	}

	config := response.Result.convert()

	if markup {

		if t.getMessageHandler().IsAuthEnabled() {
			config.Auth.Enable = true
		}

		config.Markup()
	}

	return config, nil
}

// SetConfig sets the configuration for each component with non nil config. Note that this function
// calls into each componenet as necessary.
func (t *Client) SetConfig(ctx context.Context, config *ShellyConfig) *ShellyReport {

	config = config.Clone()
	config.Sanatize()

	report := &ShellyReport{}

	if config.Bluetooth != nil {
		report.Bluetooth = &ComponentReport{}
		rebootRequired, err := t.Bluetooth().SetConfig(ctx, config.Bluetooth)
		report.Bluetooth.RebootRequired = rebootRequired
		report.Bluetooth.Error = err
	}

	if config.Cloud != nil {
		report.Cloud = &ComponentReport{}
		rebootRequired, err := t.Cloud().SetConfig(ctx, config.Cloud)
		report.Cloud.RebootRequired = rebootRequired
		report.Cloud.Error = err
	}

	if config.Ethernet != nil {
		report.Ethernet = &ComponentReport{}
		rebootRequired, err := t.Ethernet().SetConfig(ctx, config.Ethernet)
		report.Ethernet.RebootRequired = rebootRequired
		report.Ethernet.Error = err
	}

	if config.Mqtt != nil {
		report.Mqtt = &ComponentReport{}
		rebootRequired, err := t.Mqtt().SetConfig(ctx, config.Mqtt)
		report.Mqtt.RebootRequired = rebootRequired
		report.Mqtt.Error = err
	}

	if config.System != nil {
		report.System = &ComponentReport{}
		rebootRequired, err := t.System().SetConfig(ctx, config.System)
		report.System.RebootRequired = rebootRequired
		report.System.Error = err
	}

	if config.Websocket != nil {
		report.Websocket = &ComponentReport{}

		rebootRequired, err := t.Websocket().SetConfig(ctx, config.Websocket)
		report.Websocket.RebootRequired = rebootRequired
		report.Websocket.Error = err
	}

	if config.Wifi != nil {
		report.Wifi = &ComponentReport{}

		rebootRequired, err := t.Wifi().SetConfig(ctx, config.Wifi)
		report.Wifi.RebootRequired = rebootRequired
		report.Wifi.Error = err
	}

	if config.Light != nil {
		for _, v := range config.Light {
			report.Light = append(report.Light, &ComponentReport{
				ID:    &v.ID,
				Error: t.Light().SetConfig(ctx, v),
			})
		}
	}

	if config.Input != nil {
		for _, v := range config.Input {
			report.Input = append(report.Input, &ComponentReport{
				ID:    &v.ID,
				Error: t.Input().SetConfig(ctx, v),
			})
		}
	}

	if config.Switch != nil {
		for _, v := range config.Switch {
			report.Switch = append(report.Switch, &ComponentReport{
				ID:    &v.ID,
				Error: t.Switch().SetConfig(ctx, v),
			})
		}
	}

	if config.Auth != nil {
		report.Auth = &ComponentReport{}
		report.Auth.Error = t.setAuth(ctx, config.Auth)
	}

	return report
}

// GetDeviceInfo returns information about the device.
func (t *Client) GetDeviceInfo(ctx context.Context) (*DeviceInfo, error) {

	method := Component + ".GetDeviceInfo"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: &method,
	})
	if err != nil {
		return nil, err
	}

	response := &DeviceInfoResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	if response.Result == nil {
		return nil, fmt.Errorf("Result is missing from response")
	}

	return response.Result, nil
}

// CheckForUpdate checks for new firmware version for the device and returns information about it.
// If no update is available returns empty JSON object as result.
func (t *Client) CheckForUpdate(ctx context.Context) (*UpdatesReport, error) {

	method := Component + ".CheckForUpdate"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: &method,
	})
	if err != nil {
		return nil, err
	}

	response := &CheckForUpdateResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	if response.Result == nil {
		return nil, fmt.Errorf("Result is missing from response")
	}

	return &UpdatesReport{
		Src:              response.Src,
		AvailableUpdates: response.Result,
	}, nil
}

// Update updates the firmware version of the device.
func (t *Client) Update(ctx context.Context, config *ShellyUpdateConfig) error {

	method := Component + ".Update"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: &method,
		Params: config,
	})
	if err != nil {
		return err
	}

	response := &SetConfigResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return err
	}

	if response.Error != nil {
		return response.Error
	}

	return nil
}

// FactoryReset resets the configuration to its default state
func (t *Client) FactoryReset(ctx context.Context) error {

	method := Component + ".FactoryReset"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: &method,
	})
	if err != nil {
		return err
	}

	response := &SetConfigResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return err
	}

	if response.Error != nil {
		return response.Error
	}

	return nil
}

// ResetWiFiConfig resets the WiFi configuration of the device
func (t *Client) ResetWiFiConfig(ctx context.Context) error {

	method := Component + ".ResetWiFiConfig"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: &method,
	})
	if err != nil {
		return err
	}

	response := &SetConfigResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return err
	}

	if response.Error != nil {
		return response.Error
	}

	return nil
}

// Reboot reboots the device
func (t *Client) Reboot(ctx context.Context) error {

	method := Component + ".Reboot"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: &method,
	})
	if err != nil {
		return err
	}

	response := &SetConfigResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return err
	}

	if response.Error != nil {
		return response.Error
	}

	return nil
}

func (t *Client) setAuth(ctx context.Context, config *ShellyAuthConfig) error {

	method := Component + ".SetAuth"

	raw := &RawShellyAuthConfig{}

	if config.Enable {

		if config.Pass == nil {
			return fmt.Errorf("Pass is required")
		}

		deviceInfo, err := t.GetDeviceInfo(ctx)
		if err != nil {
			return err
		}

		raw.User = ShellyUser
		raw.Realm = *deviceInfo.ID
		raw.Ha1 = *config.Pass

		hashInput := raw.User + ":" + raw.Realm + ":" + *config.Pass
		h := sha256.New()
		h.Write([]byte(hashInput))
		b := h.Sum(nil)

		raw.Ha1 = hex.EncodeToString(b)
	}

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: &method,
		Params: raw,
	})
	if err != nil {
		return err
	}

	response := &SetConfigResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return err
	}

	if response.Error != nil {
		return response.Error
	}

	return nil
}

func (t *Client) PutTLSClientCert(ctx context.Context, config *ShellyTLSConfig) error {

	method := Component + ".PutTLSClientCert"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: &method,
		Params: config,
	})
	if err != nil {
		return err
	}

	response := &SetConfigResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return err
	}

	if response.Error != nil {
		return response.Error
	}

	return nil
}

func (t *Client) PutTLSClientKey(ctx context.Context, config *ShellyTLSConfig) error {

	method := Component + ".PutTLSClientKey"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: &method,
		Params: config,
	})
	if err != nil {
		return err
	}

	response := &SetConfigResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return err
	}

	if response.Error != nil {
		return response.Error
	}

	return nil
}

func (t *Client) PutUserCA(ctx context.Context, config *ShellyTLSConfig) error {

	method := Component + ".PutUserCA"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: &method,
		Params: config,
	})
	if err != nil {
		return err
	}

	response := &SetConfigResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return err
	}

	if response.Error != nil {
		return response.Error
	}

	return nil
}

func (t *Client) Close() {
	if t._messageHandler != nil {
		t._messageHandler.Close()
	}
}
