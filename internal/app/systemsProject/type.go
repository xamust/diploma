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
