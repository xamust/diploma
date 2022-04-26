package systemsProject

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"server/internal/app/models"
	"strings"
)

type VoiceCallSystem struct {
	check    *CheckData
	logger   *logrus.Logger
	fileName map[string]string
}

func (vc *VoiceCallSystem) ReadVoiceData() ([]models.VoiceCallData, error) {

	//init slice voiceData
	voiceSlice := &[]models.VoiceCallData{}

	data, err := ioutil.ReadFile(vc.fileName["voice.data"])
	if err != nil {
		vc.logger.Error(err)
		return nil, err
	}

	//TODO:need another way to '\n'...
	for _, v := range strings.Split(string(data), "\n") {
		dataVoice := strings.Split(v, ";")
		voiceData, err := vc.CheckVoiceData(dataVoice)
		if err != nil {
			vc.logger.Printf("data %v, corrupt!!!\n%s", dataVoice, err.Error())
			continue
		}
		//log.Printf("data %v, correct!!!!", dataSMS)
		*voiceSlice = append(*voiceSlice, *voiceData)
	}
	return *voiceSlice, nil
}

func (vc *VoiceCallSystem) CheckVoiceData(input []string) (*models.VoiceCallData, error) {
	return vc.check.CheckVoiceCall(input)
}
