package systemsProject

import (
	"server/internal/app/models"
)

type SystemsProject struct {
	Config           *Config
	ParsingDataFiles map[string]string
}

// GetResultData get result data...
func (s *SystemsProject) GetResultData() (*models.ResultSetT, error) {

	sms := NewSMSSystem(s.ParsingDataFiles, s.Config)
	smsData, err := sms.GetSMSData()
	if err != nil {
		return nil, err
	}

	mms := NewMMSSystem(s.Config)
	mmsData, err := mms.GetMMSData()
	if err != nil {
		return nil, err
	}

	voice := NewVoiceSystem(s.ParsingDataFiles, s.Config)
	voiceData, err := voice.GetVoiceData()
	if err != nil {
		return nil, err
	}

	email := NewEmailSystem(s.ParsingDataFiles, s.Config)
	emailData, err := email.GetEmailData()
	if err != nil {
		return nil, err
	}

	billinig := NewBillingSystem(s.ParsingDataFiles)
	billingData, err := billinig.GetBillingData()
	if err != nil {
		return nil, err
	}

	support := NewSupportSystem(s.Config)
	supportData, err := support.GetSupportData()
	if err != nil {
		return nil, err
	}

	incident := NewIncidentSystem(s.Config)
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
