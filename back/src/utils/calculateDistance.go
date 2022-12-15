package utils

import "math"

func CalculateDistance(x float64, y float64) float64 {
	var safeDistance float64 = 100000
	var xOri float64 = 250000
	var yOri float64 = 250000
	xDiff := x - xOri
	yDiff := y - yOri
	distance := math.Sqrt(math.Pow(xDiff, 2) + math.Pow(yDiff, 2))

	if safeDistance > distance {
		return distance
	}

	return -1
}
