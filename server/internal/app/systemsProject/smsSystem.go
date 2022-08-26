package systemsProject

import (
	"io/ioutil"
	"server/internal/app/checkdata"
	"server/internal/app/models"
	"strings"
)

type SMS interface {
	ReadSMS() ([]models.SMSData, error)
}

type SMSSystem struct {
	check    *checkdata.CheckData
	fileName map[string]string
}

func (s *SMSSystem) ReadSMS() ([]models.SMSData, error) {

	//init slice SMSData
	SMSSlice := &[]models.SMSData{}

	data, err := ioutil.ReadFile(s.fileName["sms.data"])
	if err != nil {
		return nil, err
	}

	//TODO:need another way to '\n'...
	for _, v := range strings.Split(string(data), "\n") {
		dataSMS := strings.Split(v, ";")
		if err := s.checkSMSData(dataSMS); err != nil {
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

func (s *SMSSystem) checkSMSData(input []string) error {
	return s.check.CheckDataSMS(input)
}
