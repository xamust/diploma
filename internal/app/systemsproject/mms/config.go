package mms

type Config struct {
	LenMMSData     int    `toml:"len_mms"`
	MMSRequestAddr string `toml:"mms_req"`
}

func NewConfig() *Config {
	return &Config{}
}
