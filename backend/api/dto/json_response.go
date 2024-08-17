package dto

type JSONResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error,omitempty"`
	Query   any   `json:"query,omitempty"`
	Data    any   `json:"data,omitempty"`
}

func NewSuccessResponse(data any) *JSONResponse {
	return &JSONResponse{
		Success: true,
		Data:    data,
	}
}

func NewErrorResponse(err error) *JSONResponse {
	return &JSONResponse{
		Success: false,
		Error:   err,
	}
}
