package mms

import "server/internal/app/models"

type MMS interface {
	readMMS() ([]models.MMSData, error)
	GetMMSData() ([][]models.MMSData, error)
}
