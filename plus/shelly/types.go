package shelly

import (
	"github.com/jodydadescott/shelly-go-sdk/plus/types"
)

type MessageHandlerFactory = types.MessageHandlerFactory
type MessageHandler = types.MessageHandler

type Config = types.ShellyConfig
type AuthResponse = types.AuthResponse
type AuthRequest = types.AuthRequest
type BluetoothStatus = types.BluetoothStatus
type BluetoothConfig = types.BluetoothConfig
type BluetoothRPC = types.BluetoothRPC
type BluetoothObserver = types.BluetoothObserver
type CloudStatus = types.CloudStatus
type CloudConfig = types.CloudConfig
type FirmwareStatus = types.FirmwareStatus
type Request = types.Request
type Response = types.Response
type Error = types.Error
type EthernetStatus = types.EthernetStatus
type EthernetConfig = types.EthernetConfig
type InputStatus = types.InputStatus
type InputConfig = types.InputConfig
type LightStatus = types.LightStatus
type LightConfig = types.LightConfig
type MqttStatus = types.MqttStatus
type MqttConfig = types.MqttConfig
type ShellyStatus = types.ShellyStatus
type ShellyReport = types.ShellyReport
type ComponentReport = types.ComponentReport
type ShellyRPCMethods = types.ShellyRPCMethods
type ShellyConfig = types.ShellyConfig
type DeviceInfo = types.DeviceInfo
type ShellyUpdateConfig = types.ShellyUpdateConfig
type ShellyAuthConfig = types.ShellyAuthConfig
type ShellyUserCAConfig = types.ShellyUserCAConfig
type ShellyTLSClientCertConfig = types.ShellyTLSClientCertConfig
type ShellyTLSClientKeyConfig = types.ShellyTLSClientKeyConfig
type UpdatesReport = types.UpdatesReport
type SwitchStatus = types.SwitchStatus
type SwitchAenergy = types.SwitchAenergy
type SwitchTemperature = types.SwitchTemperature
type SwitchConfig = types.SwitchConfig
type SystemStatus = types.SystemStatus
type SystemAvailableUpdates = types.SystemAvailableUpdates
type SystemWakeupReason = types.SystemWakeupReason
type SystemConfig = types.SystemConfig
type SystemDevice = types.SystemDevice
type SystemLocation = types.SystemLocation
type SystemDebug = types.SystemDebug
type SystemMqtt = types.SystemMqtt
type SystemWebsocket = types.SystemWebsocket
type SystemUDP = types.SystemUDP
type SystemUIData = types.SystemUIData
type SystemRPCUDP = types.SystemRPCUDP
type SystemSntp = types.SystemSntp
type Webhook = types.Webhook
type WebhookConfig = types.WebhookConfig
type WebsocketStatus = types.WebsocketStatus
type WebsocketConfig = types.WebsocketConfig
type WifiStatus = types.WifiStatus
type WifiConfig = types.WifiConfig
type WifiAPConfig = types.WifiAPConfig
type WifiRangeExtenderConfig = types.WifiRangeExtenderConfig
type WifiSTAConfig = types.WifiSTAConfig
type WifiRoamConfig = types.WifiRoamConfig
type WifiNet = types.WifiNet
type WifiAPClient = types.WifiAPClient
type WifiScanResults = types.WifiScanResults
type WifiAPClients = types.WifiAPClients

// Result internal use only
type Result struct {
	RestartRequired *bool  `json:"restart_required,omitempty"`
	Error           *Error `json:"error,omitempty"`
}

// SetConfigResponse internal use only
type SetConfigResponse struct {
	Response
	Result *Result `json:"result,omitempty"`
}

// GetConfigResponse internal use only
type GetConfigResponse struct {
	Response
	Result *RawShellyConfig `json:"result,omitempty"`
}

// GetStatusResponse internal use only
type GetStatusResponse struct {
	Response
	Result *RawShellyStatus `json:"result,omitempty"`
}

// DeviceInfoResponse internal use only
type DeviceInfoResponse struct {
	Response
	Result *DeviceInfo `json:"result,omitempty"`
}

// CheckForUpdateResponse Shelly component object
type CheckForUpdateResponse struct {
	Response
	Result *SystemAvailableUpdates `json:"result,omitempty"`
}

// ListMethodsResponse internal use only
type ListMethodsResponse struct {
	Response
	Result *ShellyRPCMethods `json:"result,omitempty"`
}

