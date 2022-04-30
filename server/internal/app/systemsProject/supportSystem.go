package systemsProject

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"server/internal/app/models"
)

type SupportService struct {
	check  *CheckData
	client *http.Client
	logger *logrus.Logger
	config *Config
}

func (s *SupportService) readSupport() (*[]models.SupportData, error) {
	return s.GetSupportData()
}

func (s *SupportService) GetSupportData() (*[]models.SupportData, error) {

	req, err := http.NewRequest(http.MethodGet, s.config.SupportRequestAddr, nil)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error! Response status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var supportMod *[]models.SupportData
	if err := json.Unmarshal(data, &supportMod); err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	//todo: new var???
	var dataSupport []models.SupportData
	for _, v := range *supportMod {
		if err := s.checkJSONSupport(&v); err != nil {
			s.logger.Warn(err.Error())
			continue
		}
		dataSupport = append(dataSupport, v)
	}

	s.logger.Print("Support data uploading complete!")
	return &dataSupport, nil

}

func (s *SupportService) checkJSONSupport(v *models.SupportData) error {
	return s.check.CheckDataSupport(v)
}
