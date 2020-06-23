package objects

type ProximityDataSensor struct {
	DistanceMm     uint32
	IsLiftInFov    bool
	FoundObject    bool
	IsUnobstructed bool
	SignalQuality  float32
}
