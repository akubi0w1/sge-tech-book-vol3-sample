package terror

type Code string

const (
	CodeOK Code = "ok"

	CodeInvalidArgument Code = "invalid_argument"
	CodeNotFound        Code = "not_found"

	CodeInternal Code = "internal"
	CodeUnknown  Code = "unknown"
)
