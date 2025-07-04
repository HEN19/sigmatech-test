package ErrorModel

import (
	"github.com/api-skeleton/constanta" // Assuming constanta package for status codes and messages
	"github.com/gin-gonic/gin"
)

// ErrorDataNotFound returns a dynamic error response for data not found
func ErrorDataNotFound(dataIdentifier string) DynamicErrorResponse {
	// Here, we pass the data identifier or other context to the error message
	return CreateErrorResponse(
		constanta.CodeBadRequestResponse,                 // HTTP Status Code for "Not Found"
		constanta.ErrorDataUnknown,                       // Custom error message from constants
		"Data not found for identifier: "+dataIdentifier, // Details about the specific error
		"Data Error", // Error type
		map[string]string{
			"data_identifier": dataIdentifier, // Adding context like the data identifier that caused the error
			"method":          "GET",          // You can also add method, or any other context information
		},
	)
}

// ErrorInvalidRequest returns a dynamic error response for invalid requests
func ErrorInvalidRequest(method *gin.Context, field string, reason string) DynamicErrorResponse {
	return CreateErrorResponse(
		constanta.CodeBadRequestResponse,               // HTTP Status Code for "Bad Request"
		constanta.ErrorDataUnknown,                     // Custom error message from constants
		"Invalid request: "+reason+" for field "+field, // Provide more specific error details
		"Client Error", // Error type
		map[string]string{
			"field":  field,                 // Adding the field name causing the error
			"reason": reason,                // Adding the reason for the error
			"method": method.Request.Method, // Assuming a POST method for this example
		},
	)
}

// ErrorInternalServerError returns a dynamic error response for internal server issues
func ErrorInternalServerError(method *gin.Context, details string) DynamicErrorResponse {
	return CreateErrorResponse(
		constanta.CodeBadRequestResponse,  // HTTP Status Code for "Internal Server Error"
		constanta.ErrorInternalDB,         // Custom error message from constants
		"Internal server error: "+details, // Error details
		"Server Error",                    // Error type
		map[string]string{
			"details": details,               // Additional details about the internal error
			"method":  method.Request.Method, // Assuming method used for the request
		},
	)

}

func NonErrorResponse() DynamicErrorResponse {
	return CreateErrorResponse(
		constanta.CodeSuccessResponse, // HTTP Status Code for "Internal Server Error"
		"",                            // Custom error message from constants
		"",                            // Error details
		"",                            // Error type
		nil,
	)
}
