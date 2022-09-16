package systemsProject

import (
	"os"
	"server/internal/app/checkdata"
	"server/internal/app/models"
	"strings"
)

type VoiceCallSystem struct {
	check    *checkdata.CheckData
	config   *Config
	fileName map[string]string
}

func NewVoiceSystem(fileName map[string]string, config *Config) *VoiceCallSystem {
	return &VoiceCallSystem{
		check:    &checkdata.CheckData{},
		config:   config,
		fileName: fileName,
	}
}

func (vc *VoiceCallSystem) readVoice() ([]models.VoiceCallData, error) {

	voiceSlice := &[]models.VoiceCallData{}

	data, err := os.ReadFile(vc.fileName["voice.data"])
	if err != nil {
		return nil, err
	}

	for _, v := range strings.Split(string(data), "\n") {
		dataVoice := strings.Split(v, ";")
		voiceData, err := vc.check.CheckVoiceCall(dataVoice, vc.config.LenVoiceCallData)
		if err != nil {
			continue
		}
		*voiceSlice = append(*voiceSlice, *voiceData)
	}
	return *voiceSlice, nil
}

func (vc *VoiceCallSystem) GetVoiceData() ([]models.VoiceCallData, error) {
	type Result struct {
		Payload []models.VoiceCallData
		Error   error
	}
	in := make(chan Result)
	defer close(in)
	go func() {
		dataVoice, err := vc.readVoice()
		if err != nil {
			in <- Result{
				Payload: nil,
				Error:   err,
			}
		}
		in <- Result{
			Payload: dataVoice,
			Error:   nil,
		}
	}()
	result := <-in
	return result.Payload, result.Error
}
