package apperror

import "fmt"

type AppError struct {
	Err       error
	File      string
	Method    string
	Submethod string
}

func (e *AppError) Error() string {
	errorString := fmt.Sprintf("%T error on %s: %s", e.Err, e.File, e.Method)
	if e.Submethod != "" {
		errorString += fmt.Sprintf(" - %s", e.Submethod)
	}
	errorString += fmt.Sprintf(": %s", e.Err.Error())

	return errorString
}

func NewError(e error, fileName string, method string, submethod string) *AppError {
	return &AppError{
		Err:       e,
		File:      fileName,
		Method:    method,
		Submethod: submethod,
	}
}
