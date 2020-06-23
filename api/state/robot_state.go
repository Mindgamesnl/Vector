package state

import "craftmend.com/Mindgamesnl/Vector/api/enums"

type RobotState struct {
	Status byte // current state of the robot
}

func (state RobotState) AreMotorsMoving() bool {
	return state.Status == enums.RobotStatus.IsMoving
}