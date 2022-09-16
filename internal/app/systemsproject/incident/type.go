package incident

import "server/internal/app/models"

type Incidents interface {
	readIncident() ([]models.IncidentData, error)
	GetIncidentData() ([]models.IncidentData, error)
}
