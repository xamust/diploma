package systemsProject

import "server/internal/app/models"

type SMS interface {
	readSMS() ([]models.SMSData, error)
	GetSMSData() ([][]models.SMSData, error)
}

type MMS interface {
	readMMS() ([]models.MMSData, error)
	GetMMSData() ([][]models.MMSData, error)
}

type Email interface {
	readEmail() ([]models.EmailData, error)
	GetEmailData() (map[string][][]models.EmailData, error)
}

type Voice interface {
	readVoice() ([]models.VoiceCallData, error)
	GetVoiceData() ([]models.VoiceCallData, error)
}

type Billing interface {
	readBilling() (*models.BillingData, error)
	GetBillingData() (*models.BillingData, error)
	calcDataBilling(input []string) (uint8, error)
}

type Support interface {
	readSupport() (*[]models.SupportData, error)
	GetSupportData() ([]int, error)
}

type Incidents interface {
	readIncident() ([]models.IncidentData, error)
	GetIncidentData() ([]models.IncidentData, error)
}
