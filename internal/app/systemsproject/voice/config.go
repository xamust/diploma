package voice

type Config struct {
	LenVoiceCallData int `toml:"len_voice"`
}

func NewConfig() *Config {
	return &Config{}
}
