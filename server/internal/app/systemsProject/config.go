package systemsProject

type Config struct {
	LenSmsData          int    `toml:"len_sms"`
	LenVoiceCallData    int    `toml:"len_voice"`
	LenEmailData        int    `toml:"len_email"`
	MMSRequestAddr      string `toml:"mms_req"`
	IncidentRequestAddr string `toml:"incident_req"`
	SupportRequestAddr  string `toml:"support_req"`
}

func NewConfig() *Config {
	return &Config{}
}
