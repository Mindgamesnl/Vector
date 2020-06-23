package screen

import "github.com/Mindgamesnl/Vector/api"

type ScreenModule struct {
	
}

func CreateScreenModule(api.Robot) ScreenModule {
	return ScreenModule{}
}