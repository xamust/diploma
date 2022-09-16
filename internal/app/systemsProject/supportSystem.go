package systemsProject

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"server/internal/app/checkdata"
	"server/internal/app/models"
)

type SupportService struct {
	check  *checkdata.CheckData
	client *http.Client
	config *Config
}

func NewSupportSystem(config *Config) *SupportService {
	return &SupportService{
		check:  &checkdata.CheckData{},
		client: &http.Client{},
		config: config,
	}
}

func (s *SupportService) readSupport() (*[]models.SupportData, error) {

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
	if err = json.Unmarshal(data, &supportMod); err != nil {
		return nil, err
	}
	var dataSupport []models.SupportData
	for _, v := range *supportMod {
		if err := s.check.CheckDataSupport(&v); err != nil {
			continue
		}
		dataSupport = append(dataSupport, v)
	}
	return &dataSupport, nil
}

func (s *SupportService) GetSupportData() ([]int, error) {
	type Result struct {
		Payload []int
		Error   error
	}
	in := make(chan Result)
	go func() {
		supportData, err := s.readSupport()
		if err != nil {
			in <- Result{
				Payload: nil,
				Error:   err,
			}
		}
		var countLoad, countTime, ticketCount float64
		calculatedTime := 60 / float64(s.config.TickerPerHour)
		for _, v := range *supportData {
			calcTime := float64(v.ActiveTickets) * calculatedTime
			countTime += calcTime
			ticketCount += float64(v.ActiveTickets)
			switch {
			case calcTime < 9:
				countLoad += 1
			case calcTime >= 9 && calcTime <= 16:
				countLoad += 2
			case calcTime > 16:
				countLoad += 3
			}
		}
		in <- Result{
			Payload: []int{int(math.Round(countLoad / float64(len(*supportData)))), 60 / s.config.TickerPerHour * int(ticketCount)},
			Error:   nil,
		}

	}()
	result := <-in
	return result.Payload, result.Error
}
