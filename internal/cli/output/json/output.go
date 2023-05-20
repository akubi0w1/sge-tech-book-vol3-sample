package json

import (
	"encoding/json"
	"fmt"
	"io"
)

// Output
type Output struct {
	io.Writer
}

// NewOutput
func NewOutput(out io.Writer) *Output {
	return &Output{
		Writer: out,
	}
}

func (o *Output) Log(msg string) {
	o.outputLog(msg)
}

func (o *Output) Logf(format string, args ...interface{}) {
	o.outputLog(format, args...)
}

func (o *Output) outputLog(format string, args ...interface{}) {
	format += "\n"
	if _, err := o.Write([]byte(fmt.Sprintf(format, args...))); err != nil {
		o.outputError(err)
	}
}

func (o *Output) Format(data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		o.outputError(err)
		return
	}

	if _, err = o.Write(jsonData); err != nil {
		o.outputError(err)
		return
	}
}

func (o *Output) outputError(err error) {
	fmt.Printf("Failed to output: %v\n", err)
}
