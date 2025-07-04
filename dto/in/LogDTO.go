package in

import (
	"fmt"

	"github.com/api-skeleton/constanta/ErrorModel"
	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	Id        int64  `json:"-"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}

func (input *UserRequest) ValidationRegistration(c *gin.Context) ErrorModel.DynamicErrorResponse {

	// ... validation logic here
	if input.FirstName == "" {
		return ErrorModel.ErrorInvalidRequest(c, "first_name", "first_name is required")
	}

	if input.LastName == "" {
		return ErrorModel.ErrorInvalidRequest(c, "last_name", "last_name is required")
	}

	if input.Gender != "L" && input.Gender != "P" {
		return ErrorModel.ErrorInvalidRequest(c, "gender", "gender must be L (Laki-Laki) or P (Perempuan)")

	}

	if input.Email != fmt.Sprintf("%s@%s", input.Username, "gmail.com") {
		return ErrorModel.ErrorInvalidRequest(c, "email", "email must be valid")
	}

	return input.mandatoryValidation()
}

func (input *UserRequest) mandatoryValidation() ErrorModel.DynamicErrorResponse {
	if input.Username == "" {
		return ErrorModel.ErrorInvalidRequest(nil, "username", "username is required")
	}
	if input.Password == "" {
		ErrorModel.ErrorInvalidRequest(nil, "username", "username is required")
	}

	return ErrorModel.NonErrorResponse()
}
