package UserService

import (
	"fmt"
	"time"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta"
	"github.com/api-skeleton/model"
	"github.com/gin-gonic/gin"

	"github.com/api-skeleton/dao"
	"github.com/api-skeleton/dto/out"
	"github.com/api-skeleton/utils"
)

func LoginService(c *gin.Context) (err error) {
	var user model.UserModel

	now := time.Now()
	fmt.Println("HIT -> LoginService.go On ", now.Format("2006-01-02 15:04:05"))

	// Get user request body from Gin context
	userBody, err := utils.GetUserBody(c)

	// Map the request body to the repository model
	userRepo := mapToUserModel(userBody)

	// Connect to the database
	db := config.Connect()
	defer db.Close()

	// Validate input (Username/Password)
	if userRepo.Username.String == "" || userRepo.Password.String == "" {
		out.ResponseOut(c, nil, false, constanta.CodeBadRequestResponse, "Username/Password tidak boleh kosong")
		return err
	}

	// Check if the user exists and credentials are correct
	user, err = dao.UserDAO.LoginCheck(db, userRepo)
	if err != nil {
		out.ResponseOut(c, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return err
	}

	// If user not found or invalid credentials
	if user.ID.Int64 == 0 {
		out.ResponseOut(c, nil, false, constanta.CodeBadRequestResponse, "Username/Password salah")
		return err
	}

	// Generate the JWT token for the user
	token, err := config.GenerateToken(user)
	if err != nil {
		out.ResponseOut(c, nil, false, constanta.CodeInternalServerErrorResponse, err.Error())
		return
	}

	// Respond with the generated token
	out.ResponseOut(c, token, true, constanta.CodeSuccessResponse, "Login berhasil")
	return nil
}
