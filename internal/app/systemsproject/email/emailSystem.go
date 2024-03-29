package email

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"server/internal/app/checkdata"
	"server/internal/app/models"
)

type EmailSystem struct {
	check    *checkdata.CheckData
	config   *Config
	fileName map[string]string
}

func NewEmailSystem(fileName map[string]string, config *Config) *EmailSystem {
	return &EmailSystem{
		check:    &checkdata.CheckData{},
		config:   config,
		fileName: fileName,
	}
}

func (e *EmailSystem) readEmail() ([]models.EmailData, error) {

	var emailSlice []models.EmailData
	data, err := os.ReadFile(e.fileName[dEmail])
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(bytes.NewReader(data))
	r.Comma = ';'
	for {
		dataEmail, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		emailData, err := e.check.CheckEmailData(dataEmail, e.config.LenEmailData)
		if err != nil {
			continue
		}
		emailSlice = append(emailSlice, *emailData)
	}

	return emailSlice, nil
}

func (e *EmailSystem) GetEmailData() (map[string][][]models.EmailData, error) {
	type Result struct {
		Payload map[string][][]models.EmailData
		Error   error
	}

	inEmail := make(chan Result)
	go func() {
		emailData, err := e.readEmail()
		if err != nil {
			inEmail <- Result{
				Payload: nil,
				Error:   err,
			}
		}
		tempEmailMap := make(map[string][]models.EmailData)
		resultMap := make(map[string][][]models.EmailData)
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

		for s2 := range tempEmailMap {
			resultMap[s2] = append(resultMap[s2], tempEmailMap[s2][:3], tempEmailMap[s2][len(tempEmailMap[s2])-3:])
		}
		inEmail <- Result{
			Payload: resultMap,
			Error:   err,
		}
		close(inEmail)
	}()
	result := <-inEmail
	return result.Payload, result.Error
}
