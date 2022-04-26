package systemsProject

import (
	"github.com/sirupsen/logrus"
	"server/internal/app/models"
	"sort"
)

type SystemsProject struct {
	Logger           *logrus.Logger
	Config           *Config
	ParsingDataFiles map[string]string
}

//sms system..
func (s *SystemsProject) GetSMSData() ([][]models.SMSData, error) {
	//sms
	//init sms service
	sms := &SMSSystem{
		logger:   s.Logger,
		check:    &CheckData{Config: s.Config},
		fileName: s.ParsingDataFiles,
	}
	dataSMS, err := sms.ReadSMS()
	if err != nil {
		s.Logger.Errorf(err.Error())
		return nil, err
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
	return [][]models.SMSData{dataSMS, dataSMSDouble}, nil
}

//mms system...

//voice system...
func (s *SystemsProject) getVoiceData() ([]models.VoiceCallData, error) {

	//init voice system...
	voice := &VoiceCallSystem{
		logger:   s.Logger,
		check:    &CheckData{Config: s.Config},
		fileName: s.ParsingDataFiles,
	}

	dataVoice, err := voice.ReadVoiceData()
	if err != nil {
		s.Logger.Errorf(err.Error())
		return nil, err
	}
	return dataVoice, nil
}

//email system...
func (s *SystemsProject) getEmailData() (map[string][][]models.EmailData, error) {
	//init email system...
	email := &EmailSystem{
		logger:   s.Logger,
		check:    &CheckData{Config: s.Config},
		fileName: s.ParsingDataFiles,
	}
	emailData, err := email.ReadEmailData()
	if err != nil {
		s.Logger.Errorf(err.Error())
		return nil, err
	}
	//temp hashmap...
	tempEmailMap := make(map[string][]models.EmailData)

	//resultMap...
	resultMap := make(map[string][][]models.EmailData)

	//map create and fill...
	for _, data := range *emailData {
		tempEmailMap[data.Country] = append(tempEmailMap[data.Country], data)
		//sort temp hashmap by the way...
		for i := 0; i < len(tempEmailMap[data.Country])-1; i++ {
			for j := 0; j < len(tempEmailMap[data.Country])-i-1; j++ {
				if tempEmailMap[data.Country][j+1].DeliveryTime < tempEmailMap[data.Country][j].DeliveryTime {
					tempEmailMap[data.Country][j+1], tempEmailMap[data.Country][j] = tempEmailMap[data.Country][j], tempEmailMap[data.Country][j+1]
				}
			}
		}
	}

	for s2, _ := range tempEmailMap {
		resultMap[s2] = append(resultMap[s2], tempEmailMap[s2][0:3], tempEmailMap[s2][len(tempEmailMap)-5:len(tempEmailMap)-2])
	}
	return resultMap, nil
}

//billing system...
func (s *SystemsProject) getBillingData() (*models.BillingData, error) {
	//init billing system...
	billing := &BillingSystem{
		logger:   s.Logger,
		check:    &CheckData{Config: s.Config},
		fileName: s.ParsingDataFiles,
	}

	billingData, err := billing.ReadBillingData()
	if err != nil {
		s.Logger.Errorf(err.Error())
		return nil, err
	}

	return billingData, nil
}

//get result data...
func (s *SystemsProject) GetResultData() (*models.ResultSetT, error) {
	/*
		type item struct {
			dataSMS       [][]models.SMSData
			dastaMMS      [][]models.MMSData
			dataVoiceCall []models.VoiceCallData
			dataEmail     map[string][][]models.EmailData
			dataBilling   models.BillingData
			dataSupport   []int
			dataIncidents []models.IncidentData
			err           error
		}
		dataS := make(chan item)

		go func() {
			var sms item
			sms.dataSMS, sms.err = s.GetSMSData()
			dataS <- sms
		}()
		sms := <-dataS
		close(dataS)
		if sms.err != nil {
			s.Logger.Error(sms.err)
			return nil, sms.err
		}
	*/
	sms, err := s.GetSMSData()
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	voice, err := s.getVoiceData()
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	email, err := s.getEmailData()
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	billinig, err := s.getBillingData()
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	return &models.ResultSetT{
		SMS:       sms,
		MMS:       nil,
		VoiceCall: voice,
		Email:     email,
		Billing:   *billinig,
		Support:   nil,
		Incidents: nil,
	}, nil
}
