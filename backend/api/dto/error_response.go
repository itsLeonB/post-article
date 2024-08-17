package dto

import (
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"-"`
	Name    string `json:"name"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func (e *ErrorResponse) Error() string {
	errorString := fmt.Sprintf("%s: %s.", e.Name, e.Message)
	if e.Details != nil {
		errorString += fmt.Sprintf("Details: %s", e.Details)
	}
	return errorString
}

func InternalServerError() *ErrorResponse {
	return &ErrorResponse{
		Code:    http.StatusInternalServerError,
		Name:    "InternalServerError",
		Message: "server encounters unexpected error",
	}
}
