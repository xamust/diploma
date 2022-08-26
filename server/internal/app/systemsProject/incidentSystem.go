package systemsProject

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/internal/app/checkdata"
	"server/internal/app/models"
)

type Incidents interface {
	GetIncidentData() ([]models.IncidentData, error)
}

type IncidentSystem struct {
	check  *checkdata.CheckData
	client *http.Client
	config *Config
}

func (i *IncidentSystem) GetIncidentData() ([]models.IncidentData, error) {

	//todo:request???!?!?!
	req, err := http.NewRequest(http.MethodGet, i.config.IncidentRequestAddr, nil)
	if err != nil {
		return nil, err
	}
	resp, err := i.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error! Response status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var incidentMod *[]models.IncidentData
	if err := json.Unmarshal(data, &incidentMod); err != nil {
		return nil, err
	}

	//todo: new var???
	var dataIncident []models.IncidentData
	for _, v := range *incidentMod {
		if err := i.CheckJSONIncident(&v); err != nil {
			continue
		}
		dataIncident = append(dataIncident, v)
	}

	return dataIncident, nil
}

func (i *IncidentSystem) CheckJSONIncident(v *models.IncidentData) error {
	return i.check.CheckDataIncident(v)
}
