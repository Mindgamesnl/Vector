package math

type Pose struct {
	PositionalLocation Position
	Rotation           Quaternion
	OriginId           uint64
}

func CreateZeroPose() Pose {
	return Pose{
		Position{X: 0, Y: 0, Z: 0},
		Quaternion{Q0: 0, Q1: 0, Q2: 0, Q3: 0},
		0,
	}
}
