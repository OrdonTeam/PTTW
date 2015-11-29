package pttw

func Normalize(slice []float64) []float64 {
	max := Max(slice)
	normalized := []float64{}
	for i := 0; i < len(slice); i++ {
		normalized = append(normalized, slice[i] / max)
	}
	return normalized
}

func Max(slice []float64) float64 {
	max := slice[0]
	for i := 0; i < len(slice); i++ {
		if (max < slice[i]) {
			max = slice[i]
		}
	}
	return max
}