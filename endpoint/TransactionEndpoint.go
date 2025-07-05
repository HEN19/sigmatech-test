package endpoint

import "github.com/gin-gonic/gin"

func TransactionWithoutParamEndpoint(c *gin.Context) {
	switch c.Request.Method {
	case "POST":
		// UserService.UserRegistration(response, request)
		break
	case "GET":
		break
	}
}

func TransactionWithParamEndpoint(c *gin.Context) {
	switch c.Request.Method {
	case "PUT":
		// UserService.UserProfileUpdate(response, request)
		break
	case "GET":
		break
	case "POST":
		break
	case "DELETE":
		break
	}
}
