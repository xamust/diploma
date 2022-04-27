package systemsProject

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"server/internal/app/models"
	"sort"
)

type SystemsProject struct {
	Logger           *logrus.Logger
	Config           *Config
	ParsingDataFiles map[string]string
}

//sms system..
func (s *SystemsProject) getSMSData() ([][]models.SMSData, error) {
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
func (s *SystemsProject) getMMSData() ([][]models.MMSData, error) {

	//init mms service
	mms := &MMSSystem{
		logger: s.Logger,
		check:  &CheckData{Config: s.Config},
		client: &http.Client{},
		config: s.Config,
	}

	dataMMS, err := mms.ReadMMS()
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	models.FullCountryNameMMS(dataMMS)
	// костыль с данными,ссылочный тип с указателями %)
	dataMMSDouble := make([]models.MMSData, len(dataMMS))
	copy(dataMMSDouble, dataMMS)
	//sort by provider
	sort.Slice(dataMMS, func(i, j int) bool {
		return dataMMS[i].Provider < dataMMS[j].Provider
	})
	//sort by country name
	sort.Slice(dataMMSDouble, func(i, j int) bool {
		return dataMMSDouble[i].Country < dataMMSDouble[j].Country
	})

	return [][]models.MMSData{dataMMS, dataMMSDouble}, nil
}

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

		resultMap[s2] = append(resultMap[s2], tempEmailMap[s2][:3], tempEmailMap[s2][len(tempEmailMap[s2])-5:])

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

//support system...

//incident system...
func (s *SystemsProject) getIncidentData() ([]models.IncidentData, error) {
	//incidents
	//init incident service
	incident := &IncidentSystem{
		logger: s.Logger,
		check:  &CheckData{Config: s.Config},
		client: &http.Client{},
		config: s.Config,
	}
	incidentData, err := incident.ReadIncident()
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	sort.Slice(incidentData, func(i, j int) bool {
		return incidentData[i].Status < incidentData[j].Status
	})

	return incidentData, nil
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
			sms.dataSMS, sms.err = s.getSMSData()
			dataS <- sms
		}()
		sms := <-dataS
		close(dataS)
		if sms.err != nil {
			s.Logger.Error(sms.err)
			return nil, sms.err
		}
	*/
	sms, err := s.getSMSData()
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	mms, err := s.getMMSData()
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	voice, err := s.getVoiceData()
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	//email, err := s.getEmailData()
	//if err != nil {
	//	s.Logger.Error(err)
	//	return nil, err
	//}
	billinig, err := s.getBillingData()
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}
	incident, err := s.getIncidentData()
	if err != nil {
		s.Logger.Error(err)
		return nil, err
	}

	//test := []models.EmailData{"AT": {{"AT", "Yahoo", 18}, {"AT", "GMX", 261}, {"AT", "Comcast", 315}, {"AT", "Protonmail", 350}, {"AT", "RediffMail", 461}, {"AT", "Orange", 579}}}
	test := [][]models.EmailData{{{"AT", "Yahoo", 18}, {"AT", "GMX", 261}, {"AT", "Comcast", 315}}, {{"AT", "Protonmail", 350}, {"AT", "RediffMail", 461}, {"AT", "Orange", 579}}}

	return &models.ResultSetT{
		SMS:       sms,
		MMS:       mms,
		VoiceCall: voice,
		Email:     test,
		Billing:   *billinig,
		Support:   []int{3, 63},
		Incidents: incident,
	}, nil
}