// RawShellyStatus internal use only
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Shelly
type RawShellyStatus struct {
	Bluetooth *BluetoothStatus `json:"ble,omitempty" yaml:"ble,omitempty"`
	Cloud     *CloudStatus     `json:"cloud,omitempty" yaml:"cloud,omitempty"`
	Mqtt      *MqttStatus      `json:"mqtt,omitempty" yaml:"mqtt,omitempty"`
	Ethernet  *EthernetStatus  `json:"eth,omitempty" yaml:"eth,omitempty"`
	System    *SystemStatus    `json:"sys,omitempty" yaml:"sys,omitempty"`
	Wifi      *WifiStatus      `json:"wifi,omitempty" yaml:"wifi,omitempty"`
	Websocket *WebsocketStatus `json:"ws,omitempty" yaml:"ws,omitempty"`
	Light0    *LightStatus     `json:"light:0,omitempty" yaml:"light:0,omitempty"`
	Light1    *LightStatus     `json:"light:1,omitempty" yaml:"light:1,omitempty"`
	Light2    *LightStatus     `json:"light:2,omitempty" yaml:"light:2,omitempty"`
	Light3    *LightStatus     `json:"light:3,omitempty" yaml:"light:3,omitempty"`
	Light4    *LightStatus     `json:"light:4,omitempty" yaml:"light:4,omitempty"`
	Light5    *LightStatus     `json:"light:5,omitempty" yaml:"light:5,omitempty"`
	Light6    *LightStatus     `json:"light:6,omitempty" yaml:"light:6,omitempty"`
	Light7    *LightStatus     `json:"light:7,omitempty" yaml:"light:7,omitempty"`
	Input0    *InputStatus     `json:"input:0,omitempty" yaml:"input:0,omitempty"`
	Input1    *InputStatus     `json:"input:1,omitempty" yaml:"input:1,omitempty"`
	Input2    *InputStatus     `json:"input:2,omitempty" yaml:"input:2,omitempty"`
	Input3    *InputStatus     `json:"input:3,omitempty" yaml:"input:3,omitempty"`
	Input4    *InputStatus     `json:"input:4,omitempty" yaml:"input:4,omitempty"`
	Input5    *InputStatus     `json:"input:5,omitempty" yaml:"input:5,omitempty"`
	Input6    *InputStatus     `json:"input:6,omitempty" yaml:"input:6,omitempty"`
	Input7    *InputStatus     `json:"input:7,omitempty" yaml:"input:7,omitempty"`
	Switch0   *SwitchStatus    `json:"switch:0,omitempty" yaml:"switch:0,omitempty"`
	Switch1   *SwitchStatus    `json:"switch:1,omitempty" yaml:"switch:1,omitempty"`
	Switch2   *SwitchStatus    `json:"switch:2,omitempty" yaml:"switch:2,omitempty"`
	Switch3   *SwitchStatus    `json:"switch:3,omitempty" yaml:"switch:3,omitempty"`
	Switch4   *SwitchStatus    `json:"switch:4,omitempty" yaml:"switch:4,omitempty"`
	Switch5   *SwitchStatus    `json:"switch:5,omitempty" yaml:"switch:5,omitempty"`
	Switch6   *SwitchStatus    `json:"switch:6,omitempty" yaml:"switch:6,omitempty"`
	Switch7   *SwitchStatus    `json:"switch:7,omitempty" yaml:"switch:7,omitempty"`
}

