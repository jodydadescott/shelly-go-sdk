package shelly

type Config struct {
	Hostname     string
	Password     string
	DebugEnabled bool
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
