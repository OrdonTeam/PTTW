package pttw

func AddSignalToNoise(signal []float64, noise []float64, snr float64) []float64 {
	sum := []float64{}
	multiplier := SnrToLinear(snr)
	for i := 0; i < len(signal); i++ {
		sum = append(sum, signal[i] + noise[i] / multiplier)
	}
	return sum
}