package voice

import "server/internal/app/models"

const (
	dVoice = "voice.data"
)

type Voice interface {
	readVoice() ([]models.VoiceCallData, error)
	GetVoiceData() ([]models.VoiceCallData, error)
}
