package systemsProject

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"server/internal/app/models"
	"strings"
)

type EmailSystem struct {
	check    *CheckData
	logger   *logrus.Logger
	fileName map[string]string
}

func (e *EmailSystem) ReadEmailData() ([]models.EmailData, error) {
	//init slice SMSData
	emailSlice := []models.EmailData{}

	data, err := ioutil.ReadFile(e.fileName["email.data"])
	if err != nil {
		e.logger.Error(err)
		return nil, err
	}
	//TODO:need another way to '\n'...
	for _, v := range strings.Split(string(data), "\n") {
		dataEmail := strings.Split(v, ";")
		emailData, err := e.CheckEmailData(dataEmail)
		if err != nil {
			e.logger.Printf("data %v, corrupt!!! %s", dataEmail, err.Error())
			continue
		}
		//log.Printf("data %v, correct!!!!", dataSMS)
		emailSlice = append(emailSlice, *emailData)
	}
	return emailSlice, nil
}

func (e *EmailSystem) CheckEmailData(input []string) (*models.EmailData, error) {
	return e.check.CheckEmailData(input)
}
