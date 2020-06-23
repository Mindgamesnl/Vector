package control

import "github.com/Mindgamesnl/Vector/api"

type ControlModule struct {

}

func CreateControlModule(api.Robot) ControlModule {
	return ControlModule{}
}