package pttw
import (
	"fmt"
	"io"
)

func Read(bytes io.Reader) ([]float64, error) {
	slice := []float64{}
	for {
		var value float64
		n, err := fmt.Fscan(bytes, &value)
		if n == 0 || err != nil {
			return slice, err
		}
		slice = append(slice, value)
	}
	return slice, nil
}