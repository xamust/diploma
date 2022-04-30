package systemsProject

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"server/internal/app/models"
)

type IncidentSystem struct {
	check  *CheckData
	client *http.Client
	logger *logrus.Logger
	config *Config
}

func (i *IncidentSystem) ReadIncident() ([]models.IncidentData, error) {
	return i.GetIncidentData()
}

func (i *IncidentSystem) GetIncidentData() ([]models.IncidentData, error) {

	req, err := http.NewRequest(http.MethodGet, i.config.IncidentRequestAddr, nil)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}
	resp, err := i.client.Do(req)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error! Response status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	var incidentMod *[]models.IncidentData
	if err := json.Unmarshal(data, &incidentMod); err != nil {
		i.logger.Error(err.Error())
		return nil, err
	}

	//todo: new var???
	var dataIncident []models.IncidentData
	for _, v := range *incidentMod {
		if err := i.CheckJSONIncident(&v); err != nil {
			i.logger.Warn(err)
			continue
		}
		dataIncident = append(dataIncident, v)
	}

	i.logger.Print("Incident data uploading complete!")
	return dataIncident, nil
}

func (i *IncidentSystem) CheckJSONIncident(v *models.IncidentData) error {
	return i.check.CheckDataIncident(v)
}
