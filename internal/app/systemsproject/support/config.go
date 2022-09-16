package support

type Config struct {
	SupportRequestAddr string `toml:"support_req"`
	SupportPersonal    int    `toml:"support_pers"`
	TickerPerHour      int    `toml:"ticket_per_hour"`
}

func NewConfig() *Config {
	return &Config{}
}