func (t *RawShellyStatus) convert() *ShellyStatus {

	c := &ShellyStatus{
		Bluetooth: t.Bluetooth,
		Cloud:     t.Cloud,
		Mqtt:      t.Mqtt,
		Ethernet:  t.Ethernet,
		System:    t.System,
		Wifi:      t.Wifi,
	}

	if t.Light0 != nil {
		c.Light = append(c.Light, t.Light0)
	}
	if t.Light1 != nil {
		c.Light = append(c.Light, t.Light1)
	}
	if t.Light2 != nil {
		c.Light = append(c.Light, t.Light2)
	}
	if t.Light3 != nil {
		c.Light = append(c.Light, t.Light3)
	}
	if t.Light4 != nil {
		c.Light = append(c.Light, t.Light4)
	}
	if t.Light5 != nil {
		c.Light = append(c.Light, t.Light5)
	}
	if t.Light6 != nil {
		c.Light = append(c.Light, t.Light6)
	}
	if t.Light7 != nil {
		c.Light = append(c.Light, t.Light7)
	}

	if t.Input0 != nil {
		c.Input = append(c.Input, t.Input0)
	}
	if t.Input1 != nil {
		c.Input = append(c.Input, t.Input1)
	}
	if t.Input2 != nil {
		c.Input = append(c.Input, t.Input2)
	}
	if t.Input3 != nil {
		c.Input = append(c.Input, t.Input3)
	}
	if t.Input4 != nil {
		c.Input = append(c.Input, t.Input4)
	}
	if t.Input5 != nil {
		c.Input = append(c.Input, t.Input5)
	}
	if t.Input6 != nil {
		c.Input = append(c.Input, t.Input6)
	}
	if t.Input7 != nil {
		c.Input = append(c.Input, t.Input7)
	}

	if t.Switch0 != nil {
		c.Switch = append(c.Switch, t.Switch0)
	}
	if t.Switch1 != nil {
		c.Switch = append(c.Switch, t.Switch1)
	}
	if t.Switch2 != nil {
		c.Switch = append(c.Switch, t.Switch2)
	}
	if t.Switch3 != nil {
		c.Switch = append(c.Switch, t.Switch3)
	}
	if t.Switch4 != nil {
		c.Switch = append(c.Switch, t.Switch4)
	}
	if t.Switch5 != nil {
		c.Switch = append(c.Switch, t.Switch5)
	}
	if t.Switch6 != nil {
		c.Switch = append(c.Switch, t.Switch6)
	}
	if t.Switch7 != nil {
		c.Switch = append(c.Switch, t.Switch7)
	}

	return c
}

// RawShellyConfig internal use only
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Shelly#configuration
type RawShellyConfig struct {
	Bluetooth *BluetoothConfig `json:"ble,omitempty" yaml:"ble,omitempty"`
	Cloud     *CloudConfig     `json:"cloud,omitempty" yaml:"cloud,omitempty"`
	Mqtt      *MqttConfig      `json:"mqtt,omitempty" yaml:"mqtt,omitempty"`
	Ethernet  *EthernetConfig  `json:"eth,omitempty" yaml:"eth,omitempty"`
	System    *SystemConfig    `json:"sys,omitempty" yaml:"sys,omitempty"`
	Wifi      *WifiConfig      `json:"wifi,omitempty" yaml:"wifi,omitempty"`
	Websocket *WebsocketConfig `json:"ws,omitempty" yaml:"ws,omitempty"`
	Light0    *LightConfig     `json:"light:0,omitempty" yaml:"light:0,omitempty"`
	Light1    *LightConfig     `json:"light:1,omitempty" yaml:"light:1,omitempty"`
	Light2    *LightConfig     `json:"light:2,omitempty" yaml:"light:2,omitempty"`
	Light3    *LightConfig     `json:"light:3,omitempty" yaml:"light:3,omitempty"`
	Light4    *LightConfig     `json:"light:4,omitempty" yaml:"light:4,omitempty"`
	Light5    *LightConfig     `json:"light:5,omitempty" yaml:"light:5,omitempty"`
	Light6    *LightConfig     `json:"light:6,omitempty" yaml:"light:6,omitempty"`
	Light7    *LightConfig     `json:"light:7,omitempty" yaml:"light:7,omitempty"`
	Input0    *InputConfig     `json:"input:0,omitempty" yaml:"input:0,omitempty"`
	Input1    *InputConfig     `json:"input:1,omitempty" yaml:"input:1,omitempty"`
	Input2    *InputConfig     `json:"input:2,omitempty" yaml:"input:2,omitempty"`
	Input3    *InputConfig     `json:"input:3,omitempty" yaml:"input:3,omitempty"`
	Input4    *InputConfig     `json:"input:4,omitempty" yaml:"input:4,omitempty"`
	Input5    *InputConfig     `json:"input:5,omitempty" yaml:"input:5,omitempty"`
	Input6    *InputConfig     `json:"input:6,omitempty" yaml:"input:6,omitempty"`
	Input7    *InputConfig     `json:"input:7,omitempty" yaml:"input:7,omitempty"`
	Switch0   *SwitchConfig    `json:"switch:0,omitempty" yaml:"switch:0,omitempty"`
	Switch1   *SwitchConfig    `json:"switch:1,omitempty" yaml:"switch:1,omitempty"`
	Switch2   *SwitchConfig    `json:"switch:2,omitempty" yaml:"switch:2,omitempty"`
	Switch3   *SwitchConfig    `json:"switch:3,omitempty" yaml:"switch:3,omitempty"`
	Switch4   *SwitchConfig    `json:"switch:4,omitempty" yaml:"switch:4,omitempty"`
	Switch5   *SwitchConfig    `json:"switch:5,omitempty" yaml:"switch:5,omitempty"`
	Switch6   *SwitchConfig    `json:"switch:6,omitempty" yaml:"switch:6,omitempty"`
	Switch7   *SwitchConfig    `json:"switch:7,omitempty" yaml:"switch:7,omitempty"`
}

