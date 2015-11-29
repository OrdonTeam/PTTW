package pttw

import (
	"io"
	"strconv"
)

func Write(slice []float64, writer io.Writer) error {
	for i := 0; i < len(slice) - 1; i++ {
		err := writeFloat(slice[i], writer)
		if err != nil {
			return err
		}
		_, err = writer.Write([]byte(" "))
		if err != nil {
			return err
		}
	}
	return writeFloat(slice[len(slice) - 1], writer)
}

func writeFloat(value float64, writer io.Writer) error {
	bytes := []byte(strconv.FormatFloat(value, 'f', -1, 64))
	_, err := writer.Write(bytes)
	return err
}