package systemsProject

import (
	"io/ioutil"
	"server/internal/app/checkdata"
	"server/internal/app/models"
	"strings"
)

type Email interface {
	ReadEmailData() ([]models.EmailData, error)
}

type EmailSystem struct {
	check    *checkdata.CheckData
	fileName map[string]string
}

func (e *EmailSystem) ReadEmailData() ([]models.EmailData, error) {
	//init slice SMSData
	emailSlice := []models.EmailData{}

	data, err := ioutil.ReadFile(e.fileName["email.data"])
	if err != nil {
		return nil, err
	}
	//TODO:need another way to '\n'...
	for _, v := range strings.Split(string(data), "\n") {
		dataEmail := strings.Split(v, ";")
		emailData, err := e.CheckEmailData(dataEmail)
		if err != nil {
			continue
		}
		emailSlice = append(emailSlice, *emailData)
	}
	return emailSlice, nil
}

func (e *EmailSystem) CheckEmailData(input []string) (*models.EmailData, error) {
	return e.check.CheckEmailData(input)
}
