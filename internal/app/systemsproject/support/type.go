package support

import "server/internal/app/models"

type Support interface {
	readSupport() (*[]models.SupportData, error)
	GetSupportData() ([]int, error)
}
