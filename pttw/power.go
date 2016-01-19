package pttw

func PowerOfSignal(signal []float64) float64 {
	powerSum := float64(0.0)
	for _, sample := range signal {
		powerSum += sample * sample
	}
	return powerSum / float64(len(signal))
}