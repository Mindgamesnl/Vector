package audio

import "craftmend.com/Mindgamesnl/Vector/api"

type AudioModule struct {

}

func CreateAudioModule(api.Robot) AudioModule {
	return AudioModule{}
}