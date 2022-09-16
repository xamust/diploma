package systemsproject

import (
	"server/internal/app/systemsproject/email"
	"server/internal/app/systemsproject/incident"
	"server/internal/app/systemsproject/mms"
	"server/internal/app/systemsproject/sms"
	"server/internal/app/systemsproject/support"
	"server/internal/app/systemsproject/voice"
)

type Config struct {
	SMS      *sms.Config
	MMS      *mms.Config
	Voice    *voice.Config
	Email    *email.Config
	Incident *incident.Config
	Support  *support.Config
}

func NewConfig() *Config {
	return &Config{
		Email:    email.NewConfig(),
		Incident: incident.NewConfig(),
		MMS:      mms.NewConfig(),
		SMS:      sms.NewConfig(),
		Voice:    voice.NewConfig(),
		Support:  support.NewConfig(),
	}
}
