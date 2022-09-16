package systemsProject

import (
	"os"
	"server/internal/app/checkdata"
	"server/internal/app/models"
	"sort"
	"strings"
)

type SMSSystem struct {
	check    *checkdata.CheckData
	config   *Config
	fileName map[string]string
}

func NewSMSSystem(fileName map[string]string, config *Config) *SMSSystem {
	return &SMSSystem{
		check:    &checkdata.CheckData{},
		config:   config,
		fileName: fileName,
	}
}

func (s *SMSSystem) readSMS() ([]models.SMSData, error) {

	//init slice SMSData
	SMSSlice := &[]models.SMSData{}

	data, err := os.ReadFile(s.fileName["sms.data"])
	if err != nil {
		return nil, err
	}

	//TODO:need another way to '\n'...
	for _, v := range strings.Split(string(data), "\n") {
		dataSMS := strings.Split(v, ";")
		if err = s.check.CheckDataSMS(dataSMS, s.config.LenSMSData); err != nil {
			continue
		}

		*SMSSlice = append(*SMSSlice, models.SMSData{
			Country:      dataSMS[0],
			Bandwidth:    dataSMS[1],
			ResponseTime: dataSMS[2],
			Provider:     dataSMS[3],
		})
	}

	return *SMSSlice, nil
}

func (s *SMSSystem) GetSMSData() ([][]models.SMSData, error) {

	type Result struct {
		Payload [][]models.SMSData
		Error   error
	}

	in := make(chan Result)

	go func() {
		dataSMS, err := s.readSMS()
		if err != nil {
			in <- Result{
				Payload: nil,
				Error:   err,
			}
		}
		models.FullCountryNameSMS(dataSMS)
		dataSMSDouble := make([]models.SMSData, len(dataSMS))
		copy(dataSMSDouble, dataSMS)
		sort.Slice(dataSMS, func(i, j int) bool {
			return dataSMS[i].Provider < dataSMS[j].Provider
		})
		sort.Slice(dataSMSDouble, func(i, j int) bool {
			return dataSMSDouble[i].Country < dataSMSDouble[j].Country
		})
		in <- Result{
			Payload: [][]models.SMSData{dataSMS, dataSMSDouble},
			Error:   nil,
		}
	}()
	result := <-in
	return result.Payload, result.Error
}
