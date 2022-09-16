package sms

import "server/internal/app/models"

const (
	dSMS = "sms.data"
)

type SMS interface {
	readSMS() ([]models.SMSData, error)
	GetSMSData() ([][]models.SMSData, error)
}
