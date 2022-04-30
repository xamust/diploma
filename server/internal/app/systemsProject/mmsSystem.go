package systemsProject

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"server/internal/app/models"
)

type MMSSystem struct {
	check  *CheckData
	client *http.Client
	logger *logrus.Logger
	config *Config
}

func (m *MMSSystem) ReadMMS() ([]models.MMSData, error) {
	return m.GetMMSData()
}

func (m *MMSSystem) GetMMSData() ([]models.MMSData, error) {

	req, err := http.NewRequest(http.MethodGet, m.config.MMSRequestAddr, nil)
	if err != nil {
		m.logger.Error(err.Error())
		return nil, err
	}
	resp, err := m.client.Do(req)
	if err != nil {
		m.logger.Error(err.Error())
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error! Response status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		m.logger.Error(err.Error())
		return nil, err
	}

	var mmsMod *[]models.MMSData
	if err := json.Unmarshal(data, &mmsMod); err != nil {
		m.logger.Error(err.Error())
		return nil, err
	}

	//todo: new var???
	var dataMMS []models.MMSData
	for _, v := range *mmsMod {
		if err := m.CheckJSONMMS(&v); err != nil {
			m.logger.Warn(err)
			continue
		}
		dataMMS = append(dataMMS, v)
	}
	m.logger.Print("MMS data uploading complete!")
	return dataMMS, nil
}

func (m *MMSSystem) CheckJSONMMS(v *models.MMSData) error {
	return m.check.CheckDataMMS(v)
}
