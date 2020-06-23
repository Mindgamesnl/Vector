package math

import "math"

type Quaternion struct {
	Q0 float64 // W
	Q1 float64 // I
	Q2 float64 // J
	Q3 float64 // K
}

func (quat Quaternion) CalculateAngleZ() float64 {
	return math.Atan2(2 * (quat.Q1 * quat.Q2 + quat.Q0 * quat.Q3), 1 - 2 * (math.Pow(quat.Q2, 2) + math.Pow(quat.Q3, 2)))
}

func QuaternionFromRadians(radiansZ float64) Quaternion {
	return Quaternion{
		Q0: math.Cos(radiansZ / 2),
		Q1: 0,
		Q2: 0,
		Q3: math.Sin(radiansZ / 2),
	}
}