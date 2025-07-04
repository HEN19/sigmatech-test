package ErrorModel

import (
	"time"
)

// DynamicErrorResponse represents the dynamic error model
type DynamicErrorResponse struct {
	Code        int               `json:"code"`
	Message     string            `json:"message"`
	Details     string            `json:"details,omitempty"`
	ErrorType   string            `json:"error_type,omitempty"`
	Timestamp   string            `json:"timestamp"`
	Additional  map[string]string `json:"additional,omitempty"`
}

// CreateErrorResponse creates a structured error response with dynamic fields
func CreateErrorResponse(code int, message string, details string, errorType string, additional map[string]string) DynamicErrorResponse {
	return DynamicErrorResponse{
		Code:      code,
		Message:   message,
		Details:   details,
		ErrorType: errorType,
		Timestamp: time.Now().Format(time.RFC3339),
		Additional: additional,
	}
}
