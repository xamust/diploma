package systems

import (
	"fmt"
	"server/internal/app/collect"
	"server/internal/app/models"
	"sort"
)

type Systems struct {
	sms   *SMSService
	check *collect.CheckData
}

func (s *Systems) GetSMSData() [][]models.SMSData {
	//sms
	//init sms service
	s.sms = &SMSService{check: s.check}
	dataSMS, err := s.sms.ReadSMS()
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
