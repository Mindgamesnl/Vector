package world

import "github.com/Mindgamesnl/Vector/api"

type MapNavModule struct {

}

func CreateMapNavModule(api.Robot) MapNavModule {
	return MapNavModule{}
}