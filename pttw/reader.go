package pttw
import (
	"fmt"
	"io"
)

func Read(reader io.Reader) ([]float64, error) {
	slice := []float64{}
	for {
		var value float64
		n, err := fmt.Fscan(reader, &value)
		if n == 0 || err != nil {
			if (err.Error() == "EOF") {
				break
			}else {
				return slice, err
			}
		}
		slice = append(slice, value)
	}
	return slice, nil
}