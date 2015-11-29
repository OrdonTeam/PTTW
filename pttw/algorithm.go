package pttw
import "io"

func Algorithm(reader io.Reader, writer io.Writer, snr float64, nextRandom func() float64) {
	signal, _ := Read(reader)
	noise := Generator(nextRandom)(len(signal))
	signalWithNoise := AddSignalToNoise(signal, noise, snr)
	Write(signalWithNoise, writer)
}