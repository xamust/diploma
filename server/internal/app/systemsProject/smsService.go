package systemsProject

import (
	"io/ioutil"
	"log"
	"server/internal/app/models"
	"strings"
)

type SMSService struct {
	check    *CheckData
	fileName map[string]string
}

func (s *SMSService) ReadSMS() ([]models.SMSData, error) {

	//init slice SMSData
	SMSSlice := &[]models.SMSData{}

	data, err := ioutil.ReadFile(s.fileName["sms.data"])
	if err != nil {
		log.Print(err)
		return nil, err
	}

	//TODO:need another way to '\n'...
	for _, v := range strings.Split(string(data), "\n") {
		dataSMS := strings.Split(v, ";")
		if err := s.checkSMSData(dataSMS); err != nil {
			log.Printf("data %v, corrupt!!!\n%s", dataSMS, err.Error())
			continue
		}
		//log.Printf("data %v, correct!!!!", dataSMS)
		*SMSSlice = append(*SMSSlice, models.SMSData{
			Country:      dataSMS[0],
			Bandwidth:    dataSMS[1],
			ResponseTime: dataSMS[2],
			Provider:     dataSMS[3],
		})
	}
	return *SMSSlice, nil
}

func (s *SMSService) checkSMSData(input []string) error {
	return s.check.CheckDataSMS(input)
}
