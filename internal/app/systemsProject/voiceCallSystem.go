package systemsProject

import (
	"os"
	"server/internal/app/checkdata"
	"server/internal/app/models"
	"strings"
)

type Voice interface {
	ReadVoiceData() ([]models.VoiceCallData, error)
}

type VoiceCallSystem struct {
	check    *checkdata.CheckData
	config   *Config
	fileName map[string]string
}

func (vc *VoiceCallSystem) ReadVoiceData() ([]models.VoiceCallData, error) {

	//init slice voiceData
	voiceSlice := &[]models.VoiceCallData{}

	//todo: map ????
	data, err := os.ReadFile(vc.fileName["voice.data"])
	if err != nil {
		return nil, err
	}

	//TODO:need another way to '\n'...
	for _, v := range strings.Split(string(data), "\n") {
		dataVoice := strings.Split(v, ";")
		voiceData, err := vc.CheckVoiceData(dataVoice)
		if err != nil {
			continue
		}
		*voiceSlice = append(*voiceSlice, *voiceData)
	}
	return *voiceSlice, nil
}

func (vc *VoiceCallSystem) CheckVoiceData(input []string) (*models.VoiceCallData, error) {
	return vc.check.CheckVoiceCall(input, vc.config.LenVoiceCallData)
}
