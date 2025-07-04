package UserService

import (
	"database/sql"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta"
	"github.com/api-skeleton/constanta/ErrorModel"
	"github.com/api-skeleton/utils"
	"github.com/gin-gonic/gin"

	"github.com/api-skeleton/dao"
	"github.com/api-skeleton/dto/in"
	"github.com/api-skeleton/dto/out"
	"github.com/api-skeleton/model"
)

func UserRegistration(c *gin.Context) {
	var reqBody model.UserModel
	// Get the user body from the request
	userRequest, err := utils.GetUserBody(c)
	if err != nil {
		c.JSON(constanta.CodeBadRequestResponse, ErrorModel.ErrorInternalServerError(c, err.Error()))
		return
	}

	// Perform validation before mapping
	errValidation := userRequest.ValidationRegistration(c)
	if errValidation.Code != constanta.CodeSuccessResponse {
		c.JSON(errValidation.Code, errValidation)
		return
	}

	// Map the request body to the model
	reqBody = mapToUserModel(userRequest)

	// Connect to DB (use connection pool for better performance)
	db := config.Connect()
	defer db.Close()

	// Insert user into DB
	err = dao.UserDAO.InsertUser(db, reqBody)
	if err != nil {
		c.JSON(constanta.CodeInternalServerErrorResponse, ErrorModel.ErrorInternalServerError(c, err.Error()))
		return
	}

	// Respond with success
	out.ResponseOut(c, nil, true, constanta.CodeSuccessResponse, constanta.SuccessRegistrationData)
}

// mapToUserModel maps the UserRequest to UserModel, performing validation if needed.
func mapToUserModel(reqBody in.UserRequest) model.UserModel {
	// Map to model
	return model.UserModel{
		ID:        sql.NullInt64{Int64: reqBody.Id},
		Username:  sql.NullString{String: reqBody.Username},
		Password:  sql.NullString{String: reqBody.Password},
		FirstName: sql.NullString{String: reqBody.FirstName},
		LastName:  sql.NullString{String: reqBody.LastName},
		Gender:    sql.NullString{String: reqBody.Gender},
		Telephone: sql.NullString{String: reqBody.Telephone},
		Email:     sql.NullString{String: reqBody.Email},
		Address:   sql.NullString{String: reqBody.Address},
	}
}
