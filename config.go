package shelly

import "time"

type Config struct {
	Hostname     string
	Password     string
	DebugEnabled bool
	SendTimeout  time.Duration
}

func (t *Config) GetHostname() string {
	return t.Hostname
}

func (t *Config) GetPassword() string {
	return t.Password
}

func (t *Config) IsDebugEnabled() bool {
	return t.DebugEnabled
}

func (t *Config) GetSendTimeout() time.Duration {
	return t.SendTimeout
}
