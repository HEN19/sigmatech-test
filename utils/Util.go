package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/api-skeleton/dto/in"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func StructToJSON(input interface{}) (output string) {
	b, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		output = ""
		return
	}
	output = string(b)
	return
}

func readBody(request *http.Request) (output string) {
	byteBody, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		return ""
	}
	return string(byteBody)
}

func GetUserBody(c *gin.Context) (in.UserRequest, error) {
	var userRequest in.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		return userRequest, err
	}
	return userRequest, nil
}

func GetCustomerBody(c *gin.Context) (in.CustomerRequest, error) {
	var customerReq in.CustomerRequest
	if err := c.ShouldBindJSON(&customerReq); err != nil {
		return customerReq, err
	}
	return customerReq, nil
}

func ReadParam(request *http.Request) (id int64, err error) {
	strId, ok := mux.Vars(request)["Id"]
	idParam, errConvert := strconv.Atoi(strId)
	id = int64(idParam)

	if !ok || errConvert != nil {
		err = errConvert
		return
	}
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
