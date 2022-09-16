package systemsProject

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/internal/app/checkdata"
	"server/internal/app/models"
	"sort"
)

type IncidentSystem struct {
	check  *checkdata.CheckData
	client *http.Client
	config *Config
}

func NewIncidentSystem(config *Config) *IncidentSystem {
	return &IncidentSystem{
		check:  &checkdata.CheckData{},
		client: &http.Client{},
		config: config,
	}
}

func (i *IncidentSystem) readIncident() ([]models.IncidentData, error) {
	var dataIncident []models.IncidentData
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
	if err = json.Unmarshal(data, &incidentMod); err != nil {
		return nil, err
	}
	for _, v := range *incidentMod {
		if err = i.check.CheckDataIncident(&v); err != nil {
			continue
		}
		dataIncident = append(dataIncident, v)
	}
	return dataIncident, nil
}

func (i *IncidentSystem) GetIncidentData() ([]models.IncidentData, error) {
	type Result struct {
		Payload []models.IncidentData
		Error   error
	}
	in := make(chan Result)
	go func() {
		incidentData, err := i.GetIncidentData()
		if err != nil {
			in <- Result{
				Payload: nil,
				Error:   err,
			}
		}
		sort.Slice(incidentData, func(i, j int) bool {
			return incidentData[i].Status < incidentData[j].Status
		})

		in <- Result{
			Payload: incidentData,
			Error:   nil,
		}
	}()
	result := <-in
	return result.Payload, result.Error
}
