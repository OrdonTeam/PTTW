package pttw

import (
	"io"
	"strconv"
)

func Write(slice []float64, writer io.Writer) error {
	for i := 0; i < len(slice); i++ {
		bytes := []byte(strconv.FormatFloat(slice[i], 'f', -1, 64))
		_, err := writer.Write(bytes)
		if err != nil {
			return err
		}
	}
	return nil
}