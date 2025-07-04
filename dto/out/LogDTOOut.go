package out

import (
	"github.com/api-skeleton/utils"
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	API APIMessage `json:"API"`
}

type APIMessage struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

// String converts the APIResponse to a JSON string.
func (ar APIResponse) String() string {
	return utils.StructToJSON(ar)
}

// ResponseOut sends a structured JSON response using Gin context.
func ResponseOut(c *gin.Context, data interface{}, success bool, code int, message string) {
	// Set the Content-Type header
	c.Header("Content-Type", "application/json")

	// Create the response structure
	apiResponse := APIResponse{
		API: APIMessage{
			Success: success,
			Code:    code,
			Message: message,
			Content: data,
		},
	}

	// Send the response with the appropriate HTTP status code
	c.JSON(code, apiResponse)
}
