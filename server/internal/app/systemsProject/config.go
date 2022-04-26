package systemsProject

type Config struct {
	LenSmsData       int `toml:"len_sms"`
	LenVoiceCallData int `toml:"len_voice"`
	LenEmailData     int `toml:"len_email"`
}

func NewConfig() *Config {
	return &Config{}
}
