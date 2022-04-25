package collect

type Config struct {
	LenSmsData       int    `toml:"len_sms"`
	LenVoiceCallData int    `toml:"len_voice"`
	LenEmailData     int    `toml:"len_email"`
	DataFolder       string `toml:"data_folder"`
}

func NewConfig() *Config {
	return &Config{}
}
