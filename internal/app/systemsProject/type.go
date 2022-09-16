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
