package pttw

func AddSignalToNoise(signal []float64, noise []float64, snr int64) []float64 {
	sum := []float64{}
	for i := 0; i < len(signal); i++ {
		sum = append(sum, signal[i] + noise[i])
	}
	return sum
}