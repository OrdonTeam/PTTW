package pttw

import (
	"io"
	"strconv"
	"bytes"
)

func Write(slice []float64, writer io.Writer) error {
	bytes := bytes.Buffer{}
	for i := 0; i < len(slice) - 1; i++ {
		bytes.WriteString(strconv.FormatFloat(slice[i], 'f', -1, 64))
		bytes.WriteString(" ")
	}
	bytes.WriteString(strconv.FormatFloat(slice[len(slice) - 1], 'f', -1, 64))
	_, err := writer.Write(bytes.Bytes())
	return err
}
