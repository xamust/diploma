package systemsProject

import (
	"os"
	"server/internal/app/checkdata"
	"server/internal/app/models"
	"strings"
)

type SMS interface {
	ReadSMS() ([]models.SMSData, error)
}

type SMSSystem struct {
	check    *checkdata.CheckData
	config   *Config
	fileName map[string]string
}

func (s *SMSSystem) ReadSMS() ([]models.SMSData, error) {

	//init slice SMSData
	SMSSlice := &[]models.SMSData{}

	data, err := os.ReadFile(s.fileName["sms.data"])
	if err != nil {
		return nil, err
	}

	//TODO:need another way to '\n'...
	for _, v := range strings.Split(string(data), "\n") {
		dataSMS := strings.Split(v, ";")
		if err = s.check.CheckDataSMS(dataSMS, s.config.LenSmsData); err != nil {
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
