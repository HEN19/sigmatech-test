package UserService

import (
	"database/sql"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta"
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
		out.ResponseOut(c, nil, false, constanta.CodeBadRequestResponse, "Invalid request body")
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
		out.ResponseOut(c, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}

	// Respond with success
	out.ResponseOut(c, nil, true, constanta.CodeSuccessResponse, constanta.SuccessRegistrationData)
}

// getUserBody parses the request body into a UserRequest struct.
func getUserBody(c *gin.Context) (in.UserRequest, error) {
	var userRequest in.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		return userRequest, err
	}
	return userRequest, nil
}

// mapToUserModel maps the UserRequest to UserModel, performing validation if needed.
func mapToUserModel(reqBody in.UserRequest) model.UserModel {
	// Perform validation before mapping
	reqBody.ValidationRegistration()

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
