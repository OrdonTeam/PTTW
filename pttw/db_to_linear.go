package pttw
import "math"

func SnrToLinear(db float64) float64 {
	return math.Pow(10, db / 20)
}