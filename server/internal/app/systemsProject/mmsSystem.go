package systemsProject

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/internal/app/checkdata"
	"server/internal/app/models"
)

type MMS interface {
	GetMMSData() ([]models.MMSData, error)
}

type MMSSystem struct {
	check  *checkdata.CheckData
	client *http.Client
	config *Config
}

func (m *MMSSystem) GetMMSData() ([]models.MMSData, error) {

	req, err := http.NewRequest(http.MethodGet, m.config.MMSRequestAddr, nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(req)
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

	var mmsMod *[]models.MMSData
	if err := json.Unmarshal(data, &mmsMod); err != nil {
		return nil, err
	}

	//todo: new var???
	var dataMMS []models.MMSData
	for _, v := range *mmsMod {
		if err := m.CheckJSONMMS(&v); err != nil {
			continue
		}
		dataMMS = append(dataMMS, v)
	}
	return dataMMS, nil
}

func (m *MMSSystem) CheckJSONMMS(v *models.MMSData) error {
	return m.check.CheckDataMMS(v)
}
