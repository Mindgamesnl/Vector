package objects

type ProximityDataSensor struct {
	DistanceMm     int32
	IsLiftInFov    bool
	FoundObject    bool
	IsUnobstructed bool
	SignalQuality  float32
}

func CreateZeroProximitySensorData() ProximityDataSensor {
	return ProximityDataSensor{
		-1, false, false, false, 0,
	}
}