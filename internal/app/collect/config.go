package collect

// Config настройка конфигурации Collect
type Config struct {
	DataFolder string `toml:"data_folder"`
}

// NewConfig инициализация конфигурации
func NewConfig() *Config {
	return &Config{}
}
