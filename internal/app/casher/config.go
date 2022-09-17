package casher

type Config struct {
	StorageTimeout int `toml:"storage_time"`
}

func NewConfig() *Config {
	return &Config{}
}
