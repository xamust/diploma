package systemsProject

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/internal/app/checkdata"
	"server/internal/app/models"
)

type Support interface {
	GetSupportData() (*[]models.SupportData, error)
}

type SupportService struct {
	check  *checkdata.CheckData
	client *http.Client
	config *Config
}

func (s *SupportService) GetSupportData() (*[]models.SupportData, error) {

	req, err := http.NewRequest(http.MethodGet, s.config.SupportRequestAddr, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
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

	var supportMod *[]models.SupportData
	if err := json.Unmarshal(data, &supportMod); err != nil {
		return nil, err
	}

	//todo: new var???
	var dataSupport []models.SupportData
	for _, v := range *supportMod {
		if err := s.checkJSONSupport(&v); err != nil {
			continue
		}
		dataSupport = append(dataSupport, v)
	}

	return &dataSupport, nil

}

func (s *SupportService) checkJSONSupport(v *models.SupportData) error {
	return s.check.CheckDataSupport(v)
}
