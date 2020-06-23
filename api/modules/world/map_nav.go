package world

import "craftmend.com/Mindgamesnl/Vector/api"

type MapNavModule struct {

}

func CreateMapNavModule(api.Robot) MapNavModule {
	return MapNavModule{}
}