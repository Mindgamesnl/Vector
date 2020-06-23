package math

type Pose struct {
	PositionalLocation Position
	Rotation           Quaternion
	OriginId           uint64
}
