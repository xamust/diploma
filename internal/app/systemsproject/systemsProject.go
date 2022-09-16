package systemsproject

import (
	"server/internal/app/models"
	"server/internal/app/systemsproject/billing"
	"server/internal/app/systemsproject/email"
	"server/internal/app/systemsproject/incident"
	"server/internal/app/systemsproject/mms"
	"server/internal/app/systemsproject/sms"
	"server/internal/app/systemsproject/support"
	"server/internal/app/systemsproject/voice"
)

type SystemsProject struct {
	Config           *Config
	ParsingDataFiles map[string]string
}

// GetResultData get result data...
func (s *SystemsProject) GetResultData() (*models.ResultSetT, error) {

	sms := sms.NewSMSSystem(s.ParsingDataFiles, s.Config.SMS)
	smsData, err := sms.GetSMSData()
	if err != nil {
		return nil, err
	}

	mms := mms.NewMMSSystem(s.Config.MMS)
	mmsData, err := mms.GetMMSData()
	if err != nil {
		return nil, err
	}

	voice := voice.NewVoiceSystem(s.ParsingDataFiles, s.Config.Voice)
	voiceData, err := voice.GetVoiceData()
	if err != nil {
		return nil, err
	}

	email := email.NewEmailSystem(s.ParsingDataFiles, s.Config.Email)
	emailData, err := email.GetEmailData()
	if err != nil {
		return nil, err
	}

	billinig := billing.NewBillingSystem(s.ParsingDataFiles)
	billingData, err := billinig.GetBillingData()
	if err != nil {
		return nil, err
	}

	support := support.NewSupportSystem(s.Config.Support)
	supportData, err := support.GetSupportData()
	if err != nil {
		return nil, err
	}

	incident := incident.NewIncidentSystem(s.Config.Incident)
	incidentData, err := incident.GetIncidentData()
	if err != nil {
		return nil, err
	}
	return &models.ResultSetT{
		SMS:       smsData,
		MMS:       mmsData,
		VoiceCall: voiceData,
		Email:     emailData,
		Billing:   *billingData,
		Support:   supportData,
		Incidents: incidentData,
	}, nil
}
