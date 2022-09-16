package email

import "server/internal/app/models"

const (
	dEmail = "email.data"
)

type Email interface {
	readEmail() ([]models.EmailData, error)
	GetEmailData() (map[string][][]models.EmailData, error)
}
