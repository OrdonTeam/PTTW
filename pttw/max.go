package pttw

import "math"

func MaxAbs(slice []float64) float64 {
	max := slice[0]
	for i := 0; i < len(slice); i++ {
		if (max < math.Abs(slice[i])) {
			max = math.Abs(slice[i])
		}
	}
	return max
}
