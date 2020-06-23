package audio

import "github.com/Mindgamesnl/Vector/api"

type AudioModule struct {

}

func CreateAudioModule(api.Robot) AudioModule {
	return AudioModule{}
}