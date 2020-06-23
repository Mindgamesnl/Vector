package enums

type RobotStatusValues struct {
	IsMoving           int64
	IsCarryingBlock    int64
	IsPickingOrPlacing int64
	IsPickedUp         int64
	IsButtonPressed    int64
	IsFalling          int64
	IsAnimating        int64
	IsPathing          int64
	LiftInPos          int64
	HeadInPos          int64
	CalmPowerMode      int64
	IsOnCharger        int64
	IsCharging         int64
	CliffDetected      int64
	AreWheelsMoving    int64
	IsBeingHeld        int64
	IsMotionDetected   int64
}

var RobotStatus = RobotStatusValues{
	IsMoving:           0x1,
	IsCarryingBlock:    0x2,
	IsPickingOrPlacing: 0x4,
	IsPickedUp:         0x8,
	IsButtonPressed:    0x10,
	IsFalling:          0x20,
	IsAnimating:        0x40,
	IsPathing:          0x80,
	LiftInPos:          0x100,
	HeadInPos:          0x200,
	CalmPowerMode:      0x400,
	IsOnCharger:        0x1000,
	IsCharging:         0x2000,
	CliffDetected:      0x4000,
	AreWheelsMoving:    0x8000,
	IsBeingHeld:        0x10000,
	IsMotionDetected:   0x20000,
}
