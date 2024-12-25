package errorcode

var (
	ErrorCodeNotImplemented = newError(1000, "Not implemented")
	ErrorCodeInvalidParam   = newError(1001, "Invalid param")
)

type ErrorCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrorCode) Error() string {
	return e.Message
}

func (e ErrorCode) GetCode() int {
	return e.Code
}

func newError(code int, message string) *ErrorCode {
	return &ErrorCode{Code: code, Message: message}
}
