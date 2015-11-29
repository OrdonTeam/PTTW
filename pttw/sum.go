package pttw

func Sum(first []float64, second []float64) []float64 {
	sum := []float64{}
	for i := 0; i < len(first); i++ {
		sum = append(sum, first[i] + second[i])
	}
	return sum
}