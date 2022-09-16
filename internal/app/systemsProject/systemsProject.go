package systemsProject

import (
	"math"
	"net/http"
	"server/internal/app/checkdata"
	"server/internal/app/models"
	"sort"
)

type SystemsProject struct {
	Config           *Config
	ParsingDataFiles map[string]string
}

// voice system...
func (s *SystemsProject) getVoiceData() ([]models.VoiceCallData, error) {

	//init voice system...
	voice := &VoiceCallSystem{
		check:    &checkdata.CheckData{},
		fileName: s.ParsingDataFiles,
		config:   s.Config,
	}

	dataVoice, err := voice.ReadVoiceData()
	if err != nil {
		return nil, err
	}
	return dataVoice, nil
}

// email system...
func (s *SystemsProject) getEmailData() (map[string][][]models.EmailData, error) {
	//init email system...
	email := &EmailSystem{
		check:    &checkdata.CheckData{},
		fileName: s.ParsingDataFiles,
		config:   s.Config,
	}
	emailData, err := email.ReadEmailData()
	if err != nil {
		return nil, err
	}

	//temp hashmap...
	tempEmailMap := make(map[string][]models.EmailData)

	//resultMap...
	resultMap := make(map[string][][]models.EmailData)

	//map create and fill...
	for _, data := range emailData {

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
		resultMap[s2] = append(resultMap[s2], tempEmailMap[s2][:3], tempEmailMap[s2][len(tempEmailMap[s2])-3:])
	}
	return resultMap, nil
}

// another parent struct models/ParentStruct.go.14
func (s *SystemsProject) getAnotherEmailData() ([][]models.EmailData, error) {

	anotherResultMass := make([][]models.EmailData, 0)
	anotherEmail, err := s.getEmailData()
	if err != nil {
		return nil, err
	}
	for _, v := range anotherEmail {
		anotherResultMass = append(anotherResultMass, v...)
	}
	return anotherResultMass, nil
}

// billing system...
func (s *SystemsProject) getBillingData() (*models.BillingData, error) {
	//init billing system...
	billing := &BillingSystem{
		check:    &checkdata.CheckData{},
		fileName: s.ParsingDataFiles,
	}

	billingData, err := billing.ReadBillingData()
	if err != nil {
		return nil, err
	}

	return billingData, nil
}

// support system...
func (s *SystemsProject) getSupportData() ([]int, error) {

	//init billing system...
	support := &SupportService{
		check:  &checkdata.CheckData{},
		client: &http.Client{},
		config: s.Config,
	}
	//
	supportData, err := support.GetSupportData()
	if err != nil {
		return nil, err
	}
	var countLoad, countTime, ticketCount float64
	calculatedTime := 60 / float64(s.Config.TickerPerHour)
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
	return []int{int(math.Round(countLoad / float64(len(*supportData)))), 60 / s.Config.TickerPerHour * int(ticketCount)}, nil
}

// incident system...
func (s *SystemsProject) getIncidentData() ([]models.IncidentData, error) {
	incident := &IncidentSystem{
		check:  &checkdata.CheckData{},
		client: &http.Client{},
		config: s.Config,
	}
	incidentData, err := incident.GetIncidentData()
	if err != nil {
		return nil, err
	}
	sort.Slice(incidentData, func(i, j int) bool {
		return incidentData[i].Status < incidentData[j].Status
	})

	return incidentData, nil
}

// get result data...
func (s *SystemsProject) GetResultData() (*models.ResultSetT, error) {

	sms := NewSMSSystem(s.ParsingDataFiles, s.Config)
	smsData, err := sms.GetSMSData()
	if err != nil {
		return nil, err
	}

	mms := NewMMSSystem(s.Config)
	mmsData, err := mms.GetMMSData()
	if err != nil {
		return nil, err
	}

	voice, err := s.getVoiceData()
	if err != nil {
		return nil, err
	}
	email, err := s.getEmailData()
	//another parent struct models/ParentStruct.go.14
	//email, err := s.getAnotherEmailData()
	if err != nil {
		return nil, err
	}
	billinig, err := s.getBillingData()
	if err != nil {
		return nil, err
	}

	support, err := s.getSupportData()
	if err != nil {
		return nil, err
	}
	incident, err := s.getIncidentData()
	if err != nil {
		return nil, err
	}

	return &models.ResultSetT{
		SMS:       smsData,
		MMS:       mmsData,
		VoiceCall: voice,
		Email:     email,
		Billing:   *billinig,
		Support:   support,
		Incidents: incident,
	}, nil
}
