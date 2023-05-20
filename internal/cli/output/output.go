package output

import (
	"io"

	ojson "github.com/akubi0w1/sge-tech-book-vol3-sample/internal/cli/output/json"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/terror"
)

type Output interface {
	// Log
	Log(msg string)
	// Logf
	Logf(format string, args ...interface{})
	// Format
	Format(data interface{})
}

// NewOutputWithFormatter
func NewOutputWithFormatter(formatter string, writer io.Writer) (Output, error) {
	switch formatter {
	case "json":
		return ojson.NewOutput(writer), nil

	default:
		return nil, terror.Newf(terror.CodeInvalidArgument, "invalid output formatter. formatter must be one of json or console.")
	}
}
