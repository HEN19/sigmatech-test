package endpoint

import (
	"fmt"
	"strconv"

	"github.com/api-skeleton/constanta"
	"github.com/api-skeleton/constanta/ErrorModel"
	"github.com/api-skeleton/dto/out"
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

		_, errInsert := CustomerService.InsertCustomerService(&userRequest)
		if errInsert.Code != constanta.CodeSuccessResponse {
			c.JSON(constanta.CodeInternalServerErrorResponse, ErrorModel.ErrorInternalServerError(c, errInsert.Details))
			return
		}
		out.ResponseOut(c, nil, true, constanta.CodeSuccessResponse, constanta.SuccessAddData)

		break
	case "GET":
		pageStr := c.DefaultQuery("page", "1")
		limitStr := c.DefaultQuery("limit", "10")

		page, _ := strconv.Atoi(pageStr)
		limit, _ := strconv.Atoi(limitStr)
		customers, errGet := CustomerService.GetListCustomerService(page, limit)
		if errGet.Code != constanta.CodeSuccessResponse {
			c.JSON(constanta.CodeInternalServerErrorResponse, ErrorModel.ErrorInternalServerError(c, errGet.Details))
			return
		}
		out.ResponseOut(c, customers, true, constanta.CodeSuccessResponse, constanta.SuccessGetData)
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

		customer, errGet := CustomerService.GetDetailCustomer(id)
		if errGet.Code != constanta.CodeSuccessResponse {
			c.JSON(constanta.CodeInternalServerErrorResponse, ErrorModel.ErrorInternalServerError(c, errGet.Details))
			return
		}
		out.ResponseOut(c, customer, true, constanta.CodeSuccessResponse, constanta.SuccessGetData)
		break
	case "POST":
		break
	case "DELETE":
		break
	}
}
