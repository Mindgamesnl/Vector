package enums

type RobotStatusValues struct {
	IsMoving           byte
	IsCarryingBlock    byte
	IsPickingOrPlacing byte
	IsPickedUp         byte
	IsButtonPressed    byte
	IsFalling          byte
	IsAnimating        byte
	IsPathing          byte
	LiftInPos          byte
	HeadInPos          byte
	CalmPowerMode      byte
	IsOnCharger        byte
	IsCharging         byte
	CliffDetected      byte
	AreWheelsMoving    byte
	IsBeingHeld        byte
	IsMotionDetected   byte
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
