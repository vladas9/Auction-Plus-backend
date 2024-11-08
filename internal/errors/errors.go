package errors

import (
	"fmt"
	"net/http"
	"runtime"
)

type ApiError struct {
	ErrorMsg any   `json:"error"`
	Status   int   `json:"status"`
	Err      error `json:"-"`
}

func (a *ApiError) Error() string {
	str := fmt.Sprintf("%v", a.ErrorMsg)
	if a.Err == nil {
		return str
	}
	return fmt.Sprintf("%s: %v", str, a.Err)
}

type ErrMessage interface{}

type Validator interface {
	Validate() error
}

// validation of request fails, i.e. request fields
func NotValid(msg ErrMessage, err error) error {
	return &ApiError{
		Status:   http.StatusBadRequest,
		ErrorMsg: msg,
		Err:      formatErr(err),
	}
}

// requested resources cannot be found. e.g. auction requested doesn't exist
func NotFound(msg ErrMessage, err error) error {
	return &ApiError{
		Status:   http.StatusNotFound,
		ErrorMsg: msg,
		Err:      formatErr(err),
	}
}

// auth token missing or expired
func Unauthorized(msg ErrMessage, err error) error {
	return &ApiError{
		Status:   http.StatusUnauthorized,
		ErrorMsg: msg,
		Err:      formatErr(err),
	}
}

// user is logged in but does not have permission
func Forbidden(msg ErrMessage, err error) error {
	return &ApiError{
		Status:   http.StatusForbidden,
		ErrorMsg: msg,
		Err:      formatErr(err),
	}
}

// e.g. creating a user with existing email,
// placing bid lower than current one, etc.
func Conflict(msg ErrMessage, err error) error {
	return &ApiError{
		Status:   http.StatusConflict,
		ErrorMsg: msg,
		Err:      formatErr(err),
	}
}

// an error that should not happen, is server's fault
func Internal(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("Internal error: %w", formatErr(err))
}

// wraps error with current function name
func Next(err error) error {
	if err == nil {
		return nil
	}

	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("unknown function: %w", err)
	}
	fn := runtime.FuncForPC(pc)
	functionName := fn.Name()

	return fmt.Errorf("%s: %w", functionName, err)
}

// formats errors for logger
func formatErr(err error) error {
	if err == nil {
		return nil
	}

	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return fmt.Errorf("unknown function: %w", err)
	}
	fn := runtime.FuncForPC(pc)
	functionName := fn.Name()

	return fmt.Errorf("%s: %w\n\t--> %s:%d", functionName, err, file, line)
}
