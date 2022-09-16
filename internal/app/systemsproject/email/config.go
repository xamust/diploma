package email

type Config struct {
	LenEmailData int `toml:"len_email"`
}

func NewConfig() *Config {
	return &Config{}
}
