package websocket

import (
	"github.com/jodydadescott/shelly-go-sdk/plus/types"
)

type Request = types.Request
type Response = types.Response
type Error = types.Error
type MessageHandlerFactory = types.MessageHandlerFactory
type MessageHandler = types.MessageHandler

type Status = types.WebsocketStatus
type Config = types.WebsocketConfig

// Result internal use only
type Result struct {
	RestartRequired *bool  `json:"restart_required,omitempty"`
	Error           *Error `json:"error,omitempty"`
}

// GetConfigResponse internal use only
type GetConfigResponse struct {
	Response
	Result *Config `json:"result,omitempty"`
	Params *Params `json:"params,omitempty"`
}

// SetConfigResponse internal use only
type SetConfigResponse struct {
	Response
	Result *Result `json:"result,omitempty"`
}

// GetStatusResponse internal use onlyå
type GetStatusResponse struct {
	Response
	Result *Status `json:"result,omitempty"`
}

// Params internal use only
type Params struct {
	Config *Config `json:"config,omitempty"`
}
