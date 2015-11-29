package pttw

func Normalize(slice []float64) []float64 {
	max := MaxAbs(slice)
	normalized := []float64{}
	for i := 0; i < len(slice); i++ {
		normalized = append(normalized, slice[i] / max)
	}
	return normalized
}