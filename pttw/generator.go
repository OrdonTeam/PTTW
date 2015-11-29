package pttw

func Generator(nextRandom func() float64) (func(int) []float64) {
	return func(size int) []float64 {
		slice := []float64{}
		for i := 0; i < size; i++ {
			slice = append(slice, nextRandom())
		}
		return slice
	}
}