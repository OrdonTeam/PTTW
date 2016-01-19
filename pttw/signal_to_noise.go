package pttw
import "math"

func AddSignalToNoise(signal []float64, noise []float64, snr float64) []float64 {
	multiplier := CalculateMultiplier(signal, noise, snr)
	sum := []float64{}
	for i := 0; i < len(signal); i++ {
		sum = append(sum, signal[i] + noise[i] * multiplier)
	}
	return sum
}

func CalculateMultiplier(signal []float64, noise []float64, snr float64) float64 {
	requiredDifferenceInPowers := SnrToLinear(snr)
	powerOfSignal := PowerOfSignal(signal)
	requiredNoisePower := powerOfSignal / requiredDifferenceInPowers
	powerOfNoise := PowerOfSignal(noise)
	return math.Sqrt(requiredNoisePower / powerOfNoise)
}