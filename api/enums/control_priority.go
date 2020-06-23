package enums

type ControlPriorityValues struct {
	OverrideBehaviors uint8
	Default           uint8
	ReserveControl    uint8
}

var ControlPriority = ControlPriorityValues{
	OverrideBehaviors: 10,
	Default: 20,
	ReserveControl: 30,
}
