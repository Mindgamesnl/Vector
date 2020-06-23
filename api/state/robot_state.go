package state

import "github.com/Mindgamesnl/Vector/api/enums"

type RobotState struct {
	Status byte // current state of the robot
}

func CreateIdleRobotStatus() RobotState {
	return RobotState{enums.RobotStatus.IsOnCharger}
}

func (state RobotState) AreMotorsMoving() bool {
	return state.Status == enums.RobotStatus.IsMoving
}