package unk

func Map(val, initialStart, initialMax, desiredStart, desiredMax float64) float64 {
	initialRange := initialMax - initialStart
	desiredRange := desiredMax - desiredStart
	mappedValue := desiredStart + ((val-initialStart)/initialRange)*desiredRange

	return mappedValue
}