func (t *RawShellyConfig) convert() *ShellyConfig {

	c := &ShellyConfig{
		Bluetooth: t.Bluetooth,
		Cloud:     t.Cloud,
		Mqtt:      t.Mqtt,
		Ethernet:  t.Ethernet,
		System:    t.System,
		Wifi:      t.Wifi,
		Websocket: t.Websocket,
	}

	if t.Light0 != nil {
		c.Light = append(c.Light, t.Light0)
	}
	if t.Light1 != nil {
		c.Light = append(c.Light, t.Light1)
	}
	if t.Light2 != nil {
		c.Light = append(c.Light, t.Light2)
	}
	if t.Light3 != nil {
		c.Light = append(c.Light, t.Light3)
	}
	if t.Light4 != nil {
		c.Light = append(c.Light, t.Light4)
	}
	if t.Light5 != nil {
		c.Light = append(c.Light, t.Light5)
	}
	if t.Light6 != nil {
		c.Light = append(c.Light, t.Light6)
	}
	if t.Light7 != nil {
		c.Light = append(c.Light, t.Light7)
	}

	if t.Input0 != nil {
		c.Input = append(c.Input, t.Input0)
	}
	if t.Input1 != nil {
		c.Input = append(c.Input, t.Input1)
	}
	if t.Input2 != nil {
		c.Input = append(c.Input, t.Input2)
	}
	if t.Input3 != nil {
		c.Input = append(c.Input, t.Input3)
	}
	if t.Input4 != nil {
		c.Input = append(c.Input, t.Input4)
	}
	if t.Input5 != nil {
		c.Input = append(c.Input, t.Input5)
	}
	if t.Input6 != nil {
		c.Input = append(c.Input, t.Input6)
	}
	if t.Input7 != nil {
		c.Input = append(c.Input, t.Input7)
	}

	if t.Switch0 != nil {
		c.Switch = append(c.Switch, t.Switch0)
	}
	if t.Switch1 != nil {
		c.Switch = append(c.Switch, t.Switch1)
	}
	if t.Switch2 != nil {
		c.Switch = append(c.Switch, t.Switch2)
	}
	if t.Switch3 != nil {
		c.Switch = append(c.Switch, t.Switch3)
	}
	if t.Switch4 != nil {
		c.Switch = append(c.Switch, t.Switch4)
	}
	if t.Switch5 != nil {
		c.Switch = append(c.Switch, t.Switch5)
	}
	if t.Switch6 != nil {
		c.Switch = append(c.Switch, t.Switch6)
	}
	if t.Switch7 != nil {
		c.Switch = append(c.Switch, t.Switch7)
	}

	return c
}

// ShellyAuthConfig internal use only
// https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/Shelly#configuration
type RawShellyAuthConfig struct {
	// User is used by the following methods:
	// SetAuth: Must be set to admin. Only one user is supported. Required
	User string `json:"user,omitempty" yaml:"user,omitempty"`
	// Realm is used by the following methods:
	// SetAuth : Must be the id of the device. Only one realm is supported. Required
	Realm string `json:"realm,omitempty" yaml:"realm,omitempty"`
	// Ha1 is used by the following methods:
	// SetAuth : "user:realm:password" encoded in SHA256 (null to disable authentication). Required
	Ha1 string `json:"ha1,omitempty" yaml:"ha1,omitempty"`
}

// RawShellyTLSConfig internal use only
type RawShellyTLSConfig struct {
	// Data is used by the following methods:
	// PutUserCA : Contents of the PEM file (null if you want to delete the existing data). Required
	// PutTLSClientCert : Contents of the client.crt file (null if you want to delete the existing data). Required
	// PutTLSClientKey : Contents of the client.key file (null if you want to delete the existing data). Required
	Data *string `json:"data,omitempty" yaml:"data,omitempty"`
	// Append is used by the following methods:
	// PutUserCA : true if more data will be appended afterwards, default false.
	// PutTLSClientCert : true if more data will be appended afterwards, default false
	// PutTLSClientKey : true if more data will be appended afterwards, default false
	Append *bool `json:"append,omitempty" yaml:"append,omitempty"`
}
