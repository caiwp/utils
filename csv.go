package utils

import (
	"encoding/csv"
	"io"

	"github.com/gocarina/gocsv"
)

func customCSVWrite(out io.Writer) *gocsv.SafeCSVWriter {
	writer := gocsv.NewSafeCSVWriter(csv.NewWriter(out))
	writer.Comma = '\t'

	return writer
}

func MarshalWithoutHeaders(in interface{}, out io.Writer) (err error) {
	gocsv.SetCSVWriter(customCSVWrite)
	return gocsv.MarshalWithoutHeaders(in, out)
}
