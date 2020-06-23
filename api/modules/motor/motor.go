package motor

import "craftmend.com/Mindgamesnl/Vector/api"

type MotorModule struct {

}

func CreateMotorModule(api.Robot) MotorModule {
	return MotorModule{}
}