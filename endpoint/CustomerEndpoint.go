package endpoint

import (
	"fmt"
	"strconv"

	"github.com/api-skeleton/constanta"
	"github.com/api-skeleton/constanta/ErrorModel"
	"github.com/api-skeleton/service/CustomerService"
	"github.com/api-skeleton/utils"
	"github.com/gin-gonic/gin"
)

func CustomerWithoutParamEndpoint(c *gin.Context) {
	switch c.Request.Method {
	case "POST":
		// var reqBody model.CustomerModel
		// Get the user body from the request
		userRequest, err := utils.GetCustomerBody(c)
		if err != nil {
			c.JSON(constanta.CodeBadRequestResponse, ErrorModel.ErrorInternalServerError(c, err.Error()))
			return
		}

		// Perform validation before mapping
		errValidation := userRequest.ValidationCustomerInsert(c)
		if errValidation.Code != constanta.CodeSuccessResponse {
			c.JSON(errValidation.Code, errValidation)
			return
		}

		CustomerService.InsertCustomerService(&userRequest)

		break
	case "GET":
		pageStr := c.DefaultQuery("page", "1")
		limitStr := c.DefaultQuery("limit", "10")

		page, _ := strconv.Atoi(pageStr)
		limit, _ := strconv.Atoi(limitStr)
		CustomerService.GetListCustomerService(page, limit)
		break
	}
}

func CustomerWithParamEndpoint(c *gin.Context) {
	id := c.Param("id")
	switch c.Request.Method {
	case "PUT":
		fmt.Println(id)
		break
	case "GET":
		break
	case "POST":
		break
	case "DELETE":
		break
	}
}
