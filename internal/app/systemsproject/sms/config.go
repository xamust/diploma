package sms

type Config struct {
	LenSMSData int `toml:"len_sms"`
}

func NewConfig() *Config {
	return &Config{}
}
