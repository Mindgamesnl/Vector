package control

import "craftmend.com/Mindgamesnl/Vector/api"

type ControlModule struct {

}

func CreateControlModule(api.Robot) ControlModule {
	return ControlModule{}
}