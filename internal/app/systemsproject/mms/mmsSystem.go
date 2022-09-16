package mms

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/internal/app/checkdata"
	"server/internal/app/models"
	"sort"
)

type MMSSystem struct {
	check  *checkdata.CheckData
	client *http.Client
	config *Config
}

func NewMMSSystem(config *Config) *MMSSystem {
	return &MMSSystem{
		check:  &checkdata.CheckData{},
		client: &http.Client{},
		config: config,
	}
}

func (m *MMSSystem) readMMS() ([]models.MMSData, error) {

	req, err := http.NewRequest(http.MethodGet, m.config.MMSRequestAddr, nil)
	if err != nil {
		return nil, err
	}
	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error! response status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var mmsMod *[]models.MMSData
	if err = json.Unmarshal(data, &mmsMod); err != nil {
		return nil, err
	}

	var dataMMS []models.MMSData
	for _, v := range *mmsMod {
		if err = m.check.CheckDataMMS(&v, m.config.LenMMSData); err != nil {
			continue
		}
		dataMMS = append(dataMMS, v)
	}
	return dataMMS, nil
}

// GetMMSData mms system...
func (m *MMSSystem) GetMMSData() ([][]models.MMSData, error) {
	type Result struct {
		Payload [][]models.MMSData
		Error   error
	}

	in := make(chan Result)
	go func() {
		dataMMS, err := m.readMMS()
		if err != nil {
			in <- Result{
				Payload: nil,
				Error:   err,
			}
		}
		models.FullCountryNameMMS(dataMMS)
		// copy slice
		dataMMSDouble := make([]models.MMSData, len(dataMMS))
		copy(dataMMSDouble, dataMMS)
		//sort by provider
		sort.Slice(dataMMS, func(i, j int) bool {
			return dataMMS[i].Provider < dataMMS[j].Provider
		})
		//sort by country name
		sort.Slice(dataMMSDouble, func(i, j int) bool {
			return dataMMSDouble[i].Country < dataMMSDouble[j].Country
		})
		in <- Result{
			Payload: [][]models.MMSData{dataMMS, dataMMSDouble},
			Error:   nil,
		}
	}()
	result := <-in
	return result.Payload, result.Error
}
