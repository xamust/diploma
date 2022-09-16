package sms

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"server/internal/app/checkdata"
	"server/internal/app/models"
	"sort"
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

	var SMSSlice []models.SMSData
	data, err := os.ReadFile(s.fileName[dSMS])
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(bytes.NewReader(data))
	r.Comma = ';'
	for {
		dataSMS, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		if err = s.check.CheckDataSMS(dataSMS, s.config.LenSMSData); err != nil {
			continue
		}

		SMSSlice = append(SMSSlice, models.SMSData{
			Country:      dataSMS[0],
			Bandwidth:    dataSMS[1],
			ResponseTime: dataSMS[2],
			Provider:     dataSMS[3],
		})
	}

	return SMSSlice, nil
}

func (s *SMSSystem) GetSMSData() ([][]models.SMSData, error) {
	type Result struct {
		Payload [][]models.SMSData
		Error   error
	}
	inSMS := make(chan Result)

	go func() {
		dataSMS, err := s.readSMS()
		if err != nil {
			inSMS <- Result{
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
		inSMS <- Result{
			Payload: [][]models.SMSData{dataSMS, dataSMSDouble},
			Error:   nil,
		}
		close(inSMS)
	}()
	result := <-inSMS
	return result.Payload, result.Error
}
