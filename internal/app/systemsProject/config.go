package systemsProject

type Config struct {
	LenSmsData          int    `toml:"len_sms"`
	LenVoiceCallData    int    `toml:"len_voice"`
	LenEmailData        int    `toml:"len_email"`
	MMSRequestAddr      string `toml:"mms_req"`
	IncidentRequestAddr string `toml:"incident_req"`
	SupportRequestAddr  string `toml:"support_req"`
	SupportPersonal     int    `toml:"support_pers"`
	TickerPerHour       int    `toml:"ticket_per_hour"`
}

func NewConfig() *Config {
	return &Config{}
}
