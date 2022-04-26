package systemsProject

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"server/internal/app/models"
	"sort"
)

type SystemsProject struct {
	Logger           *logrus.Logger
	Config           *Config
	ParsingDataFiles map[string]string
}

func (s *SystemsProject) GetSMSData() [][]models.SMSData {
	//sms
	//init sms service
	sms := &SMSService{check: &CheckData{Config: s.Config}, fileName: s.ParsingDataFiles}
	dataSMS, err := sms.ReadSMS()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	models.FullCountryNameSMS(dataSMS)

	// костыль с данными,ссылочный тип с указателями %)
	dataSMSDouble := make([]models.SMSData, len(dataSMS))
	copy(dataSMSDouble, dataSMS)
	sort.Slice(dataSMS, func(i, j int) bool {
		return dataSMS[i].Provider < dataSMS[j].Provider
	})
	sort.Slice(dataSMSDouble, func(i, j int) bool {
		return dataSMSDouble[i].Country < dataSMSDouble[j].Country
	})
	return [][]models.SMSData{dataSMS, dataSMSDouble}
}

func (s *SystemsProject) GetResultData() *models.ResultSetT {
	return &models.ResultSetT{
		SMS: s.GetSMSData(),
		//SMS:       nil,
		MMS:       nil,
		VoiceCall: nil,
		Email:     nil,
		Billing:   models.BillingData{},
		Support:   nil,
		Incidents: nil,
	}
}
