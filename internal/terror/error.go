package terror

import (
	"errors"
	"fmt"
)

type techBookError struct {
	code    Code
	message string
	cause   error
}

func (e *techBookError) Error() string {
	return fmt.Sprintf("Code: %s, message: %s, cause: %v", e.code, e.message, e.cause)
}

func Newf(code Code, format string, args ...interface{}) error {
	return &techBookError{
		code:    code,
		message: fmt.Sprintf(format, args...),
		cause:   nil,
	}
}

func Wrapf(code Code, err error, format string, args ...interface{}) error {
	return &techBookError{
		code:    code,
		message: fmt.Sprintf(format, args...),
		cause:   err,
	}
}

func GetCode(err error) Code {
	if err == nil {
		return CodeOK
	}

	var terr *techBookError
	if errors.As(err, &terr) {
		return terr.code
	}

	return CodeUnknown
}
