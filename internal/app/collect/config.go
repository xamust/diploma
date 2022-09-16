package collect

type Config struct {
	DataFolder string `toml:"data_folder"`
}

func NewConfig() *Config {
	return &Config{}
}
