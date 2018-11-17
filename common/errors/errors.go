package errors

import (
	"fmt"
)

// ErrorCode type error code for all
// possible errors
type ErrorCode string

const (
	ErrCodeMongoDB             = ErrorCode("01")
	ErrCodeEmptyServiceName    = ErrorCode("02")
	ErrCodeInvalideServiceName = ErrorCode("03")
)

// ErrorMessages all error messages to return to user
var ErrorMessages = map[ErrorCode]*Error{
	ErrCodeMongoDB: {
		ErrMessage: "Unable to connect, err %+v",
	},
	ErrCodeEmptyServiceName: {
		ErrMessage: "Service name cannot be empty",
	},
	ErrCodeInvalideServiceName: {
		ErrMessage: "Unable to find service with name %s",
	},
}

type Error struct {
	ErrMessage string
}

func (self *Error) Error() string {
	return self.ErrMessage
}

func newError(text string, args ...interface{}) *Error {
	return &Error{ErrMessage: fmt.Sprintf(text, args...)}
}

// ErrorLog finds error message in
// ErrorMessage map by key and construct
// a message with newError()
func ErrorLog(key ErrorCode, args ...interface{}) *Error {
	errorMessage, ok := ErrorMessages[key]
	if !ok {
		return newError("Unable to find error with code %s", key)
	}
	return newError(errorMessage.Error(), args)
}
