package collect

import (
	"github.com/sirupsen/logrus"
	"server/internal/app/models"
)

type Collect struct {
	Logger *logrus.Logger
	Config *Config
	//systems *systems.Systems
}

func (c *Collect) Start() error {
	c.Logger.Print("Start collecting...")

	//search *.data files...
	result, err := c.searchDataFiles()
	if err != nil {
		return err
	}
	c.Logger.Print("Поиск *.data файлов выполнен успешно!")

	c.Logger.Print(result)

	//init sms...

	return nil
}

func (c *Collect) searchDataFiles() (map[string]string, error) {
	parse := ParsingFolder{c.Config, c.Logger, make(map[string]string)}
	_, err := parse.FindFiles()
	if err != nil {
		return nil, err
	}
	return parse.mapFile, nil
}

func (c *Collect) GetResultData() *models.ResultSetT {
	return &models.ResultSetT{
		//SMS:       c.systems.GetSMSData(),
		SMS:       nil,
		MMS:       nil,
		VoiceCall: nil,
		Email:     nil,
		Billing:   models.BillingData{},
		Support:   nil,
		Incidents: nil,
	}
}
