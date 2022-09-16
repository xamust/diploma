package voice

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"server/internal/app/checkdata"
	"server/internal/app/models"
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
	data, err := os.ReadFile(vc.fileName[dVoice])
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(bytes.NewReader(data))
	r.Comma = ';'
	for {
		dataVoice, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
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
	inVoice := make(chan Result)
	go func() {
		dataVoice, err := vc.readVoice()
		if err != nil {
			inVoice <- Result{
				Payload: nil,
				Error:   err,
			}
		}
		inVoice <- Result{
			Payload: dataVoice,
			Error:   nil,
		}
		close(inVoice)
	}()
	result := <-inVoice
	return result.Payload, result.Error
}
