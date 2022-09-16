package incident

type Config struct {
	IncidentRequestAddr string `toml:"incident_req"`
}

func NewConfig() *Config {
	return &Config{}
}
