package systemsProject

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/internal/app/models"
)

type MMSService struct {
	check    *CheckData
	fileName map[string]string
}

func (m *MMSService) ReadMMS() ([]models.MMSData, error) {
	return m.GetMMSData()
}

func (m *MMSService) GetMMSData() ([]models.MMSData, error) {
	//todo:config file....
	//todo: http.client вынести...
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8383/mms", nil)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error! Response status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	log.Printf("Success upload MMS data, response status code %d", resp.StatusCode)

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	var mmsMod *[]models.MMSData
	if err := json.Unmarshal(data, &mmsMod); err != nil {
		log.Print(err.Error())
		return nil, err
	}

	//todo: new var???
	var dataMMS []models.MMSData
	for _, v := range *mmsMod {
		if err := m.CheckJSONMMS(&v); err != nil {
			log.Print(err)
			continue
		}
		dataMMS = append(dataMMS, v)
	}
	log.Print("MMS data uploading complete!")
	return dataMMS, nil
}

func (m *MMSService) CheckJSONMMS(v *models.MMSData) error {
	return m.check.CheckDataMMS(v)
}
